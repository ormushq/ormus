package worker

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/metricname"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Worker struct {
	TaskHandler   TaskHandler
	EventsChannel <-chan event.ProcessedEvent
}

type TaskHandler interface {
	HandleTask(ctx context.Context, newEvent event.ProcessedEvent) error
}

func (w *Worker) Run(done <-chan bool, wg *sync.WaitGroup) error {
	tracer := otela.NewTracer("worker")
	meter := otel.Meter("worker@Run")

	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Info("Start rabbitmq worker for handling tasks.")
		for {
			select {
			case newEvent := <-w.EventsChannel:
				wg.Add(1)
				go func() {
					defer wg.Done()
					defer func() {
						if r := recover(); r != nil {
							logger.L().Error(fmt.Sprintf("worker recovered panic: %v", r))
						}
					}()

					ctx := otela.GetContextFromCarrier(newEvent.TracerCarrier)
					ctx, span := tracer.Start(ctx, "worker@Run")

					defer span.End()
					span.AddEvent("event-received-in-worker")
					err := w.TaskHandler.HandleTask(ctx, newEvent)
					if err != nil {
						otela.IncrementFloat64Counter(ctx, meter, metricname.DestinationWorkerHandleEventError, "process_event_handle_error")

						span.AddEvent("error-on-handle-task", trace.WithAttributes(
							attribute.String("error", err.Error())))

						slog.Error(fmt.Sprintf("Error on handling event using integration handler.Error : %v", err))

						return
					}

					otela.IncrementFloat64Counter(ctx, meter, metricname.ProcessFlowOutputDestinationWorkerDoneJob, "event_handled_publish_done_job")
				}()
			case <-done:

				return
			}
		}
	}()

	return nil
}

func NewWorker(events <-chan event.ProcessedEvent, th TaskHandler) *Worker {
	return &Worker{
		EventsChannel: events,
		TaskHandler:   th,
	}
}
