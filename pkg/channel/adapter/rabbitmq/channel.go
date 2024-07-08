package rbbitmqchannel

import (
	"context"
	"fmt"
	"github.com/ormushq/ormus/adapter/otela"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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

func newChannelWithContext(ctx context.Context, done <-chan bool, wg *sync.WaitGroup, rabbitmqChannelParams rabbitmqChannelParams) (*rabbitmqChannel, error) {
	tracer := otela.NewTracer("rbbitmqchannel")
	_, span := tracer.Start(ctx, "rbbitmqchannel@newChannelWithContext")
	defer span.End()

	conn := rabbitmqChannelParams.rabbitmq.connection
	WaitForConnection(rabbitmqChannelParams.rabbitmq)
	ch, errChO := conn.Channel()
	if errChO != nil {
		return nil, errChO
	}
	span.AddEvent("connection-established")

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
	span.AddEvent("exchange-declared")

	errDQ := declareQueue(conn, rabbitmqChannelParams.queue)
	if errDQ != nil {
		return nil, errDQ
	}
	span.AddEvent("queue-declared")

	errBETQ := bindExchangeToQueue(conn, rabbitmqChannelParams.exchange, rabbitmqChannelParams.queue)
	if errBETQ != nil {
		return nil, errBETQ
	}
	span.AddEvent("bindings-established")

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
	rc.startWithContext(ctx)
	span.AddEvent("rabbitmq-channel-started")

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

func (rc *rabbitmqChannel) startWithContext(ctx context.Context) {
	tracer := otela.NewTracer("rbbitmqchannel")
	ctx, span := tracer.Start(ctx, "rbbitmqchannel@startWithContext")
	defer span.End()

	for i := 0; i < rc.numberInstants; i++ {
		if rc.mode.IsInputMode() {
			rc.wg.Add(1)
			go func() {
				defer rc.wg.Done()
				go rc.startInputWithContext(ctx)
			}()
		}
		if rc.mode.IsOutputMode() {
			rc.wg.Add(1)
			go func() {
				defer rc.wg.Done()
				rc.startOutputWitContext(ctx)
			}()
		}
	}
	span.AddEvent("channels-workers-started")
}

func (rc *rabbitmqChannel) startOutput() {
	rc.startOutputWitContext(context.Background())
}

func (rc *rabbitmqChannel) startOutputWitContext(ctx context.Context) {
	tracer := otela.NewTracer("rbbitmqchannel")
	_, span := tracer.Start(ctx, "rbbitmqchannel@startOutputWitContext")
	defer span.End()

	WaitForConnection(rc.rabbitmq)
	span.AddEvent("connection-established")

	rc.wg.Add(1)
	go func() {
		defer rc.wg.Done()

		ch, err := rc.rabbitmq.connection.Channel()
		if err != nil {
			logger.WithGroup(loggerGroupName).Error(errmsg.ErrFailedToOpenChannel,
				slog.String("error", err.Error()))
			rc.callMeNextTime(rc.startOutput)
			span.AddEvent("error-on-open-channel", trace.WithAttributes(
				attribute.String("error", err.Error()),
			))
			span.End()

			return
		}
		span.AddEvent("channel-opened")

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

			span.AddEvent("error-on-start-consume", trace.WithAttributes(
				attribute.String("error", err.Error()),
			))
			span.End()

			return
		}

		span.AddEvent("consume-started")

		span.End()

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
	rc.startInputWithContext(context.Background())
}

func (rc *rabbitmqChannel) startInputWithContext(ctx context.Context) {
	tracer := otela.NewTracer("rbbitmqchannel")
	_, span := tracer.Start(ctx, "rbbitmqchannel@startInputWithContext")
	defer span.End()

	WaitForConnection(rc.rabbitmq)
	span.AddEvent("connection-established")

	rc.wg.Add(1)
	go func() {
		defer rc.wg.Done()

		ch, err := rc.rabbitmq.connection.Channel()
		if err != nil {
			logger.WithGroup(loggerGroupName).Error(errmsg.ErrFailedToOpenChannel,
				slog.String("error", err.Error()))
			rc.callMeNextTime(rc.startInput)
			span.AddEvent("error-on-open-channel", trace.WithAttributes(
				attribute.String("error", err.Error()),
			))
			span.End()

			return
		}
		span.AddEvent("channel-opened")

		defer func() {
			err = ch.Close()
			if err != nil {
				logger.WithGroup(loggerGroupName).Error(errmsg.ErrFailedToCloseChannel,
					slog.String("error", err.Error()))
			}
		}()

		span.End()
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
