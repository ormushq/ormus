# Event Manager

The Event Manager simplifies handling events across your project, offering an intuitive interface to publish or consume events seamlessly.

## Usage

To begin using the Event Manager, create a new `adapter.Channel`.

Example code can be found in `examples/eventmanager/main.go`.


```go
package main

import (
	"context"
	"fmt"
	"github.com/ormushq/ormus/contract/go/internalevent"
	"github.com/ormushq/ormus/contract/go/project"
	"github.com/ormushq/ormus/eventmanager"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/channel/adapter/rabbitmqchannel"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"sync"
	"time"
)

func main() {
	done := make(chan bool)
	wg := &sync.WaitGroup{}

	// Init channel package you can use rabbitmqchannel or simple
	rc := rabbitmqchannel.New(done, wg, rabbitmqchannel.Config{
		User:            "guest",
		Password:        "guest",
		Host:            "127.0.0.1",
		Port:            5672,
		Vhost:           "/",
		ReconnectSecond: 2,
	})

	// Init event manager with list of internal event you want to use with this manager
	eveMnger, err := eventmanager.New(wg, done, rc, map[internalevent.EventName]eventmanager.CreateChannelFunc{
		internalevent.EventName_EVENT_NAME_PROJECT_CREATED: eventmanager.NewCreateChannelFunc(rc, channel.BothMode, 1, 10),
	})
	if err != nil {
		log.Fatal(err)
	}
	
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
		fmt.Println(msg)
	}()

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
	
	wg.Wait()

}

```


### Steps:
1. Initialize a new channel (either RabbitMQ or Simple).
2. Set up the Event Manager with a list of events to use in this instance.
3. Start consuming or publishing events as needed.

### Consuming Events

To consume events, call the `Consume` method on the Event Manager, passing in one or more event names. This method returns a Go channel of type `eventmanager.EventMessage`.

### Publishing Events

To publish events, use the `Publish` method, which accepts an `*internalevent.Event`. You can create an `*internalevent.Event` using the following helper functions:

- `eventmanager.NewWriteKeyGeneratedEvent`
- `eventmanager.NewUserCreatedEvent`
- `eventmanager.NewProjectCreatedEvent`
- `eventmanager.NewTaskCreatedEvent`

**Note:** Events can only be published or consumed if they were specified when initializing the Event Manager.