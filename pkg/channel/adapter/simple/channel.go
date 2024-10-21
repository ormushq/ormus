package simple

import (
	"sync"
	"time"

	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
)

type simpleChannel struct {
	wg             *sync.WaitGroup
	done           <-chan bool
	mode           channel.Mode
	inputChannel   chan channel.Message
	outputChannel  chan channel.Message
	numberInstants int
	maxRetryPolicy int
}

const timeForCallAgainDuration = 10

func newChannel(done <-chan bool, wg *sync.WaitGroup, mode channel.Mode,
	bufferSize, numberInstants, maxRetryPolicy int,
) *simpleChannel {
	sc := &simpleChannel{
		done:           done,
		wg:             wg,
		mode:           mode,
		numberInstants: numberInstants,
		maxRetryPolicy: maxRetryPolicy,
		inputChannel:   make(chan channel.Message, bufferSize),
		outputChannel:  make(chan channel.Message, bufferSize),
	}
	sc.startConsume()

	return sc
}

func (sc simpleChannel) GetMode() channel.Mode {
	return sc.mode
}

func (sc simpleChannel) GetInputChannel() chan<- channel.Message {
	return sc.inputChannel
}

func (sc simpleChannel) GetOutputChannel() <-chan channel.Message {
	return sc.outputChannel
}

func (sc simpleChannel) startConsume() {
	for i := 0; i < sc.numberInstants; i++ {
		sc.wg.Add(1)
		go func() {
			defer sc.wg.Done()
			for {
				select {
				case <-sc.done:

					return
				case msg := <-sc.inputChannel:
					logger.WithGroup(loggerGroupName).Debug("New message received in simple/adapter.go ca.inputChannel")
					sc.startDelivery(msg)
				}
			}
		}()
	}
}

func (sc simpleChannel) startDelivery(msg channel.Message) {
	sc.wg.Add(1)
	ackChan := make(chan bool)
	go func() {
		defer sc.wg.Done()
		go sc.publishMessage(msg, ackChan)
		select {
		case <-time.After(time.Second * timeForCallAgainDuration):
			sc.startDelivery(msg)
		case <-ackChan:

			return
		}
	}()
}

func (sc simpleChannel) publishMessage(msg channel.Message, c chan<- bool) {
	sc.wg.Add(1)
	go func(msg channel.Message, c chan<- bool) {
		defer sc.wg.Done()
		sc.outputChannel <- channel.Message{
			Body: msg.Body,
			Ack: func() error {
				c <- true

				return nil
			},
		}
	}(msg, c)
}
