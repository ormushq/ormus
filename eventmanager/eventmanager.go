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
	adapter channel.Adapter
	wg      *sync.WaitGroup
	done    <-chan bool
	events  map[internalevent.EventName]CreateChannelFunc
}

var once = make(map[channel.Adapter]map[string]*sync.Once)

func New(wg *sync.WaitGroup, done <-chan bool, adapter channel.Adapter, events map[internalevent.EventName]CreateChannelFunc) (Manager, error) {
	manager := Manager{
		wg:      wg,
		done:    done,
		adapter: adapter,
		events:  events,
	}
	for eventType, createChannelFunc := range events {
		err := manager.createChannel(createChannelFunc, eventType)
		if err != nil {
			return Manager{}, err
		}
	}

	return manager, nil
}

func (h Manager) Publish(msg *internalevent.Event) error {
	channelName, err := h.checkChannel(msg.EventName)
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

func (h Manager) checkChannel(eventName internalevent.EventName) (string, error) {
	_, ok := h.events[eventName]
	if !ok {
		return "", fmt.Errorf("unknown event name %s", eventName)
	}

	return h.getChannelName(eventName)
}

func (h Manager) Consume(eventTypes ...internalevent.EventName) (<-chan EventMessage, error) {
	if len(eventTypes) == 0 {
		return nil, fmt.Errorf("no event types provided")
	}

	chs := make([]<-chan channel.Message, 0)
	for _, eventType := range eventTypes {
		channelName, err := h.checkChannel(eventType)
		if err != nil {
			return nil, err
		}

		ch, err := h.adapter.GetOutputChannel(channelName)
		if err != nil {
			return nil, err
		}
		chs = append(chs, ch)
	}

	return h.convertChannel(chs...), nil
}

func (h Manager) createChannel(createChannelFunc CreateChannelFunc, eventType internalevent.EventName) error {
	channelName, err := h.getChannelName(eventType)
	if err != nil {
		return err
	}
	_, ok := once[h.adapter]
	if !ok {
		once[h.adapter] = make(map[string]*sync.Once)
	}
	_, ok = once[h.adapter][channelName]
	if !ok {
		once[h.adapter][channelName] = &sync.Once{}
	}
	once[h.adapter][channelName].Do(func() {
		err = createChannelFunc(channelName)
	})

	return err
}

func (h Manager) convertChannel(chs ...<-chan channel.Message) <-chan EventMessage {
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
					var ev internalevent.Event
					err := proto.Unmarshal(msg.Body, &ev)
					if err != nil {
						logger.LogError(err)

						continue
					}
					newCh <- NewEventMessage(&ev, msg.Ack)

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
