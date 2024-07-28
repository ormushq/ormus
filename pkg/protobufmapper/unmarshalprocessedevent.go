package protobufmapper

import (
	"github.com/ormushq/ormus/contract/goprotobuf/processedevent"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/manager/entity"
	fakeintgration "github.com/ormushq/ormus/manager/entity/integrations/fakeconfig"
	"github.com/ormushq/ormus/manager/entity/integrations/webhookintegration"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapProcessedEventFromProtobuf(pe *processedevent.ProcessedEvent) event.ProcessedEvent {
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
	case *processedevent.Integration_Fake:
		integration.Config = fakeintgration.FakeConfig{Name: config.Fake.Name}
	case *processedevent.Integration_Webhook:
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

func MapProcessedEventToProtobuf(e event.ProcessedEvent) *processedevent.ProcessedEvent {
	integration := &processedevent.Integration{
		Id:       e.Integration.ID,
		SourceId: e.Integration.SourceID,
		Name:     e.Integration.Name,
		ConnectionType: processedevent.ConnectionType(
			processedevent.ConnectionType_value[string(e.Integration.ConnectionType)]),
		Enabled:   e.Integration.Enabled,
		CreatedAt: timestamppb.New(e.Integration.CreatedAt),
		Metadata: &processedevent.DestinationMetadata{
			Id:   e.Integration.Metadata.ID,
			Name: e.Integration.Metadata.Name,
			Slug: processedevent.DestinationType(
				processedevent.DestinationType_value[string(e.Integration.Metadata.Slug)]),
		},
	}

	for i, category := range e.Integration.Metadata.Categories {
		integration.Metadata.Categories[i] = processedevent.DestinationCategory(
			processedevent.DestinationCategory_value[string(category)])
	}

	switch config := e.Integration.Config.(type) {
	case fakeintgration.FakeConfig:
		integration.Config = &processedevent.Integration_Fake{
			Fake: &processedevent.FakeConfig{
				Name: config.Name,
			},
		}
	case webhookintegration.WebhookConfig:
		headers := make(map[string]string)
		for k, v := range config.Headers {
			headers[k] = v
		}
		payload := make(map[string]string)
		for k, v := range config.Payload {
			payload[k] = v
		}
		integration.Config = &processedevent.Integration_Webhook{
			Webhook: &processedevent.WebhookConfig{
				Headers: headers,
				Payload: payload,
				Method:  processedevent.WebhookMethod(processedevent.WebhookMethod_value[string(config.Method)]),
				Url:     config.URL,
			},
		}
	}

	return &processedevent.ProcessedEvent{
		SourceId:          e.SourceID,
		Integration:       integration,
		MessageId:         e.MessageID,
		EventType:         processedevent.Type(processedevent.Type_value[string(e.EventType)]),
		Version:           uint32(e.Version),
		SentAt:            timestamppb.New(e.SentAt),
		ReceivedAt:        timestamppb.New(e.ReceivedAt),
		OriginalTimestamp: timestamppb.New(e.OriginalTimestamp),
		Timestamp:         timestamppb.New(e.Timestamp),
	}
}
