package channel

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskdelivery/param"
	event2 "github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/protobufmapper"
)

type Converter struct {
	done                         <-chan bool
	wg                           *sync.WaitGroup
	processedEventOutputChannels map[<-chan []byte]<-chan event2.ProcessedEvent
	processedEventInputChannels  map[chan<- []byte]chan<- event2.ProcessedEvent
	taskInputChannels            map[chan<- []byte]chan<- taskentity.Task
	taskOutputChannels           map[<-chan []byte]<-chan taskentity.Task
	deliveryTaskInputChannels    map[chan<- []byte]chan<- param.DeliveryTaskResponse
	deliveryTaskOutputChannels   map[<-chan []byte]<-chan param.DeliveryTaskResponse
}

func NewConverter(done <-chan bool, wg *sync.WaitGroup) *Converter {
	return &Converter{
		done:                         done,
		wg:                           wg,
		processedEventOutputChannels: make(map[<-chan []byte]<-chan event2.ProcessedEvent),
		processedEventInputChannels:  make(map[chan<- []byte]chan<- event2.ProcessedEvent),
		taskInputChannels:            make(map[chan<- []byte]chan<- taskentity.Task),
		taskOutputChannels:           make(map[<-chan []byte]<-chan taskentity.Task),
		deliveryTaskInputChannels:    make(map[chan<- []byte]chan<- param.DeliveryTaskResponse),
		deliveryTaskOutputChannels:   make(map[<-chan []byte]<-chan param.DeliveryTaskResponse),
	}
}

func (c *Converter) ConvertToOutputProcessedEventChannel(originalChannel <-chan []byte, bufferSize int) <-chan event2.ProcessedEvent {
	if outputChannel, ok := c.processedEventOutputChannels[originalChannel]; ok {
		return outputChannel
	}
	outputChannel := make(chan event2.ProcessedEvent, bufferSize)
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			select {
			case <-c.done:

				return
			case msg := <-originalChannel:
				c.wg.Add(1)
				go func() {
					defer c.wg.Done()
					pe, uErr := taskentity.ProtoUnmarshalBytesToProcessedEvnet(msg)
					if uErr != nil {
						logger.L().Debug(string(msg))
						slog.Error(fmt.Sprintf("Failed to convert bytes to processed events: %v", uErr))

						return
					}
					e := protobufmapper.MapProcessedEventFromProtobuf(pe)

					outputChannel <- e
				}()

			}
		}
	}()
	c.processedEventOutputChannels[originalChannel] = outputChannel

	return outputChannel
}

func (c *Converter) ConvertToOutputTaskChannel(originalChannel <-chan []byte, bufferSize int) <-chan taskentity.Task {
	if outputChannel, ok := c.taskOutputChannels[originalChannel]; ok {
		return outputChannel
	}
	outputChannel := make(chan taskentity.Task, bufferSize)
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			select {
			case <-c.done:

				return
			case msg := <-originalChannel:
				c.wg.Add(1)
				go func() {
					defer c.wg.Done()
					e, uErr := taskentity.UnmarshalBytesToTask(msg)
					if uErr != nil {
						slog.Error(fmt.Sprintf("Failed to convert bytes to processed events: %v", uErr))

						return
					}
					outputChannel <- e
				}()

			}
		}
	}()
	c.taskOutputChannels[originalChannel] = outputChannel

	return outputChannel
}

func (c *Converter) ConvertToInputTaskChannel(originalChannel chan<- []byte, bufferSize int) chan<- taskentity.Task {
	if inputChannel, ok := c.taskInputChannels[originalChannel]; ok {
		return inputChannel
	}
	inputChannel := make(chan taskentity.Task, bufferSize)
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			select {
			case <-c.done:

				return
			case task := <-inputChannel:
				c.wg.Add(1)
				go func() {
					defer c.wg.Done()
					e, uErr := json.Marshal(task)
					if uErr != nil {
						slog.Error(fmt.Sprintf("Failed to convert task tp bytes : %v", uErr))

						return
					}
					originalChannel <- e
				}()

			}
		}
	}()
	c.taskInputChannels[originalChannel] = inputChannel

	return inputChannel
}

func (c *Converter) ConvertToInputDeliveryTaskChannel(originalChannel chan<- []byte, bufferSize int) chan<- param.DeliveryTaskResponse {
	if inputChannel, ok := c.deliveryTaskInputChannels[originalChannel]; ok {
		return inputChannel
	}
	inputChannel := make(chan param.DeliveryTaskResponse, bufferSize)
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			select {
			case <-c.done:

				return
			case deliveryTask := <-inputChannel:
				c.wg.Add(1)
				go func() {
					defer c.wg.Done()
					e, uErr := json.Marshal(deliveryTask)
					if uErr != nil {
						slog.Error(fmt.Sprintf("Failed to convert deliverytask to byte: %v", uErr))

						return
					}
					originalChannel <- e
				}()

			}
		}
	}()
	c.deliveryTaskInputChannels[originalChannel] = inputChannel

	return inputChannel
}
