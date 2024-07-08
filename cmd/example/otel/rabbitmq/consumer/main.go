package main

import (
	"encoding/json"
	"fmt"
	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/pkg/channel"
	rbbitmqchannel "github.com/ormushq/ormus/pkg/channel/adapter/rabbitmq"
	"os"
	"os/signal"
	"sync"
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
		ServiceName:        "Test-Rabbitmq-Consumer",
		EnableMetricExpose: false,
		Exporter:           otela.ExporterGrpc,
	}
	err := otela.Configure(wg, done, cfg)
	if err != nil {
		panic(err)
	}

	port := 5672
	reconnectSecond := 10
	channelAdapter := rbbitmqchannel.New(done, wg, dconfig.RabbitMQConsumerConnection{
		User:            "guest",
		Password:        "guest",
		Host:            "rabbitmq",
		Port:            port,
		Vhost:           "",
		ReconnectSecond: reconnectSecond,
	})
	bufferSize := 100
	numberInstants := 1
	maxRetryPolicy := 10
	err = channelAdapter.NewChannel("test", channel.OutputOnly, bufferSize, numberInstants, maxRetryPolicy)
	if err != nil {
		panic(err)
	}
	outputChannel, err := channelAdapter.GetOutputChannel("test")
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-outputChannel:
				func() {
					fmt.Printf("recived message : %s\n", msg.Body)
					var decode MyMessage
					err = json.Unmarshal(msg.Body, &decode)
					if err != nil {
						panic(err)
					}
					fmt.Println(decode)
					ctx := otela.GetContextFromCarrier(decode.Carrier)

					tracer := otela.NewTracer("test-tracer")
					_, span := tracer.Start(ctx, "test-span-after-rabbit")
					defer span.End()

					span.AddEvent("task-ended")
					err = msg.Ack()
					if err != nil {
						panic(err)
					}
				}()

			case <-done:
				return
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	close(done)
	wg.Wait()
}
