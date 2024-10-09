package rabbitmq

import (
	rabbitmq "github.com/ormushq/ormus/adapter/rabbitmq"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQAdapter struct {
	adapter rabbitmq.Adapter
}

func NewRabbitMQAdapter(adapter rabbitmq.Adapter) RabbitMQAdapter {
	return RabbitMQAdapter{
		adapter: adapter,
	}
}

func (r RabbitMQAdapter) QueueDeclare(queueName string) error {
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

func (r RabbitMQAdapter) Publish(queueName string, message []byte) error {
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

func (r RabbitMQAdapter) Subscribe(queueName string) (chan *rabbitmq.Message, error) {
	err := r.QueueDeclare(queueName)
	if err != nil {
		return nil, err
	}
	err = r.adapter.GetChannel().Qos(7, 0, false)
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

func (r RabbitMQAdapter) Ack(msg *rabbitmq.Message) error { // Acknowledge specific message
	return r.adapter.GetChannel().Ack(msg.DeliveryTag, false)
}

func (r RabbitMQAdapter) Close() {
	if err := r.adapter.GetChannel().Close(); err != nil {
		log.Println("Error closing channel:", err)
	}
	if err := r.adapter.GetConnection().Close(); err != nil {
		log.Println("Error closing connection:", err)
	}
}
