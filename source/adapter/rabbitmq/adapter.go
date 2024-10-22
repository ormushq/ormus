package rabbitmq

import (
	rabbitmq "github.com/ormushq/ormus/adapter/rabbitmq"
	"github.com/streadway/amqp"
	"log"
)

type rabbitMQAdapter struct {
	adapter rabbitmq.Adapter
}

func NewRabbitMQAdapter(adapter rabbitmq.Adapter) rabbitMQAdapter {
	return rabbitMQAdapter{
		adapter: adapter,
	}
}

func (r rabbitMQAdapter) QueueDeclare(queueName string) error {
	_, err := r.adapter.GetChannel().QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	return nil

}

func (r rabbitMQAdapter) Publish(queueName string, message []byte) error {
	err := r.QueueDeclare(queueName)
	if err != nil {
		return err
	}

	err = r.adapter.GetChannel().Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)

	return err

}

func (r rabbitMQAdapter) Subscribe(queueName string) (chan *rabbitmq.Message, error) {
	err := r.QueueDeclare(queueName)
	if err != nil {
		return nil, err
	}
	prefetchCount := 7
	prefetchSize := 0
	err = r.adapter.GetChannel().Qos(prefetchCount, prefetchSize, false)
	if err != nil {
		return nil, err
	}
	msgs, err := r.adapter.GetChannel().Consume(
		queueName,
		"",    // consumer
		false, // auto-ack (set to false for manual ack)
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
		return nil, err
	}

	go func() {
		for msg := range msgs {
			r.adapter.GetMessage() <- &rabbitmq.Message{Body: msg.Body, DeliveryTag: msg.DeliveryTag} // Send received message with delivery tag
		}
	}()
	return r.adapter.GetMessage(), nil
}

func (r rabbitMQAdapter) Ack(msg *rabbitmq.Message) error { // Acknowledge specific message
	return r.adapter.GetChannel().Ack(msg.DeliveryTag, false)
}

func (r rabbitMQAdapter) Close() {
	if err := r.adapter.GetChannel().Close(); err != nil {
		log.Println("Error closing channel:", err)
	}
	if err := r.adapter.GetConnection().Close(); err != nil {
		log.Println("Error closing connection:", err)
	}
}
