package taskmanager

// Worker represents a worker that processes jobs of queue.
type Worker interface {
	ProcessJobs()
}
