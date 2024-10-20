package protobufmapper

import (
	"fmt"
	protobufEvent "github.com/ormushq/ormus/contract/go/event"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/manager/entity"
	fakeintgration "github.com/ormushq/ormus/manager/entity/integrations/fakeconfig"
	"github.com/ormushq/ormus/manager/entity/integrations/webhookintegration"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
)

// MapProcessedEventFromProtobuf TODO: update me with protobufEvent.ProcessedEvent
func MapProcessedEventFromProtobuf(pe *protobufEvent.ProcessedEvent) event.ProcessedEvent {
	integration := entity.Integration{
		ID:       pe.Integration.Id,
		SourceID: pe.Integration.SourceId,
		Name:     pe.Integration.Name,
		Metadata: entity.DestinationMetadata{
			ID:         pe.Integration.Metadata.Id,
			Name:       pe.Integration.Metadata.Name,
			Slug:       entity.DestinationType(pe.Integration.Metadata.Slug.String()),
			Categories: make([]entity.DestinationCategory, len(pe.Integration.Metadata.Categories)),
		},
		ConnectionType: entity.ConnectionType(pe.Integration.ConnectionType.String()),
		Enabled:        pe.Integration.Enabled,
		CreatedAt:      pe.Integration.CreatedAt.AsTime(),
	}

	switch config := pe.Integration.Config.(type) {
	case *protobufEvent.Integration_Fake:
		integration.Config = fakeintgration.FakeConfig{Name: config.Fake.Name}
	case *protobufEvent.Integration_Webhook:
		headers := make(map[string]string)
		for k, v := range config.Webhook.Headers {
			headers[k] = v
		}
		payload := make(map[string]string)
		for k, v := range config.Webhook.Payload {
			payload[k] = v
		}
		integration.Config = webhookintegration.WebhookConfig{
			Headers: headers,
			Payload: payload,
			Method:  webhookintegration.WebhookMethod(config.Webhook.Method.String()),
			URL:     config.Webhook.Url,
		}
	}

	for i, category := range pe.Integration.Metadata.Categories {
		integration.Metadata.Categories[i] = entity.DestinationCategory(category.String())
	}

	return event.ProcessedEvent{
		SourceID:          pe.SourceId,
		Integration:       integration,
		MessageID:         pe.MessageId,
		EventType:         event.Type(pe.EventType.String()),
		Version:           uint8(pe.Version),
		SentAt:            pe.SentAt.AsTime(),
		ReceivedAt:        pe.ReceivedAt.AsTime(),
		OriginalTimestamp: pe.OriginalTimestamp.AsTime(),
		Timestamp:         pe.Timestamp.AsTime(),
	}
}

func MapProcessedEventToProtobuf(e event.ProcessedEvent) *protobufEvent.ProcessedEvent {
	integration := &protobufEvent.Integration{
		Id:       e.Integration.ID,
		SourceId: e.Integration.SourceID,
		Name:     e.Integration.Name,
		ConnectionType: protobufEvent.ConnectionType(
			protobufEvent.ConnectionType_value[string(e.Integration.ConnectionType)]),
		Enabled:   e.Integration.Enabled,
		CreatedAt: timestamppb.New(e.Integration.CreatedAt),
		Metadata: &protobufEvent.DestinationMetadata{
			Id:   e.Integration.Metadata.ID,
			Name: e.Integration.Metadata.Name,
			Slug: protobufEvent.DestinationType(
				protobufEvent.DestinationType_value[string(e.Integration.Metadata.Slug)]),
		},
	}

	for _, category := range e.Integration.Metadata.Categories {
		integration.Metadata.Categories = append(integration.Metadata.Categories,
			protobufEvent.DestinationCategory(protobufEvent.DestinationCategory_value[string(category)]))
	}

	switch config := e.Integration.Config.(type) {
	case fakeintgration.FakeConfig:
		fmt.Println("--fakeintgration.FakeConfig--")
		integration.Config = &protobufEvent.Integration_Fake{
			Fake: &protobufEvent.FakeConfig{
				Name: config.Name,
			},
		}
	case webhookintegration.WebhookConfig:
		fmt.Println("-webhookintegration.WebhookConfig--")

		headers := make(map[string]string)
		for k, v := range config.Headers {
			headers[k] = v
		}
		payload := make(map[string]string)
		for k, v := range config.Payload {
			payload[k] = v
		}
		integration.Config = &protobufEvent.Integration_Webhook{
			Webhook: &protobufEvent.WebhookConfig{
				Headers: headers,
				Payload: payload,
				Method:  protobufEvent.WebhookMethod(protobufEvent.WebhookMethod_value[string(config.Method)]),
				Url:     config.URL,
			},
		}
	default:
		fmt.Println("--default what what what--")
		fmt.Println(reflect.TypeOf(config))
	}

	// TODO: complete me
	return &protobufEvent.ProcessedEvent{
		SourceId:          e.SourceID,
		TracerCarrier:     e.TracerCarrier,
		Integration:       integration,
		MessageId:         e.MessageID,
		EventType:         protobufEvent.Type(protobufEvent.Type_value[string(e.EventType)]),
		Version:           uint32(e.Version),
		SentAt:            timestamppb.New(e.SentAt),
		ReceivedAt:        timestamppb.New(e.ReceivedAt),
		OriginalTimestamp: timestamppb.New(e.OriginalTimestamp),
		Timestamp:         timestamppb.New(e.Timestamp),

		UserId:      "",
		AnonymousId: "",
		Event:       "",
		Name:        "",
		GroupId:     "",
		PreviousId:  "",
		Context:     &protobufEvent.Context{},
		Properties:  &protobufEvent.Properties{},
		Traits:      &protobufEvent.Traits{},
	}
}

// MapContextToProtobuf TODO: implement me
func MapContextToProtobuf(e event.Context) *protobufEvent.Context {
	return &protobufEvent.Context{
		Active:        false,
		Ip:            "",
		Locale:        "",
		Location:      nil,
		Page:          nil,
		UserAgent:     "",
		UserAgentData: nil,
		Library:       nil,
		Traits:        nil,
		Campaign:      nil,
		Referrer:      nil,
		CustomData:    nil,
	}
}

// MapPropertiesToProtobuf TODO: implement me
func MapPropertiesToProtobuf(e event.Properties) *protobufEvent.Properties {
	return &protobufEvent.Properties{}
}

// MapTraitsToProtobuf TODO: implement me
func MapTraitsToProtobuf(e event.Traits) *protobufEvent.Traits {
	return &protobufEvent.Traits{}
}
