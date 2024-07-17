package fakedeliveryhandler

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskdelivery/param"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type FakeHandler struct{}

func New() *FakeHandler {
	return &FakeHandler{}
}

const fakeProcessingTimeSecond = 2

func (h FakeHandler) Handle(t taskentity.Task) (param.DeliveryTaskResponse, error) {
	tracer := otela.NewTracer("fakedeliveryhandler")
	_, span := tracer.Start(otela.GetContextFromCarrier(t.ProcessedEvent.TracerCarrier), "fakedeliveryhandler@Handle")
	defer span.End()

	time.Sleep(fakeProcessingTimeSecond * time.Second)

	result := fmt.Sprintf("Task [%s] handled successfully! ✅ ", t.ID)

	slog.Info(result)
	span.AddEvent("result", trace.WithAttributes(
		attribute.String("result", result),
	))

	res := param.DeliveryTaskResponse{
		Attempts:       1,
		FailedReason:   nil,
		DeliveryStatus: taskentity.SuccessTaskStatus,
	}

	return res, nil
}
