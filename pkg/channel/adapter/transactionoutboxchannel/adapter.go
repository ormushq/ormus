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
	GetMessage(channelName string, limit int) ([][]byte, error)
}

type ChannelAdapter struct {
	wg            *sync.WaitGroup
	done          <-chan bool
	adapter       channel.Adapter
	storage       Storage
	logger        *zap.Logger
	inputChannels map[string]chan []byte
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
		inputChannels: make(map[string]chan []byte),
		cfg:           cfg,
	}
}

func (ca *ChannelAdapter) NewChannel(name string, mode channel.Mode, bufferSize, numberInstants, maxRetryPolicy int) error {
	ca.inputChannels[name] = make(chan []byte, bufferSize)
	err := ca.start(name)
	if err != nil {
		return err
	}

	return ca.adapter.NewChannel(name, mode, bufferSize, numberInstants, maxRetryPolicy)
}

func (ca *ChannelAdapter) GetInputChannel(name string) (chan<- []byte, error) {
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
				err := ca.storage.StoreMessage(name, msg)
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
				result, err := ca.storage.GetMessage(name, ca.cfg.NumberMsgSendPerRun)
				if err != nil {
					ca.logger.Error(err.Error())
					continue
				}
				for _, msg := range result {
					inputChannel <- msg
				}
				time.Sleep(time.Duration(ca.cfg.SendIntervalSec) * time.Second)
			}
		}
	}()

	return nil
}
