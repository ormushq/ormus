package writekey

import (
	"context"
	rabbitmq "github.com/ormushq/ormus/adapter/rabbitmq"
	proto_source "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source"
)

type WriteKeyRepo interface {
	CreateNewWriteKey(ctx context.Context, WriteKey proto_source.NewSourceEvent, ExpirationTime uint) error
	GetWriteKey(ctx context.Context, OwnerID string, ProjectID string) (*proto_source.NewSourceEvent, error)
}

type PublisherRepo interface {
	Publish(queueName string, message []byte) error
	Close()
}

type ConsumerRepo interface {
	Subscribe(queueName string) (chan *rabbitmq.Message, error)
	Close()
	Ack(msg *rabbitmq.Message) error
}

type Service struct {
	Publisher    PublisherRepo
	Consumer     ConsumerRepo
	WriteKeyRepo WriteKeyRepo
	config       source.Config
}

func New(Publisher PublisherRepo, Consumer ConsumerRepo, WriteKeyRepo WriteKeyRepo, config source.Config) Service {
	return Service{
		Publisher:    Publisher,
		Consumer:     Consumer,
		WriteKeyRepo: WriteKeyRepo,
		config:       config,
	}
}

func (s Service) CreateNewWriteKey(ctx context.Context, OwnerID string, ProjectID string, WriteKey string) error {
	err := s.WriteKeyRepo.CreateNewWriteKey(ctx, proto_source.NewSourceEvent{
		ProjectId: ProjectID,
		OwnerId:   OwnerID,
		WriteKey:  WriteKey,
	}, s.config.WriteKeyRedisExpiration)
	if err != nil {
		return richerror.New("source.service").WithWrappedError(err)
	}
	return nil
}
