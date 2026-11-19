package domain

import (
	"net/mail"
	"strings"
)

type Email struct {
    value string
}

func NewEmail(value string) (Email, error) {
    value = strings.ToLower(strings.TrimSpace(value))

    if value == "" {
        return Email{}, ErrInvalidEmail
    }

    parsed, err := mail.ParseAddress(value)
    if err != nil || parsed.Address != value {
        return Email{}, ErrInvalidEmail
    }

    return Email{value: value}, nil
}

func (e Email) Value() string {
    return e.value
}

type PasswordHash struct {
    value string
}

func NewPasswordHash(value string) (PasswordHash, error) {
    if strings.TrimSpace(value) == "" {
        return PasswordHash{}, ErrInvalidPasswordHash
    }

    return PasswordHash{value: value}, nil
}

func (p PasswordHash) Value() string {
    return p.value
}
