package integrationmanager

import (
	"context"

	"github.com/ormushq/ormus/manager/param"
)

// Manager integration contains data that we use to process and deliver events.
// this integration is made or edited in the dashboard by the project owner.
type Manager interface {
	GetIntegrationByWriteKey(ctx context.Context, writeKey string) (param.ResponseIntegration, error)
}
