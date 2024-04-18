package rabbitmqtaskmanager

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"sync"
	"time"

	"github.com/ormushq/ormus/destination/taskmanager"
)

const timeoutInSeconds = 5

type Worker struct {
	TaskConsumer taskmanager.Consumer
	TaskHandler  taskmanager.TaskHandler
}

func (w *Worker) Run(done <-chan bool, wg *sync.WaitGroup) error {
	processedEventsChannel, err := w.TaskConsumer.Consume(done, wg)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Info("Start rabbitmq worker for handling tasks.")

		for {
			select {
			case newEvent := <-processedEventsChannel:
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

func NewWorker(c taskmanager.Consumer, th taskmanager.TaskHandler) *Worker {
	return &Worker{
		TaskConsumer: c,
		TaskHandler:  th,
	}
}

func panicOnWorkersError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func printWorkersError(err error, msg string) {
	log.Printf("%s: %s", msg, err)
}
