package rbbitmqchannel

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/errmsg"
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
	maxRetryPolicy int
	bufferSize     int
}
type rabbitmqChannelParams struct {
	mode           channel.Mode
	rabbitmq       *Rabbitmq
	exchange       string
	queue          string
	bufferSize     int
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

	defer func(c *amqp.Channel) {
		err := c.Close()
		if err != nil {
			logger.WithGroup(loggerGroupName).Error(errmsg.ErrFailedToCloseChannel,
				slog.String("error", err.Error()))
		}
	}(ch)

	errDE := declareExchange(conn, rabbitmqChannelParams.exchange)
	if errDE != nil {
		return nil, errDE
	}

	errDQ := declareQueue(conn, rabbitmqChannelParams.queue)
	if errDQ != nil {
		return nil, errDQ
	}

	errBETQ := bindExchangeToQueue(conn, rabbitmqChannelParams.exchange, rabbitmqChannelParams.queue)
	if errBETQ != nil {
		return nil, errBETQ
	}

	rc := &rabbitmqChannel{
		done:           done,
		wg:             wg,
		mode:           rabbitmqChannelParams.mode,
		exchange:       rabbitmqChannelParams.exchange,
		queue:          rabbitmqChannelParams.queue,
		rabbitmq:       rabbitmqChannelParams.rabbitmq,
		maxRetryPolicy: rabbitmqChannelParams.maxRetryPolicy,
		inputChannel:   make(chan []byte, rabbitmqChannelParams.bufferSize),
		outputChannel:  make(chan channel.Message, rabbitmqChannelParams.bufferSize),
		bufferSize:     rabbitmqChannelParams.bufferSize,
	}
	rc.start()

	return rc, nil
}

func declareExchange(conn *amqp.Connection, exchange string) error {
	ch, errChO := conn.Channel()
	if errChO != nil {
		return errChO
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			logger.WithGroup(loggerGroupName).Error(errmsg.ErrFailedToCloseChannel,
				slog.String("error", err.Error()))
		}
	}(ch)

	errDE := ch.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if errDE != nil {
		ch, errChO = conn.Channel()
		if errChO != nil {
			return errChO
		}
		errDE = ch.ExchangeDeclarePassive(
			exchange, // name
			"topic",  // type
			true,     // durable
			false,    // auto-deleted
			false,    // internal
			false,    // no-wait
			nil,      // arguments
		)
		if errDE != nil {
			return errDE
		}
	}

	return nil
}

func declareQueue(conn *amqp.Connection, queue string) error {
	ch, errChO := conn.Channel()
	if errChO != nil {
		return errChO
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			logger.WithGroup(loggerGroupName).Error(errmsg.ErrFailedToCloseChannel,
				slog.String("error", err.Error()))
		}
	}(ch)

	_, errQD := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if errQD != nil {
		ch, errChO = conn.Channel()
		if errChO != nil {
			return errChO
		}
		_, errQD = ch.QueueDeclarePassive(
			queue, // name
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if errQD != nil {
			return errQD
		}
	}

	return nil
}

func bindExchangeToQueue(conn *amqp.Connection, exchange, queue string) error {
	ch, errChO := conn.Channel()
	if errChO != nil {
		return errChO
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			logger.WithGroup(loggerGroupName).Error(errmsg.ErrFailedToCloseChannel,
				slog.String("error", err.Error()))
		}
	}(ch)

	return ch.QueueBind(
		queue,    // queue name
		"",       // routing key
		exchange, // exchange
		false,
		nil)
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

func (rc *rabbitmqChannel) startOutput() {
	WaitForConnection(rc.rabbitmq)

	rc.wg.Add(1)
	go func() {
		defer rc.wg.Done()

		ch, err := rc.rabbitmq.connection.Channel()
		if err != nil {
			logger.WithGroup(loggerGroupName).Error(errmsg.ErrFailedToOpenChannel,
				slog.String("error", err.Error()))
			rc.callMeNextTime(rc.startOutput)

			return
		}

		err = ch.Qos(rc.bufferSize, 0, false)
		if err != nil {
			logger.L().Error(errmsg.ErrFailedToSetQosOnChannel)
			rc.callMeNextTime(rc.startOutput)

			return
		}

		defer func(ch *amqp.Channel) {
			err = ch.Close()
			if err != nil {
				logger.WithGroup(loggerGroupName).Error(errmsg.ErrFailedToCloseChannel,
					slog.String("error", err.Error()))
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
			logger.WithGroup(loggerGroupName).Error("failed to start consume",
				slog.String("error", err.Error()))
			rc.callMeNextTime(rc.startOutput)

			return
		}

		for {
			if ch.IsClosed() {
				logger.WithGroup(loggerGroupName).
					Debug("channel is closed rerun startOutput function")
				rc.callMeNextTime(rc.startOutput)

				return
			}
			select {
			case <-rc.done:

				return
			case msg := <-msgs:
				rc.outputChannel <- channel.Message{
					Body: msg.Body,
					Ack: func() error {
						return msg.Ack(false)
					},
				}
			}
		}
	}()
}

func (rc *rabbitmqChannel) startInput() {
	WaitForConnection(rc.rabbitmq)

	rc.wg.Add(1)
	go func() {
		defer rc.wg.Done()

		ch, err := rc.rabbitmq.connection.Channel()
		if err != nil {
			logger.WithGroup(loggerGroupName).Error(errmsg.ErrFailedToOpenChannel,
				slog.String("error", err.Error()))
			rc.callMeNextTime(rc.startInput)

			return
		}

		defer func() {
			err = ch.Close()
			if err != nil {
				logger.WithGroup(loggerGroupName).Error(errmsg.ErrFailedToCloseChannel,
					slog.String("error", err.Error()))
			}
		}()

		for {
			if ch.IsClosed() {
				logger.WithGroup(loggerGroupName).Debug("channel is closed rerun startInput function")
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
		logger.WithGroup(loggerGroupName).Error(fmt.Sprintf("job failed after %d tries", tries))

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
			logger.WithGroup(loggerGroupName).Error("error on publish to rabbitmq",
				slog.String("error", errPWC.Error()))

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
