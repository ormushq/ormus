package rbbitmqchannel

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitmqChannel struct {
	wg             *sync.WaitGroup
	done           <-chan bool
	mode           channel.Mode
	rabbitmq       *Rabbitmq
	inputChannel   chan []byte
	outputChannel  chan channel.Message
	exchange       string
	queue          string
	numberInstants int
	maxRetryPolicy int
}
type rabbitmqChannelParams struct {
	mode           channel.Mode
	rabbitmq       *Rabbitmq
	exchange       string
	queue          string
	bufferSize     int
	numberInstants int
	maxRetryPolicy int
}

const (
	timeForCallAgainDuration = 10
	retriesToTimeRatio       = 2
)

func newChannel(done <-chan bool, wg *sync.WaitGroup, rabbitmqChannelParams rabbitmqChannelParams) (*rabbitmqChannel, error) {
	conn := rabbitmqChannelParams.rabbitmq.connection
	WaitForConnection(rabbitmqChannelParams.rabbitmq)
	ch, errChO := conn.Channel()
	if errChO != nil {
		return nil, errChO
	}

	defer func() {
		err := ch.Close()
		if err != nil {
			logger.L().Error("failed to close rabbitmq channel", err)
		}
	}()

	errDE := ch.ExchangeDeclare(
		rabbitmqChannelParams.exchange, // name
		"topic",                        // type
		true,                           // durable
		false,                          // auto-deleted
		false,                          // internal
		false,                          // no-wait
		nil,                            // arguments
	)
	if errDE != nil {
		ch, errChO = conn.Channel()
		if errChO != nil {
			return nil, errChO
		}
		errDE = ch.ExchangeDeclarePassive(
			rabbitmqChannelParams.exchange, // name
			"topic",                        // type
			true,                           // durable
			false,                          // auto-deleted
			false,                          // internal
			false,                          // no-wait
			nil,                            // arguments
		)
		if errDE != nil {
			return nil, errDE
		}
	}
	_, errQD := ch.QueueDeclare(
		rabbitmqChannelParams.queue, // name
		true,                        // durable
		false,                       // delete when unused
		false,                       // exclusive
		false,                       // no-wait
		nil,                         // arguments
	)

	if errQD != nil {
		ch, errChO = conn.Channel()
		if errChO != nil {
			return nil, errChO
		}
		_, errQD = ch.QueueDeclarePassive(
			rabbitmqChannelParams.queue, // name
			true,                        // durable
			false,                       // delete when unused
			false,                       // exclusive
			false,                       // no-wait
			nil,                         // arguments
		)
		if errQD != nil {
			return nil, errQD
		}
	}

	errQB := ch.QueueBind(
		rabbitmqChannelParams.queue,    // queue name
		"",                             // routing key
		rabbitmqChannelParams.exchange, // exchange
		false,
		nil)
	if errQB != nil {
		return nil, errQB
	}

	rc := &rabbitmqChannel{
		done:           done,
		wg:             wg,
		mode:           rabbitmqChannelParams.mode,
		exchange:       rabbitmqChannelParams.exchange,
		queue:          rabbitmqChannelParams.queue,
		rabbitmq:       rabbitmqChannelParams.rabbitmq,
		numberInstants: rabbitmqChannelParams.numberInstants,
		maxRetryPolicy: rabbitmqChannelParams.maxRetryPolicy,
		inputChannel:   make(chan []byte, rabbitmqChannelParams.bufferSize),
		outputChannel:  make(chan channel.Message, rabbitmqChannelParams.bufferSize),
	}
	rc.start()

	return rc, nil
}

func (rc *rabbitmqChannel) GetMode() channel.Mode {
	return rc.mode
}

func (rc *rabbitmqChannel) GetInputChannel() chan<- []byte {
	return rc.inputChannel
}

func (rc *rabbitmqChannel) GetOutputChannel() <-chan channel.Message {
	return rc.outputChannel
}

func (rc *rabbitmqChannel) start() {
	for i := 0; i < rc.numberInstants; i++ {
		if rc.mode.IsInputMode() {
			rc.wg.Add(1)
			go func() {
				defer rc.wg.Done()
				go rc.startInput()
			}()
		}
		if rc.mode.IsOutputMode() {
			rc.wg.Add(1)
			go func() {
				defer rc.wg.Done()
				rc.startOutput()
			}()
		}
	}
}

func (rc *rabbitmqChannel) startOutput() {
	rc.wg.Add(1)
	WaitForConnection(rc.rabbitmq)
	go func() {
		defer rc.wg.Done()
		ch, err := rc.rabbitmq.connection.Channel()

		if err != nil {
			logger.L().Error(err.Error(), err)
			rc.callMeNextTime(rc.startOutput)
			return
		}

		defer func(ch *amqp.Channel) {
			err = ch.Close()
			if err != nil {
				logger.L().Error(err.Error(), err)
			}
		}(ch)

		msgs, errConsume := ch.Consume(
			rc.queue, // queue
			"",       // consumer
			false,    // auto-ack
			false,    // exclusive
			false,    // no-local
			false,    // no-wait
			nil,      // arguments
		)
		if errConsume != nil {
			logger.L().Error(errConsume.Error(), err)
			rc.callMeNextTime(rc.startOutput)
			return
		}

		for {
			if ch.IsClosed() {
				logger.L().Debug("channel is closed rerun startOutput function")
				rc.callMeNextTime(rc.startOutput)
				return
			}
			select {
			case <-rc.done:

				return
			case msg := <-msgs:
				rc.wg.Add(1)
				go func() {
					defer rc.wg.Done()

					rc.outputChannel <- channel.Message{
						Body: msg.Body,
						Ack: func() error {
							return msg.Ack(false)
						},
					}
				}()
			}
		}
	}()
}

func (rc *rabbitmqChannel) startInput() {
	rc.wg.Add(1)
	WaitForConnection(rc.rabbitmq)

	go func() {
		defer rc.wg.Done()
		ch, err := rc.rabbitmq.connection.Channel()
		if err != nil {
			logger.L().Error(err.Error(), err)
			rc.callMeNextTime(rc.startInput)
			return
		}

		defer func() {
			err = ch.Close()
			if err != nil {
				logger.L().Error(err.Error(), err)
			}
		}()

		for {
			if ch.IsClosed() {
				logger.L().Debug("channel is closed rerun startInput function")
				rc.callMeNextTime(rc.startInput)
				return
			}
			select {
			case <-rc.done:

				return
			case msg := <-rc.inputChannel:
				rc.publishToRabbitmq(ch, msg, 0)
			}
		}
	}()
}

func (rc *rabbitmqChannel) publishToRabbitmq(ch *amqp.Channel, msg []byte, tries int) {
	if tries > rc.maxRetryPolicy {
		logger.L().Error(fmt.Sprintf("job failed after %d tries", tries))
		return
	}
	rc.wg.Add(1)
	go func() {
		defer rc.wg.Done()
		time.Sleep(time.Second * time.Duration(tries*retriesToTimeRatio))
		errPWC := ch.PublishWithContext(context.Background(),
			rc.exchange, // exchange
			"",          // routing key
			false,       // mandatory
			false,       // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        msg,
			})
		if errPWC != nil {
			logger.L().Error("error on publish to rabbitmq", errPWC)

			rc.publishToRabbitmq(ch, msg, tries+1)
		}
	}()
}

func (rc *rabbitmqChannel) callMeNextTime(f func()) {
	rc.wg.Add(1)
	go func() {
		time.Sleep(time.Second * timeForCallAgainDuration)
		defer rc.wg.Done()
		f()
	}()
}
