package main

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/config"
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
	eventName := cfg.Source.NewEventQueueName
	err := otela.Configure(wg, done, otela.Config{Exporter: otela.ExporterConsole})
	if err != nil {
		panic(err.Error())
	}

	outputAdapter := rabbitmqchannel.New(done, wg, cfg.RabbitMq)
	err = outputAdapter.NewChannel(eventName, channel.OutputOnly, bufferSize, maxRetryPolicy)
	if err != nil {
		panic(err)
	}
	outputChannel, err := outputAdapter.GetOutputChannel(eventName)
	if err != nil {
		panic(err)
	}

	wg.Add(1)

	go func() {
		defer wg.Done()
		for msg := range outputChannel {
			m := encoder.DecodeNewEvent(string(msg.Body))
			log.Info(m)

			if err := msg.Ack(); err != nil {
				panic(err)
			}
		}
	}()
	wg.Wait()
}
