package simple

import (
	"fmt"
	"sync"

	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/errmsg"
)

type ChannelAdapter struct {
	wg       *sync.WaitGroup
	done     <-chan bool
	channels map[string]*simpleChannel
}

func New(done <-chan bool, wg *sync.WaitGroup) *ChannelAdapter {
	return &ChannelAdapter{
		channels: make(map[string]*simpleChannel),
		done:     done,
		wg:       wg,
	}
}

func (ca *ChannelAdapter) NewChannel(name string, mode channel.Mode, bufferSize, numberInstants, maxRetryPolicy int) error {
	ca.channels[name] = newChannel(ca.done, ca.wg, mode, bufferSize, numberInstants, maxRetryPolicy)

	return nil
}

func (ca *ChannelAdapter) GetInputChannel(name string) (chan<- []byte, error) {
	if c, ok := ca.channels[name]; ok {
		return c.GetInputChannel(), nil
	}

	return nil, fmt.Errorf(errmsg.ErrChannelNotFound, name)
}

func (ca *ChannelAdapter) GetOutputChannel(name string) (<-chan channel.Message, error) {
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
