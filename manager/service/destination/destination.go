package destination

import "github.com/ormushq/ormus/manager/entity"

type Destination interface {
	GetIntegrationByWriteKey(writeKey string) (entity.Integration, error)
	// TODO: we need a function that is responsible for update integration in destination cache db when Integration edited
}
