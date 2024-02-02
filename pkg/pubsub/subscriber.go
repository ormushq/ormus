package pubsub

import (
	"sync"

	"github.com/google/uuid"
	"github.com/ormushq/ormus/logger"
)

type subscriber struct {
	id uuid.UUID

	sending        sync.Mutex
	messageChannel chan *Message
}

func (s *subscriber) sendMessageToSubscriber(publisherName string, msg *Message) {
	s.sending.Lock()
	defer s.sending.Unlock()

	select {

	case s.messageChannel <- msg:
		logger.L().Info("sending message to subscriber successfully", "publisher name", publisherName)

	default:
		logger.L().Info("there is no message to send")
	}
}
