package ports

import (
	"context"

	"go-project/services/identity/internal/domain"
)

type PasswordHasher interface {
    Hash(
        ctx context.Context,
        password string,
    ) (domain.PasswordHash, error)
}