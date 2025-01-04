package rediswritekey

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	proto_source "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/redis/go-redis/v9"
)

func (r *DB) CreateNewWriteKey(ctx context.Context, writeKey *proto_source.NewSourceEvent, expirationTime uint) error {
	v, err := json.Marshal(
		map[string]string{"OwnerId": writeKey.OwnerId, "ProjectId": writeKey.ProjectId},
	)
	if err != nil {
		logger.L().Error(err.Error())

		return richerror.New("source").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(err.Error())
	}
	err = r.adapter.Client().Set(ctx, writeKey.WriteKey, v,
		time.Minute*time.Duration(expirationTime)).Err()
	if err != nil {
		logger.L().Error(err.Error())

		return richerror.New("source").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(err.Error())
	}

	return nil
}

func (r *DB) IsWriteKeyValid(ctx context.Context, writeKey string, expirationTime uint) (bool, error) {
	result, err, _ := r.singleFlight.Do(writeKey, func() (interface{}, error) {
		err := r.adapter.Client().Get(ctx, writeKey).Err()
		if errors.Is(err, redis.Nil) {
			// Cache miss, need to validate the write key
			resp, validateErr := r.ManagerAdapter.IsWriteKeyValid(ctx, &proto_source.ValidateWriteKeyReq{
				WriteKey: writeKey,
			})
			if validateErr != nil {
				logger.L().Error(validateErr.Error())

				var richErr richerror.RichError
				if errors.As(validateErr, &richErr) && richErr.Kind() == richerror.KindInvalid {
					return false, nil
				}

				return false, richerror.New("source.repository").
					WithWrappedError(validateErr).
					WithKind(richerror.KindUnexpected).
					WithMessage(validateErr.Error())
			}

			// If valid, create a new write key in the cache
			if resp != nil && resp.IsValid {
				err := r.CreateNewWriteKey(ctx,
					&proto_source.NewSourceEvent{
						WriteKey:  resp.WriteKey,
						ProjectId: resp.ProjectId,
						OwnerId:   resp.OwnerId,
					}, expirationTime,
				)
				if err != nil {
					logger.L().Error(err.Error())

					return nil, richerror.New("source.repository").
						WithWrappedError(err).
						WithKind(richerror.KindUnexpected).
						WithMessage(err.Error())
				}

				return true, nil
			}

			// If the write key is not valid, return false
			return false, nil
		}

		if err != nil {
			logger.L().Error(err.Error())

			return false, richerror.New("source.repository").
				WithWrappedError(err).
				WithKind(richerror.KindUnexpected).
				WithMessage(err.Error())
		}

		// If the key is found in Redis, return true as it's valid
		return true, nil
	})

	if err != nil {
		logger.L().Error(err.Error())

		return false, err
	}

	valid, ok := result.(bool)
	if !ok {
		return false, richerror.New("source.repository").
			WithMessage("invalid result type, expected bool")
	}

	return valid, nil
}
