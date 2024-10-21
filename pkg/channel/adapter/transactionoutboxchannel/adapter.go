package transactionoutboxchannel

import (
	"fmt"
	"github.com/ormushq/ormus/pkg/channel"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Storage interface {
	StoreMessage(channelName string, body []byte) error
	GetMessage(channelName string) (body []byte, ack func() error, err error)
}

type ChannelAdapter struct {
	wg            *sync.WaitGroup
	done          <-chan bool
	adapter       channel.Adapter
	storage       Storage
	logger        *zap.Logger
	inputChannels map[string]chan channel.Message
	cfg           Config
}

type Config struct {
	SendIntervalSec     int
	NumberMsgSendPerRun int
}

func New(done <-chan bool, wg *sync.WaitGroup, logger *zap.Logger, cfg Config, adapter channel.Adapter, storage Storage) *ChannelAdapter {
	return &ChannelAdapter{
		done:          done,
		wg:            wg,
		logger:        logger,
		adapter:       adapter,
		storage:       storage,
		inputChannels: make(map[string]chan channel.Message),
		cfg:           cfg,
	}
}

func (ca *ChannelAdapter) NewChannel(name string, mode channel.Mode, bufferSize, numberInstants, maxRetryPolicy int) error {
	ca.inputChannels[name] = make(chan channel.Message, bufferSize)
	err := ca.start(name)
	if err != nil {
		return err
	}

	return ca.adapter.NewChannel(name, mode, bufferSize, numberInstants, maxRetryPolicy)
}

func (ca *ChannelAdapter) GetInputChannel(name string) (chan<- channel.Message, error) {
	if c, ok := ca.inputChannels[name]; ok {
		return c, nil
	}

	return nil, fmt.Errorf("channel %s not found", name)
}

func (ca *ChannelAdapter) GetOutputChannel(name string) (<-chan channel.Message, error) {
	return ca.adapter.GetOutputChannel(name)
}

func (ca *ChannelAdapter) GetMode(name string) (channel.Mode, error) {
	return ca.adapter.GetMode(name)
}

func (ca *ChannelAdapter) start(name string) error {
	ca.wg.Add(1)
	go func() {
		defer ca.wg.Done()
		for {
			select {
			case <-ca.done:
				return
			case msg := <-ca.inputChannels[name]:
				err := ca.storage.StoreMessage(name, msg.Body)
				if err != nil {
					ca.logger.Error(err.Error())
				}
			}
		}
	}()

	inputChannel, err := ca.adapter.GetInputChannel(name)
	if err != nil {
		return err
	}
	ca.wg.Add(1)
	go func() {
		defer ca.wg.Done()
		for {
			select {
			case <-ca.done:
				return
			default:
				for i := 0; i < ca.cfg.NumberMsgSendPerRun; i++ {
					msg, ack, err := ca.storage.GetMessage(name)
					if err != nil {
						ca.logger.Error(err.Error())
						continue
					}
					inputChannel <- channel.Message{
						Body: msg,
						Ack:  ack,
					}
				}

				time.Sleep(time.Duration(ca.cfg.SendIntervalSec) * time.Second)
			}
		}
	}()

	return nil
}
