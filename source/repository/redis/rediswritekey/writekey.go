package rediswritekey

import (
	"context"
	"fmt"
	proto_source "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/pkg/richerror"
	"time"
)

func (r *DB) CreateNewWriteKey(ctx context.Context, writeKey proto_source.NewSourceEvent, ExpirationTime uint) error {
	err := r.adapter.Client().Set(ctx, fmt.Sprintf("%s+%s", writeKey.OwnerId, writeKey.ProjectId),
		writeKey.WriteKey, time.Minute*time.Duration(ExpirationTime)).Err()
	if err != nil {
		return richerror.New("source").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(err.Error())
	}

	return nil

}

func (r *DB) GetWriteKey(ctx context.Context, ownerID string, projectID string) (*proto_source.NewSourceEvent, error) {
	wk, err := r.adapter.Client().Get(ctx, fmt.Sprintf("%s+%s", ownerID, projectID)).Result()
	if err != nil {
		return nil, richerror.New("source.repository").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(err.Error())

	}

	return &proto_source.NewSourceEvent{
		OwnerId:   ownerID,
		ProjectId: projectID,
		WriteKey:  wk,
	}, nil

}
