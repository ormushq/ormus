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
	inputChannelAdapter.NewChannel(channelName, channel.InputOnlyMode, 0, 1)

	outputChannelAdapter := rbbitmqchannel.New(done, &wg, config.C().Destination.RabbitMQConsumerConnection)
	outputChannelAdapter.NewChannel(channelName, channel.OutputOnly, 0, 1)

	inputChannel, _ := inputChannelAdapter.GetInputChannel(channelName)
	outputChannel, _ := outputChannelAdapter.GetOutputChannel(channelName)

	wg.Add(1)
	go func() {
		for {
			select {
			case msg := <-outputChannel:
				//fmt.Println(string(msg.Body))
				err := msg.Ack(false)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	inputChannel <- []byte("Hello form input channel" + time.Now().UTC().String())
	wg.Add(1)
	go func() {
		for {
			//fmt.Println("Send date to input channel")
			inputChannel <- []byte("Hello form input channel" + time.Now().UTC().String())
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
