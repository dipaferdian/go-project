package domain

import (
	"strings"
	"time"
)

type UserID string

type UserStatus string

const (
    UserStatusActive   UserStatus = "active"
    UserStatusDisabled UserStatus = "disabled"
)

type User struct {
    id           UserID
    email        Email
    passwordHash PasswordHash
    status       UserStatus
    createdAt    time.Time
}

func NewUser(
    id UserID,
    email Email,
    passwordHash PasswordHash,
    createdAt time.Time,
) (*User, error) {
    if strings.TrimSpace(string(id)) == "" {
        return nil, ErrInvalidUserID
    }

    if email.Value() == "" {
        return nil, ErrInvalidEmail
    }

    if passwordHash.Value() == "" {
        return nil, ErrInvalidPasswordHash
    }

    if createdAt.IsZero() {
        createdAt = time.Now()
    }

    return &User{
        id:           id,
        email:        email,
        passwordHash: passwordHash,
        status:       UserStatusActive,
        createdAt:    createdAt,
    }, nil
}

func (u *User) ID() UserID {
    return u.id
}

func (u *User) Email() Email {
    return u.email
}

func (u *User) PasswordHash() PasswordHash {
    return u.passwordHash
}

func (u *User) Status() UserStatus {
    return u.status
}

func (u *User) CreatedAt() time.Time {
    return u.createdAt
}

func (u *User) CanAuthenticate() error {
    if u.status == UserStatusDisabled {
        return ErrUserDisabled
    }

    return nil
}

func (u *User) Disable() {
    u.status = UserStatusDisabled
}

func (u *User) Activate() {
    u.status = UserStatusActive
}

func (u *User) ChangePasswordHash(hash PasswordHash) error {
    if hash.Value() == "" {
        return ErrInvalidPasswordHash
    }

    u.passwordHash = hash
    return nil
}
