package rbbitmqadapter

import (
	"fmt"
	"github.com/ormushq/ormus/destination/channel"
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/pkg/errmsg"
	"sync"
)

type ChannelAdapter struct {
	wg       *sync.WaitGroup
	done     <-chan bool
	channels map[string]*rabbitmqChannel
	config   dconfig.RabbitMQConsumerConnection
}

func New(done <-chan bool, wg *sync.WaitGroup, config dconfig.RabbitMQConsumerConnection) *ChannelAdapter {
	return &ChannelAdapter{
		done:     done,
		wg:       wg,
		config:   config,
		channels: make(map[string]*rabbitmqChannel),
	}
}

func (ca *ChannelAdapter) NewChannel(name string, mode channel.Mode, bufferSize int, numberInstants int) {
	ca.channels[name] = newChannel(ca.done, ca.wg, rabbitmqChannelParams{mode, ca.config, name + "-exchange",
		name + "-queue", bufferSize, numberInstants})
}

func (ca *ChannelAdapter) GetInputChannel(name string) (chan<- []byte, error) {
	if c, ok := ca.channels[name]; ok {
		return c.GetInputChannel(), nil
	}
	return nil, fmt.Errorf(errmsg.ErrChannelNotFound, name)
}
func (ca *ChannelAdapter) GetOutputChannel(name string) (<-chan []byte, error) {
	if c, ok := ca.channels[name]; ok {
		return c.GetOutputChannel(), nil
	}
	return nil, fmt.Errorf(errmsg.ErrChannelNotFound, name)
}
func (ca *ChannelAdapter) GetMode(name string) (channel.Mode, error) {
	if c, ok := ca.channels[name]; ok {
		return c.GetMode(), nil
	}
	return "", fmt.Errorf(errmsg.ErrChannelNotFound, name)
}
