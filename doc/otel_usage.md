# Observability, Trace, Metric

## Configuration
For use observability, trace or metric you need to configure it.  

In below, you can find sample code to configure Otel.

```go
package main

import (
	"sync"
	"github.com/ormushq/ormus/adapter/otela"
)

var cfg = otela.Config{
	Endpoint:           "otel_collector:4317",
	ServiceName:        "test",
	EnableMetricExpose: true,
	MetricExposePath:   "metrics",
	MetricExposePort:   8081,
	Exporter:           otela.ExporterGrpc,
}

func main()  {
	wg := &sync.WaitGroup{}
	done := make(chan bool)

	err := otela.Configure(wg, done, cfg)
	if err != nil {
		panic(err)
	}
}

```

### Specification
#### Endpoint
The endpoint of otel collector
#### ServiceName
The current service name
#### EnableMetricExpose
If this service need to expose default metrics to collect them in prometheus set this to `true`
#### MetricExposePath
If `EnableMetricExpose` is `true`, the path to expose metrics defined in the `MetricExposePath
#### MetricExposePort
If `EnableMetricExpose` is `true`, the port to expose metrics defined in the `MetricExposePort
#### Exporter
This used for configure exporter mod. Currently two mode supported: `otela.ExporterGrpc`,`otela.ExporterConsole`


## Observability/Trace

```go
package main

import (
	"context"
	"github.com/ormushq/ormus/adapter/otela"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	tracer := otela.NewTracer("test-tracer")
	ctx, span := tracer.Start(context.Background(), "main", trace.WithAttributes(
		attribute.String("detail", "boot-application-trace"),
	))
	span.AddEvent("processed-events-channel-opened")
}
```
You can add more event. after done work you can close span
```go
span.End()
```
If you want to trace in other packages you pass context to it and create new span with this context.

There is two function in `otela` package to get carrier from context (`otela.GetCarrierFromContext`) and convert carrier to the context (`otela.GetContextFromCarrier`).
These methods used when you want to pass current span to another process as parent span.(multi service trace)

## Metrics

For metric you can use this:
```go
meter := otel.Meter("test-meter")
counter, _ := meter.Float64Counter("test_counter")
counter.Add(context.Background(), 1)
```

There is two helper method: `otela.AddFloat64Counter` and `otela.IncrementFloat64Counter` 