package main

import (
	"sync"

	"github.com/google/uuid"
	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/channel/adapter/rabbitmqchannel"
	"github.com/ormushq/ormus/pkg/encoder"
)

func main() {
	cfg := config.C()
	done := make(chan bool)
	wg := &sync.WaitGroup{}
	bufferSize := cfg.Source.BufferSize
	maxRetryPolicy := cfg.Source.MaxRetry
	testCount := 100

	err := otela.Configure(wg, done, otela.Config{Exporter: otela.ExporterConsole})
	if err != nil {
		panic(err.Error())
	}
	inputAdapter := rabbitmqchannel.New(done, wg, cfg.RabbitMq)
	err = inputAdapter.NewChannel(cfg.Source.NewSourceEventName, channel.InputOnlyMode, bufferSize, maxRetryPolicy)
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
