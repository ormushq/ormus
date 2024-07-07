package dtcoordinator

import (
	"fmt"
	"github.com/ormushq/ormus/adapter/otela"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
		tracer := otela.NewTracer("dtcoordinator")

		slog.Info("Starting destination type task coordinator.")

		for {
			select {
			case pe := <-processedEvents:
				wg.Add(1)
				go func() {
					defer wg.Done()
					ctx := otela.GetContextFromCarrier(pe.TracerCarrier)
					ctx, span := tracer.Start(ctx, "dtcoordinator@ProcessEvent")
					defer span.End()

					taskPublisher, ok := d.TaskPublishers[pe.DestinationType()]
					if !ok {
						span.AddEvent("error-on-get-task-publisher", trace.WithAttributes(
							attribute.String("destination-type", string(pe.DestinationType())),
							attribute.String("error-message", "Task manager not found"),
						))
						slog.Error(fmt.Sprintf("Error on finding task manager for %s", pe.DestinationType()))

						return
					}
					span.AddEvent("task-publisher-founded")

					pe.TracerCarrier = otela.GetCarrierFromContext(ctx)

					pErr := taskPublisher.Publish(pe)
					if pErr != nil {
						span.AddEvent("error-on-publish-event", trace.WithAttributes(
							attribute.String("error-message", pErr.Error()),
						))
						slog.Error(fmt.Sprintf("Error on publishing event : %s", pErr))

						return
					}
					span.AddEvent("task-published")
				}()

			case <-done:

				return
			}
		}
	}()

	return nil
}
