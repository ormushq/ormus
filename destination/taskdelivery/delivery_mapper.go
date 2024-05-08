package taskdelivery

import "github.com/ormushq/ormus/manager/entity"

var Mapper = make(map[entity.DestinationType]DeliveryHandler)

// Register registers a new delivery handler for a destination type.
func Register(destinationType entity.DestinationType, dh DeliveryHandler) {
	Mapper[destinationType] = dh
}
