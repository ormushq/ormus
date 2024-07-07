package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/pkg/channel"
	rbbitmqchannel "github.com/ormushq/ormus/pkg/channel/adapter/rabbitmq"
	"os"
	"os/signal"
	"sync"
	"time"
)

type MyMessage struct {
	Name    string
	Carrier map[string]string
}

func main() {
	wg := &sync.WaitGroup{}
	done := make(chan bool)

	cfg := otela.Config{
		Endpoint:           "otel_collector:4317",
		ServiceName:        "Test-Rabbitmq-Publisher",
		EnableMetricExpose: false,
	}
	err := otela.Configure(wg, done, cfg)
	if err != nil {
		panic(err)
	}

	channelAdapter := rbbitmqchannel.New(done, wg, dconfig.RabbitMQConsumerConnection{
		User:            "guest",
		Password:        "guest",
		Host:            "rabbitmq",
		Port:            5672,
		Vhost:           "",
		ReconnectSecond: 10,
	})

	err = channelAdapter.NewChannel("test", channel.InputOnlyMode, 100, 1, 10)
	if err != nil {
		panic(err)
	}

	inputChannel, err := channelAdapter.GetInputChannel("test")
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			default:
				func() {
					tracer := otela.NewTracer("test-tracer")
					ctx, span := tracer.Start(context.Background(), "test-span-before-rabbit")
					defer span.End()

					span.AddEvent("task-created")
					carrier := otela.GetCarrierFromContext(ctx)

					msg := MyMessage{
						Name:    "Mohsen",
						Carrier: carrier,
					}
					fmt.Printf("ctx %+v\n", span.SpanContext().SpanID().String())
					encode, err := json.Marshal(msg)
					if err != nil {
						panic(err)
					}
					fmt.Printf("encode message %s\n", encode)

					inputChannel <- encode

					time.Sleep(time.Second * 5)
				}()

			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	close(done)
	wg.Wait()
}
