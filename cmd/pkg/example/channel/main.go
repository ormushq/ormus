package main

import (
	"log"
	"sync"
	"time"

	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
	rabitmqchannel "github.com/ormushq/ormus/pkg/channel/adapter/rabbitmq"
)

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	channelName := "test"

	maxRetryPolicy := 5
	bufferSize := 5
	inputChannelAdapter := rabitmqchannel.New(done, &wg, rabitmqchannel.Config{
		User:            "guest",
		Password:        "guest",
		Host:            "127.0.0.1",
		Port:            5672,
		Vhost:           "/",
		ReconnectSecond: 2,
	})
	err := inputChannelAdapter.NewChannel(channelName, channel.InputOnlyMode, bufferSize, maxRetryPolicy)
	if err != nil {
		log.Fatal(err)
	}

	outputChannelAdapter := rabitmqchannel.New(done, &wg, rabitmqchannel.Config{
		User:            "guest",
		Password:        "guest",
		Host:            "127.0.0.1",
		Port:            5672,
		Vhost:           "/",
		ReconnectSecond: 2,
	})
	err = outputChannelAdapter.NewChannel(channelName, channel.OutputOnly, bufferSize, maxRetryPolicy)
	if err != nil {
		log.Fatal(err)
	}

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
