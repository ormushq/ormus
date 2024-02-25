package inmemorytaskmanager

import (
	"github.com/ormushq/ormus/destination/entity"
	"log"
)

// Queue represents an in-memory job queue.
type Queue struct {
	Tasks chan *entity.Task
}

func NewQueue() *Queue {
	tasks := make(chan *entity.Task, 1)
	return &Queue{
		Tasks: tasks,
	}
}

// Enqueue adds a new job to the in-memory queue.
func (q *Queue) Enqueue(task *entity.Task) error {
	log.Printf("Task [%s] is published to In-Memory queue.", task.ID)
	q.Tasks <- task

	println(q.Tasks)
	return nil
}
