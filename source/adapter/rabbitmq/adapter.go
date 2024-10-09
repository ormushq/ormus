package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQAdapter struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	messages   chan *Message
}

type Message struct {
	Body        []byte
	DeliveryTag uint64 // Delivery tag for acknowledgment
}

func NewRabbitMQAdapter(amqpURI string) (*RabbitMQAdapter, error) {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQAdapter{
		connection: conn,
		channel:    ch,
		messages:   make(chan *Message), // Initialize channel
	}, nil
}

func (r *RabbitMQAdapter) QueueDeclare(queueName string) error {
	_, err := r.channel.QueueDeclare(
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

func (r *RabbitMQAdapter) Publish(queueName string, message []byte) error {
	err := r.QueueDeclare(queueName)
	if err != nil {
		return err
	}

	err = r.channel.Publish(
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

func (r *RabbitMQAdapter) Subscribe(queueName string) (chan *Message, error) {
	err := r.QueueDeclare(queueName)
	if err != nil {
		return nil, err
	}
	err = r.channel.Qos(7, 0, false)
	if err != nil {
		return nil, err
	}
	msgs, err := r.channel.Consume(
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
			r.messages <- &Message{Body: msg.Body, DeliveryTag: msg.DeliveryTag} // Send received message with delivery tag
		}
	}()
	return r.messages, nil
}

func (r *RabbitMQAdapter) Ack(msg *Message) error { // Acknowledge specific message
	return r.channel.Ack(msg.DeliveryTag, false)
}

func (r *RabbitMQAdapter) Close() {
	if err := r.channel.Close(); err != nil {
		log.Println("Error closing channel:", err)
	}
	if err := r.connection.Close(); err != nil {
		log.Println("Error closing connection:", err)
	}
}
