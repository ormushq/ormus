package rediswritekey

import (
	"context"
	"fmt"
	proto_source "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/pkg/richerror"
	"time"
)

func (r *DB) CreateNewWriteKey(ctx context.Context, WriteKey proto_source.NewSourceEvent, ExpirationTime uint) error {
	err := r.adapter.Client().Set(ctx, fmt.Sprintf("%s+%s", WriteKey.OwnerId, WriteKey.ProjectId),
		WriteKey.WriteKey, time.Minute*time.Duration(ExpirationTime)).Err()
	if err != nil {
		return richerror.New("source").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(err.Error())
	}
	return nil
}

func (r *DB) GetWriteKey(ctx context.Context, OwnerID string, ProjectID string) (*proto_source.NewSourceEvent, error) {
	wk, err := r.adapter.Client().Get(ctx, fmt.Sprintf("%s+%s", OwnerID, ProjectID)).Result()
	if err != nil {
		return nil, richerror.New("source.repository").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(err.Error())
	}

	return &proto_source.NewSourceEvent{
		OwnerId:   OwnerID,
		ProjectId: ProjectID,
		WriteKey:  wk,
	}, nil
}
