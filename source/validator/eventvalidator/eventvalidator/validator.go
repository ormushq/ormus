package eventvalidator

import (
	"context"
)

type Repository interface {
	IsWriteKeyValid(ctx context.Context, writeKey string) (bool, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
