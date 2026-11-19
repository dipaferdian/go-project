package ports

import (
	"context"

	"go-project/services/identity/internal/domain"
)

type UserRepository interface {
    ExistsByEmail(
        ctx context.Context,
        email domain.Email,
    ) (bool, error)

    Save(
        ctx context.Context,
        user *domain.User,
    ) error
}