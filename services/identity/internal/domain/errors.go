package domain

import "errors"

var (
    ErrInvalidUserID     = errors.New("invalid user id")
    ErrInvalidEmail      = errors.New("invalid email")
    ErrInvalidPasswordHash = errors.New("invalid password hash")
    ErrInvalidUserStatus = errors.New("invalid user status")
    ErrUserDisabled      = errors.New("user is disabled")
)
