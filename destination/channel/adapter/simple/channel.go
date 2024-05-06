package simple

import (
	"fmt"
	"github.com/ormushq/ormus/destination/channel"
	"sync"
)

type simpleChannel struct {
	wg             *sync.WaitGroup
	done           <-chan bool
	mode           channel.Mode
	inputChannel   chan []byte
	outputChannel  chan []byte
	numberInstants int
}

func newChannel(done <-chan bool, wg *sync.WaitGroup, mode channel.Mode, bufferSize int, numberInstants int) *simpleChannel {
	sc := &simpleChannel{
		done:           done,
		wg:             wg,
		mode:           mode,
		numberInstants: numberInstants,
		inputChannel:   make(chan []byte, bufferSize),
		outputChannel:  make(chan []byte, bufferSize),
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
func (sc simpleChannel) GetOutputChannel() <-chan []byte {
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
					sc.outputChannel <- msg
				}
			}
		}()
	}

}
