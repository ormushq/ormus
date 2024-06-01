package main

import (
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/channel/adapter/rabbitmq"
	"log"
	"sync"
	"time"
)

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	channelName := "test"

	maxRetryPolicy := 5
	numberInstants := 5
	bufferSize := 5
	inputChannelAdapter := rbbitmqchannel.New(done, &wg, config.C().Destination.RabbitMQConsumerConnection)
	inputChannelAdapter.NewChannel(channelName, channel.InputOnlyMode, bufferSize, numberInstants, maxRetryPolicy)

	outputChannelAdapter := rbbitmqchannel.New(done, &wg, config.C().Destination.RabbitMQConsumerConnection)
	outputChannelAdapter.NewChannel(channelName, channel.OutputOnly, bufferSize, numberInstants, maxRetryPolicy)

	inputChannel, err := inputChannelAdapter.GetInputChannel(channelName)
	if err != nil {
		log.Fatal(err)
	}
	outputChannel, err := outputChannelAdapter.GetOutputChannel(channelName)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go func() {
		for msg := range outputChannel {
			err := msg.Ack()
			if err != nil {
				logger.L().Error(err.Error())

			}
		}
	}()

	inputChannel <- []byte("Hello form input channel" + time.Now().UTC().String())
	wg.Add(1)
	go func() {
		for {
			inputChannel <- []byte("Hello form input channel" + time.Now().UTC().String())
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
