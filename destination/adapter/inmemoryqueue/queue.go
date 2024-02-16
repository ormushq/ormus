package inmemoryqueue

import (
	"github.com/ormushq/ormus/destination/queue"
	"sync"
)

// Queue represents an in-memory job queue.
type Queue struct {
	jobs  []queue.Job
	mutex sync.Mutex
}

// Enqueue adds a new job to the in-memory queue.
func (q *Queue) Enqueue(job queue.Job) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.jobs = append(q.jobs, job)
}

// Dequeue retrieves and removes the oldest job from the in-memory queue.
func (q *Queue) Dequeue(topic string) queue.Job {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	for i, job := range q.jobs {
		if job.Topic == topic {
			// Found a job with a matching topic, remove it from the queue and return it
			q.jobs = append(q.jobs[:i], q.jobs[i+1:]...)
			return job
		}
	}

	return queue.Job{} // Return an empty job if no matching job is found
}
