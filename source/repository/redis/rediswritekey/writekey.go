package rediswritekey

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	proto_source "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/redis/go-redis/v9"
)

func (r *DB) CreateNewWriteKey(ctx context.Context, writeKey *proto_source.NewSourceEvent, expirationTime uint) error {
	v, err := json.Marshal(
		map[string]string{"OwnerId": writeKey.OwnerId, "ProjectId": writeKey.ProjectId},
	)
	if err != nil {
		return richerror.New("source").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(err.Error())
	}
	err = r.adapter.Client().Set(ctx, writeKey.WriteKey, v,
		time.Minute*time.Duration(expirationTime)).Err()
	if err != nil {
		return richerror.New("source").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(err.Error())
	}

	return nil
}

func (r *DB) IsWriteKeyValid(ctx context.Context, writeKey string) (bool, error) {
	err := r.adapter.Client().Get(ctx, writeKey).Err()
	if errors.Is(err, redis.Nil) { // Use errors.Is to check for redis.Nil
		return false, nil
	}
	if err != nil {
		return false, richerror.New("source.repository").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(err.Error())
	}

	return true, nil
}
