package workers

import (
	"sync"

	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/service/projectservice"
	"github.com/ormushq/ormus/pkg/channel/adapter/simple"
)

type Worker struct {
	prSvc          *projectservice.Service
	internalBroker *simple.ChannelAdapter
}

func New(prSvc *projectservice.Service, internalBroker *simple.ChannelAdapter) *Worker {
	return &Worker{
		prSvc:          prSvc,
		internalBroker: internalBroker,
	}
}

func (w *Worker) Run(done <-chan bool, wg *sync.WaitGroup) {
	internalBroker, err := w.internalBroker.GetOutputChannel("CreateDefaultProject")
	if err != nil {
		logger.L().Debug("error on getting internal broker channel")

		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-internalBroker:
				err := w.prSvc.Create()
				if err != nil {
					logger.L().Error("creating project")

					break
				}
				err = msg.Ack()
				if err != nil {
					logger.L().Error("message ack")

					break
				}
			case <-done:
				return

			}
		}
	}()
}
