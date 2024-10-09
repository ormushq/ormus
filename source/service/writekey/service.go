package writekey

import (
	"context"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source"
	"github.com/ormushq/ormus/source/params"
)

type Repository interface {
	// TODO - implementation writekeyadapter
	IsValidWriteKey(ctx context.Context, writeKey string) (bool, error)
}

type WriteKeyRepo interface {
	CreateNewWriteKey(ctx context.Context, WriteKey params.WriteKey, ExpirationTime uint) error
	GetWriteKey(ctx context.Context, OwnerID string, ProjectID string) (*params.WriteKey, error)
}

type Service struct {
	repo         Repository
	WriteKeyRepo WriteKeyRepo
	config       source.Config
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

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

func (s Service) CreateNewWriteKey(ctx context.Context, OwnerID string, ProjectID string, WriteKey string) error {
	err := s.WriteKeyRepo.CreateNewWriteKey(ctx, params.WriteKey{
		ProjectID: ProjectID,
		OwnerID:   OwnerID,
		WriteKey:  WriteKey,
	}, s.config.WritekeyRedisExpiration)
	if err != nil {
		return richerror.New("source.service").WithWrappedError(err)
	}
	return nil
}
