package main

import (
	"sync"

	"github.com/google/uuid"
	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/pkg/channel"
	rbbitmqchannel "github.com/ormushq/ormus/pkg/channel/adapter/rabbitmq"
	"github.com/ormushq/ormus/pkg/encoder"
)

func main() {
	cfg := config.C()
	done := make(chan bool)
	wg := &sync.WaitGroup{}
	dbConfig := dconfig.RabbitMQConsumerConnection{
		User:            cfg.RabbitMq.UserName,
		Password:        cfg.RabbitMq.Password,
		Host:            cfg.RabbitMq.Host,
		Port:            cfg.RabbitMq.Port,
		Vhost:           cfg.RabbitMq.Vhost,
		ReconnectSecond: cfg.RabbitMq.ReconnectSecond,
	}
	bufferSize := cfg.Source.BufferSize
	numberInstants := cfg.Source.NumberInstants
	maxRetryPolicy := cfg.Source.MaxRetry
	testCount := 100

	err := otela.Configure(wg, done, otela.Config{Exporter: otela.ExporterConsole})
	if err != nil {
		panic(err.Error())
	}
	inputAdapter := rbbitmqchannel.New(done, wg, dbConfig)
	err = inputAdapter.NewChannel(cfg.Source.NewSourceEventName, channel.InputOnlyMode, bufferSize, numberInstants, maxRetryPolicy)
	if err != nil {
		panic(err.Error())
	}
	inputChannel, err := inputAdapter.GetInputChannel(cfg.Source.NewSourceEventName)
	if err != nil {
		panic(err.Error())
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for messageID := 0; messageID < testCount; messageID++ {
			msg := encoder.EncodeNewSourceEvent(&source.NewSourceEvent{
				ProjectId: uuid.New().String(),
				OwnerId:   uuid.New().String(),
				WriteKey:  uuid.New().String(),
			})
			inputChannel <- []byte(msg)

		}
	}()
	wg.Wait()
}
