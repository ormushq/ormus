package eventmanager

import "github.com/ormushq/ormus/pkg/channel"

func NewCreateChannelFunc(ca channel.Adapter, mod channel.Mode, bufferSize, maxRetryPolicy int) CreateChannelFunc {
	return func(channelName string) error {
		return ca.NewChannel(channelName, mod, bufferSize, maxRetryPolicy)
	}
}
