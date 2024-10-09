package encoder

import (
	"encoding/base64"
	"github.com/ormushq/ormus/contract/go/source"
	"google.golang.org/protobuf/proto"
)

func EncodeNewSourceEvent(NewSource source.NewSourceEvent) string {
	payload, err := proto.Marshal(&NewSource)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(payload)

}

func DecodeNewSourceEvent(NewSourceEvent string) *source.NewSourceEvent {
	payload, err := base64.StdEncoding.DecodeString(NewSourceEvent)
	if err != nil {
		return nil
	}
	mu := source.NewSourceEvent{}
	if err := proto.Unmarshal(payload, &mu); err != nil {
		return nil
	}

	return &source.NewSourceEvent{
		ProjectId: (mu.ProjectId),
		OwnerId:   mu.OwnerId,
		WriteKey:  mu.WriteKey,
	}
}
