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


## Trace helper

### `GetCarrierFromContext(ctx context.Context) map[string]string`
To extract the carrier from the context, use this method. It's particularly useful when passing a tracer ID between services running in separate binaries, especially in event-driven or messaging-based design patterns. On the client service, you can easily set the tracer ID with this value, pass it to the host, and then use the following function to generate a context with the tracer ID for use in the host service.

### `GetContextFromCarrier(carrier map[string]string) context.Context`
As mentioned above, you can use this method to convert the tracer ID into a context.


## Trace builder

``go
TraceBuilder(packageName, functionName string, options ...TracerOptions) (context.Context, trace.Span)
``
You can use this method for easy integration with the Otel collector, like this:
```go
ctx, span := otela.TraceBuilder("package name", "Function name",
    TracerOptions...),
)
```
Here is a helper function for generating `TraceOption`, such as:
- `WithCarrier(carrier map[string]string) TracerOptions`
- `WithContext(ctx context.Context) TracerOptions`
- `WithTracerOptions(options ...trace.TracerOption) TracerOptions`
- `WithSpanOptions(options ...trace.SpanStartOption) TracerOptions`
- `WithSpanOptionAttributes(attributes ...attribute.KeyValue) TracerOptions`


## Log echo requests

In the delivery layer, you can register the Otel log as shown below:

```go
    s.Router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogHost:          true,
		LogRemoteIP:      true,
		LogRequestID:     true,
		LogMethod:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogLatency:       true,
		LogError:         true,
		LogProtocol:      true,
		LogValuesFunc:    otela.EchoRequestLoggerLogValuesFunc("httpserver", "Serve"),
	}))
```
It then starts the tracer on the request and injects the trace ID into the context. You can continue to append spans to the tracer as shown below:

```go
package statushandler

import (
	"github.com/ormushq/ormus/adapter/otela"
	"github.com/labstack/echo/v4"
)


    func (h Handler) healthCheck(c echo.Context) error {

        // You can pass newCtx to service for set this span as parent of spans that start in services 
        newCtx, span := otela.TraceBuilder("statushandler", "status",
        otela.WithContext(ctx.Request().Context()),
        )
        defer span.End()
        
        span.AddEvent("Starting status handler")
        // Do something
        span.AddEvent("Ending status handler")
		
        ...
    }
        
```