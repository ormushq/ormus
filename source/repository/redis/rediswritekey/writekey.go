package rediswritekey

import (
	"context"
	"fmt"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source/params"
	"time"
)

func (r *DB) CreateNewWriteKey(ctx context.Context, WriteKey params.WriteKey, ExpirationTime uint) error {
	err := r.adapter.Client().Set(ctx, fmt.Sprintf("%s+%s", WriteKey.OwnerID, WriteKey.ProjectID),
		WriteKey.WriteKey, time.Minute*time.Duration(ExpirationTime)).Err()
	if err != nil {
		return richerror.New("source").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(err.Error())
	}
	return nil
}

func (r *DB) GetWriteKey(ctx context.Context, OwnerID string, ProjectID string) (*params.WriteKey, error) {
	wk, err := r.adapter.Client().Get(ctx, fmt.Sprintf("%s+%s", OwnerID, ProjectID)).Result()
	if err != nil {
		return nil, richerror.New("source.repository").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(err.Error())
	}

	return &params.WriteKey{
		OwnerID:   OwnerID,
		ProjectID: ProjectID,
		WriteKey:  wk,
	}, nil
}
