package eventvalidator

import (
	"context"
	proto_source "github.com/ormushq/ormus/contract/go/source"
	source_proto "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/source"
)

type Repository interface {
	IsWriteKeyValid(ctx context.Context, writeKey string) (bool, error)
	CreateNewWriteKey(ctx context.Context, writeKey *proto_source.NewSourceEvent, expirationTime uint) error
}

type ManagerWriteKeyValidation interface {
	IsWriteKeyValid(ctx context.Context, req *source_proto.ValidateWriteKeyReq) (*source_proto.ValidateWriteKeyResp, error)
}

type Validator struct {
	repo              Repository
	managerValidation ManagerWriteKeyValidation
	config            source.Config
}

func New(repo Repository, writeKeyValidation ManagerWriteKeyValidation, config source.Config) Validator {
	return Validator{
		repo:              repo,
		managerValidation: writeKeyValidation,
		config:            config,
	}
}
