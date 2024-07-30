package writekey

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/contract/protobuf/manager/goproto/writekey"
	"github.com/ormushq/ormus/manager/entity"
	"time"
)

type DB struct {
	adapter       redis.Adapter
	managerClient writekey.WriteKeyManagerClient
}

func New(adapter redis.Adapter, managerClient writekey.WriteKeyManagerClient) DB {
	return DB{
		adapter:       adapter,
		managerClient: managerClient,
	}
}

func (db DB) IsValidWriteKey(ctx context.Context, writeKey string) (*entity.WriteKeyMetaData, error) {
	// try to get it from redis first
	result, err := db.getFromRedis(ctx, writeKey)
	if err == nil {
		return result, nil
	}

	// Key was not in the cache, make a GRPC call to manager
	resp, err := db.managerClient.GetWriteKey(ctx, &writekey.GetWriteKeyRequest{
		WriteKey: writeKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get write key from manager: %w", err)
	}

	metadata := &entity.WriteKeyMetaData{
		WriteKey:   resp.Metadata.WriteKey,
		OwnerID:    resp.Metadata.OwnerId,
		SourceID:   resp.Metadata.SourceId,
		CreatedAt:  resp.Metadata.CreatedAt.AsTime(),
		LastUsedAt: resp.Metadata.LastUsedAt.AsTime(),
		Status:     entity.WriteKeyStatus(resp.Metadata.Status.String()),
	}

	// cache the metadata
	if err := db.updateRedis(ctx, writeKey, metadata); err != nil { //TODO: retrying strategy?
		// Continue even though there is an error
		// TODO: logging maybe! I suppose nothing happen when Redis's not updated, we send GRPC.
	}

	return metadata, nil
}

func (db DB) getFromRedis(ctx context.Context, writeKey string) (*entity.WriteKeyMetaData, error) {
	metadataJSON, err := db.adapter.Client().Get(ctx, "writekey:"+writeKey).Result()
	if err == nil {
		var metadata entity.WriteKeyMetaData
		if err := json.Unmarshal([]byte(metadataJSON), &metadata); err != nil {
			return nil, fmt.Errorf("failed to unmarshal writekey metadata: %w", err)
		}

		return &metadata, nil
	}

	return &entity.WriteKeyMetaData{}, err
}

func (db DB) updateRedis(ctx context.Context, writeKey string, metadata *entity.WriteKeyMetaData) error {
	metadataJSON, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("failed to marshal writekey metadata: %w", err)
	}
	// we set TTL for write keys, we don't need to fill up the memory with
	// all the already generated write keys.
	return db.adapter.Client().Set(ctx, "writekey:"+writeKey, metadataJSON, 24*time.Hour).Err()
}

// InvalidateWriteKey is mechanism of write key invalidation in cache(Redis) when
// a write key is updated or deleted in the core(manager) service.
func (db DB) InvalidateWriteKey(ctx context.Context, writeKey string) error {
	err := db.adapter.Client().Del(ctx, "writekey:"+writeKey).Err()
	if err != nil {
		return fmt.Errorf("failed to invalidate write key in Redis: %w", err)
	}

	return nil
}
