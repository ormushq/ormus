package projectservice

import (
	"github.com/ormushq/ormus/logger"
	"math/rand"
	"strconv"
	"time"
)

func (s Service) Create() error {
	//const op = "projectservice.Create"
	inOutChan, err := s.internalBroker.GetOutputChannel("CreateDefaultProject")
	if err != nil {
		return err
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		select {
		case msg := <-inOutChan:
			_, err := s.repo.Create("Default"+strconv.Itoa(rand.Intn(100)), string(msg.Body))
			if err != nil {
				return err
			}
			logger.L().Debug(string(msg.Body))
			err = msg.Ack()
			if err != nil {
				return err
			}
			logger.L().Debug("Default project created")
		}

	}
}
