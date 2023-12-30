package destination

import (
	"context"
	"github.com/ormushq/ormus/destination/entity"
)

type Destination interface {
	GetIntegrationByWriteKey(ctx context.Context, writeKey string) (entity.Integration, error)
	UpdateIntegrationByWriteKey(ctx context.Context, writeKey string, integration entity.Integration) error
}
