package auth

import (
	"context"
	"errors"
	"testing"

	"go-project/services/identity/internal/application/ports"
	"go-project/services/identity/internal/domain"
)

// Fake UserRepository

type fakeUserRepository struct {
    exists bool
    saved  *domain.User
    saveErr error
    existsErr error
}

func (f *fakeUserRepository) ExistsByEmail(
    ctx context.Context,
    email domain.Email,
) (bool, error) {
    return f.exists, f.existsErr
}

func (f *fakeUserRepository) Save(
    ctx context.Context,
    user *domain.User,
) error {
    if f.saveErr != nil {
        return f.saveErr
    }

    f.saved = user
    return nil
}

// Fake PasswordHasher

type fakePasswordHasher struct {
    hash domain.PasswordHash
    err  error
}

func (f *fakePasswordHasher) Hash(
    ctx context.Context,
    password string,
) (domain.PasswordHash, error) {
    return f.hash, f.err
}

// Fake IDGenerator

type fakeIDGenerator struct {
    id  domain.UserID
    err error
}

func (f *fakeIDGenerator) NewUserID() (domain.UserID, error) {
    return f.id, f.err
}

func newTestRegisterUser(
    repo *fakeUserRepository,
    hasher *fakePasswordHasher,
    generator *fakeIDGenerator,
) *RegisterUser {
    return NewRegisterUser(repo, hasher, generator)
}

func TestRegisterUserExecute(t *testing.T) {
    passwordHash, err := domain.NewPasswordHash("hashed-password")
    if err != nil {
        t.Fatal(err)
    }

    tests := []struct {
        name       string
        input      RegisterInput
        repo       *fakeUserRepository
        hasher     *fakePasswordHasher
        generator  *fakeIDGenerator
        wantErr    error
    }{
        {
            name: "successful registration",
            input: RegisterInput{
                Email:    "user@example.com",
                Password: "secret-password",
            },
            repo: &fakeUserRepository{},
            hasher: &fakePasswordHasher{
                hash: passwordHash,
            },
            generator: &fakeIDGenerator{
                id: domain.UserID("user-123"),
            },
        },
        {
            name: "invalid email",
            input: RegisterInput{
                Email:    "invalid-email",
                Password: "secret-password",
            },
            repo:      &fakeUserRepository{},
            hasher:    &fakePasswordHasher{hash: passwordHash},
            generator: &fakeIDGenerator{id: "user-123"},
            wantErr:   domain.ErrInvalidEmail,
        },
        {
            name: "empty password",
            input: RegisterInput{
                Email:    "user@example.com",
                Password: "",
            },
            repo:      &fakeUserRepository{},
            hasher:    &fakePasswordHasher{hash: passwordHash},
            generator: &fakeIDGenerator{id: "user-123"},
            wantErr:   ErrInvalidPassword,
        },
        {
            name: "email already registered",
            input: RegisterInput{
                Email:    "user@example.com",
                Password: "secret-password",
            },
            repo: &fakeUserRepository{
                exists: true,
            },
            hasher:    &fakePasswordHasher{hash: passwordHash},
            generator: &fakeIDGenerator{id: "user-123"},
            wantErr:   ports.ErrEmailAlreadyExists,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            uc := newTestRegisterUser(
                tt.repo,
                tt.hasher,
                tt.generator,
            )

            result, err := uc.Execute(
                context.Background(),
                tt.input,
            )

            if tt.wantErr != nil {
                if !errors.Is(err, tt.wantErr) {
                    t.Fatalf(
                        "expected error %v, got %v",
                        tt.wantErr,
                        err,
                    )
                }

                if result != nil {
                    t.Fatal("expected nil result")
                }

                return
            }

            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }

            if result == nil {
                t.Fatal("expected registration result")
            }

            if result.UserID != "user-123" {
                t.Errorf("unexpected user ID: %s", result.UserID)
            }

            if result.Email != "user@example.com" {
                t.Errorf("unexpected email: %s", result.Email)
            }

            if result.Status != domain.UserStatusActive {
                t.Errorf("unexpected status: %s", result.Status)
            }

            if tt.repo.saved == nil {
                t.Fatal("expected user to be saved")
            }
        })
    }
}
