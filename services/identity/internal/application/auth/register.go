package auth

import (
	"context"
	"errors"
	"time"

	"go-project/services/identity/internal/application/ports"
	"go-project/services/identity/internal/domain"
)

var (
    ErrInvalidPassword = errors.New("password cannot be empty")
)

type RegisterInput struct {
    Email    string
    Password string
}

type RegisterOutput struct {
    UserID    domain.UserID
    Email     string
    Status    domain.UserStatus
    CreatedAt time.Time
}

type RegisterUser struct {
    users         ports.UserRepository
    passwordHasher ports.PasswordHasher
    idGenerator    ports.IDGenerator
}

func NewRegisterUser(
    users ports.UserRepository,
    passwordHasher ports.PasswordHasher,
    idGenerator ports.IDGenerator,
) *RegisterUser {
    return &RegisterUser{
        users:          users,
        passwordHasher: passwordHasher,
        idGenerator:    idGenerator,
    }
}

func (uc *RegisterUser) Execute(
    ctx context.Context,
    input RegisterInput,
) (*RegisterOutput, error) {
    if input.Password == "" {
        return nil, ErrInvalidPassword
    }

    email, err := domain.NewEmail(input.Email)
    if err != nil {
        return nil, err
    }

    exists, err := uc.users.ExistsByEmail(ctx, email)
    if err != nil {
        return nil, err
    }

    if exists {
        return nil, ports.ErrEmailAlreadyExists
    }

    passwordHash, err := uc.passwordHasher.Hash(
        ctx,
        input.Password,
    )
    if err != nil {
        return nil, err
    }

    userID, err := uc.idGenerator.NewUserID()
    if err != nil {
        return nil, err
    }

    user, err := domain.NewUser(
        userID,
        email,
        passwordHash,
        time.Now(),
    )
    if err != nil {
        return nil, err
    }

    if err := uc.users.Save(ctx, user); err != nil {
        return nil, err
    }

    return &RegisterOutput{
        UserID:    user.ID(),
        Email:     user.Email().Value(),
        Status:    user.Status(),
        CreatedAt: user.CreatedAt(),
    }, nil
}
