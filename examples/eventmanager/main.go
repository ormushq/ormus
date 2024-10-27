package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ormushq/ormus/contract/go/internalevent"
	"github.com/ormushq/ormus/contract/go/project"
	"github.com/ormushq/ormus/eventmanager"
	"github.com/ormushq/ormus/pkg/channel"
	rbbitmqchannel "github.com/ormushq/ormus/pkg/channel/adapter/rabbitmq"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	done := make(chan bool)
	wg := &sync.WaitGroup{}

	port := 5672
	reconnectSecond := 2
	// Init channel package you can use rabbitmqchannel or simple
	rc := rbbitmqchannel.New(done, wg, rbbitmqchannel.Config{
		User:            "guest",
		Password:        "guest",
		Host:            "127.0.0.1",
		Port:            port,
		Vhost:           "/",
		ReconnectSecond: reconnectSecond,
	})
	fmt.Println("RabbitMQ channel initialized")

	bufferSize := 1
	maxRetryPolicy := 10

	// Init event manager with list of internal event you want to use with this manager
	eveMnger, err := eventmanager.New(wg, done, rc, map[internalevent.EventName]eventmanager.CreateChannelFunc{
		internalevent.EventName_EVENT_NAME_PROJECT_CREATED: eventmanager.NewCreateChannelFunc(rc, channel.BothMode, bufferSize, maxRetryPolicy),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Event manager initialized")

	// For consume you can call Consume method. It returns a channel with type of eventmanager.EventMessage
	// The consume method accept multiple event name
	wg.Add(1)
	go func() {
		defer wg.Done()

		outputChannel, err := eveMnger.Consume(internalevent.EventName_EVENT_NAME_PROJECT_CREATED)
		if err != nil {
			log.Fatal(err)
		}
		msg := <-outputChannel
		fmt.Println("Message received")
		fmt.Println(msg)
	}()
	fmt.Println("Consumer initialized")

	// You can publish event with call publish on event manager. It accepts `*internalevent.Event.`
	// For create `*internalevent.Event` you can use helpers.
	err = eveMnger.Publish(eventmanager.NewProjectCreatedEvent(context.Background(), &internalevent.ProjectCreatedEvent{
		Project: &project.Project{
			Id:          "TestID",
			CreatedAt:   timestamppb.New(time.Now()),
			UpdatedAt:   timestamppb.New(time.Now()),
			DeletedAt:   timestamppb.New(time.Now()),
			Name:        "Test",
			Description: "Test",
			UserId:      "TestUserID",
		},
	}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Message published")

	time.Sleep(1 * time.Second)
	close(done)

	wg.Wait()
}
