package domain

import (
	"errors"
	"testing"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{
			name:  "valid email",
			input: "user@example.com",
			want:  "user@example.com",
		},
		{
			name:  "email with uppercase",
			input: "USER@EXAMPLE.COM",
			want:  "user@example.com",
		},
		{
			name:    "empty email",
			input:   "",
			wantErr: ErrInvalidEmail,
		},
		{
			name:    "invalid email",
			input:   "invalid-email",
			wantErr: ErrInvalidEmail,
		},
		{
			name:    "email without domain",
			input:   "user@",
			wantErr: ErrInvalidEmail,
		},
		{
			name:    "email with spaces",
			input:   " user@example.com ",
			want:    "user@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEmail(tt.input)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got.Value() != tt.want {
				t.Errorf("expected %q, got %q", tt.want, got.Value())
			}
		})
	}
}

func TestNewPasswordHash(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:  "valid password hash",
			input: "$2a$12$example-hash",
		},
		{
			name:    "empty password hash",
			input:   "",
			wantErr: ErrInvalidPasswordHash,
		},
		{
			name:    "whitespace password hash",
			input:   "   ",
			wantErr: ErrInvalidPasswordHash,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPasswordHash(tt.input)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got.Value() != tt.input {
				t.Errorf("expected %q, got %q", tt.input, got.Value())
			}
		})
	}
}
