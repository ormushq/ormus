package dtcoordinator

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/ormushq/ormus/destination/taskmanager"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/manager/entity"
)

type TaskPublisherMap map[entity.DestinationType]taskmanager.Publisher

type DestinationTypeCoordinator struct {
	TaskPublishers TaskPublisherMap
}

func New(taskPublishers TaskPublisherMap) DestinationTypeCoordinator {
	return DestinationTypeCoordinator{
		TaskPublishers: taskPublishers,
	}
}

func (d DestinationTypeCoordinator) Start(processedEvents <-chan event.ProcessedEvent, done <-chan bool, wg *sync.WaitGroup) error {
	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Info("Starting destination type task coordinator.")

		for {
			select {
			case pe := <-processedEvents:

				taskPublisher, ok := d.TaskPublishers[pe.DestinationType()]
				if !ok {
					slog.Error(fmt.Sprintf("Error on finding task manager for %s", pe.DestinationType()))

					break
				}

				pErr := taskPublisher.Publish(pe)
				if pErr != nil {
					slog.Error(fmt.Sprintf("Error on publishing event : %s", pErr))

					break
				}

			case <-done:

				return
			}
		}
	}()

	return nil
}
