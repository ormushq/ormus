package writekey

import (
	"context"
)

// Repository is an interface representing what methods should be implemented by the repository.
type Repository interface {
	// TODO - implementation redis
	IsValidWriteKey(ctx context.Context, writeKey string) (bool, error)
}

// Service show dependencies writeKey service.
type Service struct {
	repo Repository
}

// Constructor writeKey service.
func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

// IsValid checks whether the writeKey is valid or not.
func (s Service) IsValid(ctx context.Context, writeKey string) (bool, error) {
	// TODO - How errmsg handling ? Rich-errmsg or ...?
	isValid, err := s.repo.IsValidWriteKey(ctx, writeKey)
	if err != nil {
		// TODO - logger
		return false, err
	}
	if !isValid {
		return false, err
	}

	return true, nil

}
