package encoder

import (
	"encoding/base64"

	"github.com/ormushq/ormus/contract/go/destination"
	"google.golang.org/protobuf/proto"
)

func EncodeProcessedEvent(event *destination.DeliveredEventsList) string {
	payload, err := proto.Marshal(event)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(payload)
}

func DecodeProcessedEvent(event string) *destination.DeliveredEventsList {
	payload, err := base64.StdEncoding.DecodeString(event)
	if err != nil {
		return nil
	}
	mu := destination.DeliveredEventsList{}
	if err := proto.Unmarshal(payload, &mu); err != nil {
		return nil
	}

	return &destination.DeliveredEventsList{
		Events: mu.Events,
	}
}
