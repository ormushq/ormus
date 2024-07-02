package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
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
		EnableMetricExpose: true,
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
	go service2(wg, done, c)
	err := otela.Configure(wg, done, cfg)
	if err != nil {
		panic(err)
	}
	startWOrk(c)

	fmt.Println("Press Ctrl+C to stop")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	close(done)
	wg.Wait()
}

func service2(wg *sync.WaitGroup, done <-chan bool, c <-chan context.Context) {
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

	span.AddEvent("test-event2")
}

func startWOrk(c chan<- context.Context) {
	tracer := otela.NewTracer("test-tracer")

	ctx, span := tracer.Start(context.Background(), "test-span")

	defer span.End()

	span.AddEvent("test-event")

	doWork(ctx, tracer)

	meter := otel.Meter("test-meter")
	counter, err := meter.Float64Counter("test_counter")
	if err != nil {
		fmt.Println(err)
	} else {
		cv := 1.0
		counter.Add(context.Background(), cv)
	}

	meter = otel.Meter("test-meter2")
	counter1, err1 := meter.Float64Histogram("test_counter2")
	if err1 != nil {
		fmt.Println(err1)
	} else {
		cv := 10.0
		counter1.Record(context.Background(), cv)
	}

	meter = otel.Meter("test-meter3")
	counter, err = meter.Float64Counter("test_counter3")
	if err != nil {
		fmt.Println(err)
	} else {
		cv := 20.0
		counter.Add(context.Background(), cv)
	}

	c <- ctx
}

func doWork(ctx context.Context, tracer trace.Tracer) {
	_, span := tracer.Start(ctx, "doWork")
	defer span.End()

	span.AddEvent("Starting work")
	sleepTime := 2
	time.Sleep(time.Duration(sleepTime) * time.Second)

	span.AddEvent("Work completed")
}
