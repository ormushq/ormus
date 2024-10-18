package eventmanager

// event name pattern : PublisherServiceName.EntityName.Trigger.
const (
	EventNameWriteKeyGenerated = "manager.write_key.generated"
	EventNameUserCreated       = "manager.user.created"
	EventNameProjectCreated    = "manager.project.created"
	EventNameTaskCreated       = "destination.task.created"
)
