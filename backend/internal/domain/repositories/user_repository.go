package repositories

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/valueobjects"
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("User not found")

type UserRepository interface {
	FindById(ctx context.Context, user entities.UserId) (*entities.ValidatedUser, error)
	FindByEmail(ctx context.Context, email valueobjects.Email) (*entities.ValidatedUser, error)
	Create(ctx context.Context, user *entities.ValidatedUser) (*entities.ValidatedUser, error)
	IsEmailUnique(ctx context.Context, email valueobjects.Email) (bool, error)
}
