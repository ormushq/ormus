package rbbitmqchannel

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/errmsg"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ChannelAdapter struct {
	wg                       *sync.WaitGroup
	done                     <-chan bool
	channels                 map[string]*rabbitmqChannel
	config                   dconfig.RabbitMQConsumerConnection
	rabbitmq                 *Rabbitmq
	rabbitmqConnectionClosed chan bool
}

type Rabbitmq struct {
	connection *amqp.Connection
	cond       *sync.Cond
}

const loggerGroupName = "pkg.channel.rabbit"

func New(done <-chan bool, wg *sync.WaitGroup, config dconfig.RabbitMQConsumerConnection) *ChannelAdapter {
	return NewWithContext(context.Background(), done, wg, config)
}

func NewWithContext(ctx context.Context, done <-chan bool, wg *sync.WaitGroup, config dconfig.RabbitMQConsumerConnection) *ChannelAdapter {
	tracer := otela.NewTracer("rbbitmqchannel")
	_, span := tracer.Start(ctx, "rbbitmqchannel@NewWithContext")
	defer span.End()

	cond := sync.NewCond(&sync.Mutex{})
	rabbitmq := Rabbitmq{
		cond:       cond,
		connection: &amqp.Connection{},
	}
	c := &ChannelAdapter{
		done:                     done,
		wg:                       wg,
		config:                   config,
		rabbitmq:                 &rabbitmq,
		channels:                 make(map[string]*rabbitmqChannel),
		rabbitmqConnectionClosed: make(chan bool),
	}

	for {
		err := c.connect()
		time.Sleep(time.Second * time.Duration(config.ReconnectSecond))

		if err == nil {
			break
		}
		logger.WithGroup(loggerGroupName).Error("rabbitmq connection failed",
			slog.String("error", err.Error()))
	}
	span.AddEvent("connection-established")

	return c
}

func (ca *ChannelAdapter) connect() error {
	ca.rabbitmq.cond.L.Lock()
	defer ca.rabbitmq.cond.L.Unlock()

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		ca.config.User, ca.config.Password, ca.config.Host,
		ca.config.Port, ca.config.Vhost))
	if err != nil {
		return err
	}
	ca.rabbitmq.connection = conn
	logger.WithGroup(loggerGroupName).Debug("connected to rabbitmq server")
	ca.rabbitmq.cond.Broadcast()

	ca.waitForConnectionClose()

	return nil
}

func (ca *ChannelAdapter) waitForConnectionClose() {
	ca.wg.Add(1)
	go func() {
		defer ca.wg.Done()

		connectionClosedChannel := make(chan *amqp.Error)
		ca.rabbitmq.connection.NotifyClose(connectionClosedChannel)

		for {
			select {
			case <-ca.done:
				return
			case err := <-connectionClosedChannel:
				logger.WithGroup(loggerGroupName).Error("connection closed",
					slog.String("error", err.Error()))
				for {
					e := ca.connect()
					time.Sleep(time.Second * time.Duration(ca.config.ReconnectSecond))

					if e == nil {
						break
					}
					logger.WithGroup(loggerGroupName).Error("Connection failed to rabbitmq",
						slog.String("error", e.Error()))

				}

				return
			}
		}
	}()
}

func (ca *ChannelAdapter) NewChannel(name string, mode channel.Mode, bufferSize, numberInstants, maxRetryPolicy int) error {
	return ca.NewChannelWithContext(context.Background(), name, mode, bufferSize, numberInstants, maxRetryPolicy)
}

func (ca *ChannelAdapter) NewChannelWithContext(ctx context.Context, name string, mode channel.Mode, bufferSize, numberInstants, maxRetryPolicy int) error {
	tracer := otela.NewTracer("rbbitmqchannel")
	_, span := tracer.Start(ctx, "rbbitmqchannel@NewChannel")
	defer span.End()

	ch, err := newChannelWithContext(
		ctx,
		ca.done,
		ca.wg,
		rabbitmqChannelParams{
			mode:           mode,
			rabbitmq:       ca.rabbitmq,
			exchange:       name + "-exchange",
			queue:          name + "-queue",
			bufferSize:     bufferSize,
			numberInstants: numberInstants,
			maxRetryPolicy: maxRetryPolicy,
		})
	if err != nil {
		return err
	}
	span.AddEvent("channel-created")
	ca.channels[name] = ch

	return nil
}

func (ca *ChannelAdapter) GetInputChannel(name string) (chan<- channel.Message, error) {
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

func WaitForConnection(rabbitmq *Rabbitmq) {
	rabbitmq.cond.L.Lock()
	defer rabbitmq.cond.L.Unlock()
	for rabbitmq.connection.IsClosed() {
		rabbitmq.cond.Wait()
	}
}
