package workers

import (
	"encoding/json"
	"sync"

	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/managerparam"
	"github.com/ormushq/ormus/manager/managerparam/projectparam"
	"github.com/ormushq/ormus/manager/service/projectservice"
	"github.com/ormushq/ormus/pkg/channel/adapter/simplechannel"
)

type Worker struct {
	prSvc          projectservice.Service
	internalBroker *simplechannel.ChannelAdapter
}

func New(prSvc projectservice.Service, internalBroker *simplechannel.ChannelAdapter) *Worker {
	return &Worker{
		prSvc:          prSvc,
		internalBroker: internalBroker,
	}
}

func (w *Worker) Run(done <-chan bool, wg *sync.WaitGroup) {
	logger.L().Debug("workers.Run")
	internalBroker, err := w.internalBroker.GetOutputChannel(managerparam.CreateDefaultProject)
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
				var req projectparam.CreateThoughChannel
				err = json.Unmarshal(msg.Body, &req)
				if err != nil {
					logger.L().Error(err.Error())
				}

				_, createErr := w.prSvc.Create(projectparam.CreateRequest(req))
				if createErr != nil {
					logger.L().Error("err on creating project", "err msg:", createErr)

					break
				}
				ackErr := msg.Ack()
				if ackErr != nil {
					logger.L().Error("err on acking message", "err msg:", ackErr)

					break
				}
			case <-done:
				return

			}
		}
	}()
}
