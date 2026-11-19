package ports

import "go-project/services/identity/internal/domain"

type IDGenerator interface {
    NewUserID() (domain.UserID, error)
}