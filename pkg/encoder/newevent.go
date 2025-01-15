package encoder

import (
	"encoding/base64"

	"github.com/ormushq/ormus/contract/go/source"
	"google.golang.org/protobuf/proto"
)

func EncodeNewEvent(event *source.NewEvent) string {
	payload, err := proto.Marshal(event)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(payload)
}

func DecodeNewEvent(event string) *source.NewEvent {
	payload, err := base64.StdEncoding.DecodeString(event)
	if err != nil {
		return nil
	}
	mu := source.NewEvent{}
	if err := proto.Unmarshal(payload, &mu); err != nil {
		return nil
	}

	return &source.NewEvent{
		Id:         mu.Id,
		Type:       mu.Type,
		Name:       mu.Name,
		Properties: mu.Properties,
		SendAt:     mu.SendAt,
		ReceivedAt: mu.ReceivedAt,
		Timestamp:  mu.Timestamp,
		Event:      mu.Event,
		WriteKey:   mu.WriteKey,
		CreatedAt:  mu.CreatedAt,
		UpdatedAt:  mu.UpdatedAt,
	}
}
