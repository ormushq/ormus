
package destination

import (
	"context"

	"github.com/ormushq/ormus/manager/param"
)

// Destination integration contains data that we use to process and deliver events.
// this integration is made or edited in the dashboard by the project owner
// this interface is a contractor between the manager and the destination layer https://github.com/ormushq/ormus/issues/35
type Destination interface {
	GetIntegrationByWriteKey(ctx context.Context, writeKey string) (param.ResponseIntegration, error)
	UpdateIntegrationByWriteKey(ctx context.Context, writeKey string, integration param.RequestIntegration) error
}
