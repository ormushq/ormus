package simple

import (
	"fmt"
	channel2 "github.com/ormushq/ormus/pkg/channel"
	"sync"
	"time"
)

type simpleChannel struct {
	wg             *sync.WaitGroup
	done           <-chan bool
	mode           channel2.Mode
	inputChannel   chan []byte
	outputChannel  chan channel2.Message
	numberInstants int
}

const timeForCallAgainDuration = 10

func newChannel(done <-chan bool, wg *sync.WaitGroup, mode channel2.Mode, bufferSize, numberInstants int) *simpleChannel {
	sc := &simpleChannel{
		done:           done,
		wg:             wg,
		mode:           mode,
		numberInstants: numberInstants,
		inputChannel:   make(chan []byte, bufferSize),
		outputChannel:  make(chan channel2.Message, bufferSize),
	}
	sc.startConsume()

	return sc
}

func (sc simpleChannel) GetMode() channel2.Mode {
	return sc.mode
}

func (sc simpleChannel) GetInputChannel() chan<- []byte {
	return sc.inputChannel
}

func (sc simpleChannel) GetOutputChannel() <-chan channel2.Message {
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
					fmt.Println("New message received in simple/adapter.go ca.inputChannel", msg)
					sc.startDelivery(msg)
				}
			}
		}()
	}
}

func (sc simpleChannel) startDelivery(msg []byte) {
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

func (sc simpleChannel) publishMessage(msg []byte, c chan<- bool) {
	sc.wg.Add(1)
	go func(msg []byte, c chan<- bool) {
		defer sc.wg.Done()
		sc.outputChannel <- channel2.Message{
			Body: msg,
			Ack: func() error {
				c <- true

				return nil
			},
		}
	}(msg, c)
}
