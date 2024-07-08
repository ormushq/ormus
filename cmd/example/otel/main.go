package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/metric"
	"sync"
	"time"

	"github.com/ormushq/ormus/adapter/otela"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	port = 8081
	cfg  = otela.Config{
		Endpoint:           "otel_collector:4317",
		EnableMetricExpose: false,
		MetricExposePath:   "metrics",
		MetricExposePort:   port,
	}
)

func main() {
	wg := &sync.WaitGroup{}
	done := make(chan bool)
	name := "test-service"
	cfg.ServiceName = name

	wg.Add(1)
	chanSize := 10
	c := make(chan context.Context, chanSize)
	go startService2(wg, done, c)
	err := otela.Configure(wg, done, cfg)
	if err != nil {
		panic(err)
	}
	startService1(c)

	sleepTime := 5
	time.Sleep(time.Duration(sleepTime) * time.Second)

	close(done)
	wg.Wait()
}
func startService1(c chan<- context.Context) {
	tracer := otela.NewTracer("test-tracer")

	ctx, span := tracer.Start(context.Background(), "test-span")

	defer span.End()

	span.AddEvent("test-event")

	subService1(ctx, tracer)

	sendMetric()

	c <- ctx
}

func subService1(ctx context.Context, tracer trace.Tracer) {
	_, span := tracer.Start(ctx, "subService2")
	defer span.End()

	span.AddEvent("Starting work")
	sleepTime := 2
	time.Sleep(time.Duration(sleepTime) * time.Second)

	sendMetric()

	span.AddEvent("Work completed")
}

func startService2(wg *sync.WaitGroup, done <-chan bool, c <-chan context.Context) {
	defer wg.Done()
	ctx := <-c
	name := "test-service2"
	cfg.ServiceName = name
	err := otela.Configure(wg, done, cfg)
	if err != nil {
		panic(err)
	}
	tracer := otela.NewTracer("test-tracer2")

	_, span := tracer.Start(ctx, "test-span2")

	defer span.End()

	sendMetric()

	span.AddEvent("test-event2")
}

func sendMetric() {
	meter := otel.Meter("test.meter")
	counter, err := meter.Float64Counter("test.counter", metric.WithDescription("test.counter"))
	if err != nil {
		fmt.Println(err)
	} else {
		cv := 1.0
		counter.Add(context.Background(), cv)
	}
}
