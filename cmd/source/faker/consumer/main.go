package main

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/config"
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
	eventName := cfg.Source.NewSourceEventName
	err := otela.Configure(wg, done, otela.Config{Exporter: otela.ExporterConsole})
	if err != nil {
		panic(err.Error())
	}

	outputAdapter := rbbitmqchannel.New(done, wg, dbConfig)
	err = outputAdapter.NewChannel(eventName, channel.OutputOnly, bufferSize, numberInstants, maxRetryPolicy)
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
			m := encoder.DecodeNewSourceEvent(string(msg.Body))
			log.Info(m.WriteKey)

			if err := msg.Ack(); err != nil {
				panic(err)
			}
		}
	}()
	wg.Wait()
}
