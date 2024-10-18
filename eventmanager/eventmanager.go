package eventmanager

import (
	"fmt"
	"sync"

	"github.com/ormushq/ormus/contract/go/internalevent"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
	"google.golang.org/protobuf/proto"
)

type RegisterEventChannel func(channelName string) error

type Manager struct {
	adapter           channel.Adapter
	wg                *sync.WaitGroup
	done              <-chan bool
	createChannelFunc func(channelName string) error
}

var once = make(map[channel.Adapter]map[string]*sync.Once)

func New(wg *sync.WaitGroup, done <-chan bool, adapter channel.Adapter, createChannelFunc func(channelName string) error) Manager {
	return Manager{
		wg:                wg,
		done:              done,
		adapter:           adapter,
		createChannelFunc: createChannelFunc,
	}
}

func (h Manager) Publish(msg *internalevent.Event) error {
	channelName, err := h.getChannelName(msg.EventName)
	if err != nil {
		return err
	}

	err = h.createChannel(channelName)
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

func (h Manager) Consume(eventName internalevent.EventName, eventNames ...internalevent.EventName) (<-chan EventMessage, error) {
	eventNames = append(eventNames, eventName)
	channelNames := make([]string, 0)
	for _, eventName := range eventNames {
		channelName, err := h.getChannelName(eventName)
		if err != nil {
			return nil, err
		}
		channelNames = append(channelNames, channelName)
	}
	chs := make([]<-chan channel.Message, 0)
	for _, channelName := range channelNames {
		err := h.createChannel(channelName)
		if err != nil {
			return nil, err
		}
		ch, err := h.adapter.GetOutputChannel(channelName)
		if err != nil {
			return nil, err
		}
		chs = append(chs, ch)
	}

	return h.covertChannel(chs...), nil
}

func (h Manager) createChannel(channelName string) error {
	var err error
	once[h.adapter][channelName].Do(func() {
		err = h.createChannelFunc(channelName)
	})

	return err
}

func (h Manager) covertChannel(chs ...<-chan channel.Message) <-chan EventMessage {
	newCh := make(chan EventMessage, len(chs))

	for _, ch := range chs {
		h.wg.Add(1)
		go func(ch <-chan channel.Message) {
			defer h.wg.Done()
			for {
				select {
				case <-h.done:
					return
				case msg := <-ch:
					var ev *internalevent.Event
					err := proto.Unmarshal(msg.Body, ev)
					if err != nil {
						logger.LogError(err)

						continue
					}
					newCh <- NewEventMessage(ev, msg.Ack)

				}
			}
		}(ch)
	}

	return newCh
}

func (h Manager) getChannelName(eventName internalevent.EventName) (string, error) {
	switch eventName {
	case internalevent.EventName_EVENT_NAME_USER_CREATED:
		return EventNameUserCreated, nil
	case internalevent.EventName_EVENT_NAME_PROJECT_CREATED:
		return EventNameProjectCreated, nil
	case internalevent.EventName_EVENT_NAME_WRITE_KEY_GENERATED:
		return EventNameWriteKeyGenerated, nil
	case internalevent.EventName_EVENT_NAME_TASK_CREATED:
		return EventNameTaskCreated, nil
	}

	return "", fmt.Errorf("unsupported message type: %s", eventName)
}
