package simple

import (
	"fmt"
	"github.com/ormushq/ormus/destination/channel"
	"sync"
	"time"
)

type simpleChannel struct {
	wg             *sync.WaitGroup
	done           <-chan bool
	mode           channel.Mode
	inputChannel   chan []byte
	outputChannel  chan channel.Message
	numberInstants int
}

func newChannel(done <-chan bool, wg *sync.WaitGroup, mode channel.Mode,
	bufferSize, numberInstants int) *simpleChannel {
	sc := &simpleChannel{
		done:           done,
		wg:             wg,
		mode:           mode,
		numberInstants: numberInstants,
		inputChannel:   make(chan []byte, bufferSize),
		outputChannel:  make(chan channel.Message, bufferSize),
	}
	sc.startConsume()

	return sc
}
func (sc simpleChannel) GetMode() channel.Mode {

	return sc.mode
}
func (sc simpleChannel) GetInputChannel() chan<- []byte {

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
		case <-time.After(5 * time.Second):
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
		sc.outputChannel <- channel.Message{
			Body: msg,
			Ack: func(multiple bool) error {
				c <- true

				return nil
			},
		}
	}(msg, c)
}
