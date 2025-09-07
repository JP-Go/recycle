package services

import (
	"backend/internal/domain/repositories"
	"backend/internal/domain/valueobjects"
	"context"
)

type UniqueEmailVerifier interface {
	Verify(context.Context, valueobjects.Email) (bool, error)
}

type BasicUniqueEmailVerifier struct {
	repository repositories.UserRepository
}

func NewUniqueEmailVerifier(ur repositories.UserRepository) *BasicUniqueEmailVerifier {
	return &BasicUniqueEmailVerifier{
		repository: ur,
	}
}

func (bev *BasicUniqueEmailVerifier) Verify(ctx context.Context, email valueobjects.Email) (bool, error) {
	return bev.repository.IsEmailUnique(ctx, email)
}
