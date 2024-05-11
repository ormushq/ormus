package rbbitmqchannel

import (
	"fmt"
	"github.com/ormushq/ormus/destination/channel"
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/pkg/errmsg"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

type ChannelAdapter struct {
	wg       *sync.WaitGroup
	done     <-chan bool
	channels map[string]*rabbitmqChannel
	config   dconfig.RabbitMQConsumerConnection
	rabbitmq *Rabbitmq
}
type Rabbitmq struct {
	connection *amqp.Connection
	cond       *sync.Cond
}

func New(done <-chan bool, wg *sync.WaitGroup, config dconfig.RabbitMQConsumerConnection) *ChannelAdapter {
	cond := sync.NewCond(&sync.Mutex{})
	rabbitmq := Rabbitmq{
		cond:       cond,
		connection: &amqp.Connection{},
	}
	fmt.Printf("Main rabbitmq object address %p \n", &rabbitmq)
	c := &ChannelAdapter{
		done:     done,
		wg:       wg,
		config:   config,
		rabbitmq: &rabbitmq,
		channels: make(map[string]*rabbitmqChannel),
	}

	for {
		err := c.connect()
		time.Sleep(time.Second * time.Duration(config.ReconnectSecond))
		failOnError(err, "rabbitmq connection failed")
		if err == nil {
			break
		}
	}

	return c
}
func (ca *ChannelAdapter) connect() error {
	ca.rabbitmq.cond.L.Lock()
	defer ca.rabbitmq.cond.L.Unlock()

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		ca.config.User, ca.config.Password, ca.config.Host,
		ca.config.Port, ca.config.Vhost))
	failOnError(err, "Failed to connect to rabbitmq server")
	if err != nil {
		return err
	}
	ca.rabbitmq.connection = conn
	fmt.Println("Connected to rabbitmq server")
	ca.rabbitmq.cond.Broadcast()

	ca.wg.Add(1)
	go func() {
		defer ca.wg.Done()
		for range ca.done {
			err = conn.Close()
			failOnError(err, "failed to close a connection")

			break
		}

	}()
	go ca.waitForConnectionClose()

	return nil
}
func (ca *ChannelAdapter) waitForConnectionClose() {
	connectionClosedChannel := make(chan *amqp.Error)
	ca.rabbitmq.connection.NotifyClose(connectionClosedChannel)

	for {
		select {
		case <-ca.done:
			return
		case err := <-connectionClosedChannel:
			fmt.Println("Connection closed")
			fmt.Println(err)
			for {
				e := ca.connect()
				time.Sleep(time.Second * time.Duration(ca.config.ReconnectSecond))
				failOnError(e, "Connection failed to rabbitmq")
				if e == nil {
					break
				}
			}

			return
		}
	}
}

func (ca *ChannelAdapter) NewChannel(name string, mode channel.Mode, bufferSize, numberInstants int) {
	ca.channels[name] = newChannel(
		ca.done,
		ca.wg,
		rabbitmqChannelParams{
			mode:           mode,
			rabbitmq:       ca.rabbitmq,
			exchange:       name + "-exchange",
			queue:          name + "-queue",
			bufferSize:     bufferSize,
			numberInstants: numberInstants,
		})
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

func WaitForConnection(rabbitmq *Rabbitmq) {
	fmt.Printf("the address in wait for connection %p \n", rabbitmq)

	rabbitmq.cond.L.Lock()
	defer rabbitmq.cond.L.Unlock()
	for rabbitmq.connection.IsClosed() {
		fmt.Println(rabbitmq.connection.IsClosed())
		fmt.Println("Before wait for connection")
		rabbitmq.cond.Wait()
		fmt.Println("After wait for connection")

	}
}
