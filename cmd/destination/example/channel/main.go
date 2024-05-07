package main

import (
	"fmt"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/destination/channel"
	rbbitmqchannel "github.com/ormushq/ormus/destination/channel/adapter/rabbitmq"
	"sync"
	"time"
)

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	channelName := "test"

	inputChannelAdapter := rbbitmqchannel.New(done, &wg, config.C().Destination.RabbitMQConsumerConnection)
	inputChannelAdapter.NewChannel(channelName, channel.InputOnlyMode, 100, 5)

	outputChannelAdapter := rbbitmqchannel.New(done, &wg, config.C().Destination.RabbitMQConsumerConnection)
	outputChannelAdapter.NewChannel(channelName, channel.OutputOnly, 100, 5)

	outputChannel, _ := outputChannelAdapter.GetOutputChannel(channelName)
	inputChannel, _ := inputChannelAdapter.GetInputChannel(channelName)

	wg.Add(1)
	go func() {
		for {
			select {
			case msg := <-outputChannel:
				fmt.Print(string(msg))
			}
		}
	}()

	wg.Add(1)
	go func() {
		for {
			inputChannel <- []byte("Hello form input channel" + time.Now().UTC().String())
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
