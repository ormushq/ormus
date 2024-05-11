package worker

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/ormushq/ormus/destination/taskmanager"
	"github.com/ormushq/ormus/event"
)

const timeoutInSeconds = 5

type Worker struct {
	TaskHandler   TaskHandler
	EventsChannel <-chan event.ProcessedEvent
}

type TaskHandler interface {
	HandleTask(ctx context.Context, newEvent event.ProcessedEvent) error
}

func (w *Worker) Run(done <-chan bool, wg *sync.WaitGroup) error {
	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Info("Start rabbitmq worker for handling tasks.")
		for {
			select {
			case newEvent := <-w.EventsChannel:
				go func() {
					ctx, cancel := context.WithTimeout(context.Background(), timeoutInSeconds*time.Second)
					defer cancel()
					err := w.TaskHandler.HandleTask(ctx, newEvent)
					if err != nil {
						slog.Error(fmt.Sprintf("Error on handling event using integration handler.Error : %v", err))
					}
				}()
			case <-done:

				return
			}
		}
	}()

	return nil
}

func NewWorker(events <-chan event.ProcessedEvent, th taskmanager.TaskHandler) *Worker {
	return &Worker{
		EventsChannel: events,
		TaskHandler:   th,
	}
}
