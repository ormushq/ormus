package destination

import (
	"context"

	"github.com/ormushq/ormus/manager/param"
)

type Destination interface {
	GetIntegrationByWriteKey(ctx context.Context, writeKey string) (param.ResponseIntegration, error)
	UpdateIntegrationByWriteKey(ctx context.Context, writeKey string, integration param.RequestIntegration) error
}
