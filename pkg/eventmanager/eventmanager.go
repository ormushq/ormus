package eventmanager

import (
	"fmt"
	"github.com/ormushq/ormus/contract/go/internalevent"
	"github.com/ormushq/ormus/pkg/channel"
	"sync"

	"google.golang.org/protobuf/proto"
)

type RegisterEventChannel func(channelName string) error

type Manager struct {
	adapter            channel.Adapter
	wg                 *sync.WaitGroup
	done               <-chan bool
	errorFunc          func(error)
	declareChannelFunc func(channelName string) error
}

func New(wg *sync.WaitGroup, done <-chan bool, errorFunc func(error), adapter channel.Adapter, declareChannelFunc func(channelName string) error) Manager {
	return Manager{
		wg:                 wg,
		done:               done,
		adapter:            adapter,
		errorFunc:          errorFunc,
		declareChannelFunc: declareChannelFunc,
	}
}

func (h Manager) Publish(msg *internalevent.Event) error {
	channelName, err := h.getChannelName(msg)
	if err != nil {
		return err
	}
	err = h.declareChannelFunc(channelName)
	if err != nil {
		return err
	}

	inputChannel, err := h.adapter.GetInputChannel(channelName)
	if err != nil {
		return fmt.Errorf("no input channel found for channel %s", channelName)
	}

	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	inputChannel <- body

	return nil
}

func (h Manager) Consume(msg *internalevent.Event) (<-chan EventMessage, error) {
	channelName, err := h.getChannelName(msg)
	if err != nil {
		return nil, err
	}
	ch, err := h.adapter.GetOutputChannel(channelName)
	if err != nil {
		return nil, err
	}

	return h.covertChannel(ch), nil
}

func (h Manager) covertChannel(ch <-chan channel.Message) <-chan EventMessage {
	newCh := make(chan EventMessage)

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()

		for {
			select {
			case <-h.done:
				return
			case msg := <-ch:
				var ev *internalevent.Event
				err := proto.Unmarshal(msg.Body, ev)
				if err != nil {
					h.errorFunc(err)

					continue
				}
				newCh <- NewEventMessage(ev, msg.Ack)

			}
		}
	}()

	return newCh
}

func (h Manager) getChannelName(msg *internalevent.Event) (string, error) {
	switch msg.EventName {
	case internalevent.EventName_EVENT_NAME_USER_CREATED:
		return EventNameUserCreated, nil
	case internalevent.EventName_EVENT_NAME_PROJECT_CREATED:
		return EventNameProjectCreated, nil
	case internalevent.EventName_EVENT_NAME_WRITE_KEY_GENERATED:
		return EventNameWriteKeyGenerated, nil
	case internalevent.EventName_EVENT_NAME_TASK_CREATED:
		return EventNameTaskCreated, nil
	}

	return "", fmt.Errorf("unsupported message type: %T", msg.Payload)
}
