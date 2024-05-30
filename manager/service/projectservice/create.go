package projectservice

import (
	"fmt"
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
	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random integer between 0 and 99
	for {
		select {
		case msg := <-inOutChan:
			_, err := s.repo.Create("Default"+strconv.Itoa(rand.Intn(100)), string(msg.Body))
			if err != nil {
				return err
			}
			fmt.Println(msg.Body)
			err = msg.Ack()
			if err != nil {
				return err
			}
			fmt.Println("created")
		}

	}
}
