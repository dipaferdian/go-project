package domain

import (
	"errors"
	"testing"
	"time"
)

func createTestUser(t *testing.T) *User {
	t.Helper()

	email, err := NewEmail("user@example.com")
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	passwordHash, err := NewPasswordHash("$2a$12$example-hash")
	if err != nil {
		t.Fatalf("failed to create password hash: %v", err)
	}

	user, err := NewUser(
		UserID("user-123"),
		email,
		passwordHash,
		time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
	)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	return user
}

func TestNewUser(t *testing.T) {
	email, _ := NewEmail("user@example.com")
	passwordHash, _ := NewPasswordHash("$2a$12$example-hash")

	tests := []struct {
		name    string
		id      UserID
		email   Email
		hash    PasswordHash
		wantErr error
	}{
		{
			name:  "valid user",
			id:    UserID("user-123"),
			email: email,
			hash:  passwordHash,
		},
		{
			name:    "empty user id",
			id:      UserID(""),
			email:   email,
			hash:    passwordHash,
			wantErr: ErrInvalidUserID,
		},
		{
			name:    "invalid email",
			id:      UserID("user-123"),
			email:   Email{},
			hash:    passwordHash,
			wantErr: ErrInvalidEmail,
		},
		{
			name:    "invalid password hash",
			id:      UserID("user-123"),
			email:   email,
			hash:    PasswordHash{},
			wantErr: ErrInvalidPasswordHash,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(
				tt.id,
				tt.email,
				tt.hash,
				time.Time{},
			)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if user == nil {
				t.Fatal("expected user, got nil")
			}

			if user.ID() != tt.id {
				t.Errorf("expected ID %q, got %q", tt.id, user.ID())
			}

			if user.Status() != UserStatusActive {
				t.Errorf("expected active status, got %q", user.Status())
			}

			if user.CreatedAt().IsZero() {
				t.Error("expected createdAt to be initialized")
			}
		})
	}
}

func TestUserCanAuthenticate(t *testing.T) {
	user := createTestUser(t)

	err := user.CanAuthenticate()
	if err != nil {
		t.Fatalf("active user should be able to authenticate: %v", err)
	}

	user.Disable()

	err = user.CanAuthenticate()
	if !errors.Is(err, ErrUserDisabled) {
		t.Fatalf("expected ErrUserDisabled, got %v", err)
	}

	user.Activate()

	err = user.CanAuthenticate()
	if err != nil {
		t.Fatalf("activated user should be able to authenticate: %v", err)
	}
}

func TestUserDisable(t *testing.T) {
	user := createTestUser(t)

	user.Disable()

	if user.Status() != UserStatusDisabled {
		t.Errorf(
			"expected status %q, got %q",
			UserStatusDisabled,
			user.Status(),
		)
	}
}

func TestUserActivate(t *testing.T) {
	user := createTestUser(t)

	user.Disable()
	user.Activate()

	if user.Status() != UserStatusActive {
		t.Errorf(
			"expected status %q, got %q",
			UserStatusActive,
			user.Status(),
		)
	}
}

func TestUserChangePasswordHash(t *testing.T) {
	user := createTestUser(t)

	newHash, err := NewPasswordHash("$2a$12$new-example-hash")
	if err != nil {
		t.Fatalf("failed to create password hash: %v", err)
	}

	err = user.ChangePasswordHash(newHash)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.PasswordHash().Value() != newHash.Value() {
		t.Error("password hash was not updated")
	}

	err = user.ChangePasswordHash(PasswordHash{})
	if !errors.Is(err, ErrInvalidPasswordHash) {
		t.Errorf("expected ErrInvalidPasswordHash, got %v", err)
	}
}
