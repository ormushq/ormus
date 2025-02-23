package main

import (
	"sync"

	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/contract/go/destination"
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
	testCount := 1

	err := otela.Configure(wg, done, otela.Config{Exporter: otela.ExporterConsole})
	if err != nil {
		panic(err.Error())
	}
	inputAdapter := rabbitmqchannel.New(done, wg, cfg.RabbitMq)
	err = inputAdapter.NewChannel(cfg.Source.UndeliveredEventsQueueName, channel.InputOnlyMode, bufferSize, maxRetryPolicy)
	if err != nil {
		panic(err.Error())
	}
	inputChannel, err := inputAdapter.GetInputChannel(cfg.Source.UndeliveredEventsQueueName)
	if err != nil {
		panic(err.Error())
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for messageID := 0; messageID < testCount; messageID++ {
			msg := encoder.EncodeProcessedEvent(&destination.DeliveredEventsList{
				Events: []*destination.DeliveredEvent{
					{
						MessageId: "d5aacd53-f866-4406-8e2f-d6f1dbc96975",
					},
				},
			})
			inputChannel <- []byte(msg)

		}
	}()
	wg.Wait()
}
