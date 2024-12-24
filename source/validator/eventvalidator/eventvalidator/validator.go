package eventvalidator

import (
	"context"

	"github.com/ormushq/ormus/source"
)

type Repository interface {
	IsWriteKeyValid(ctx context.Context, writeKey string, expirationTime uint) (bool, error)
}

type Validator struct {
	repo   Repository
	config source.Config
}

func New(repo Repository, cfg source.Config) Validator {
	return Validator{
		repo:   repo,
		config: cfg,
	}
}
