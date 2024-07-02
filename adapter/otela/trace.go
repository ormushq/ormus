package otela

import (
	"context"
	"fmt"

	"github.com/ormushq/ormus/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.25.0"
	"go.opentelemetry.io/otel/trace"
)

func (opr *otelProvider) newTraceExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	return otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithGRPCConn(opr.conn))
}

func (opr *otelProvider) newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	// Ensure default SDK resources and the required service name are set.
	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(op.serviceName),
		)),
	)
}

func (opr *otelProvider) initTrace() error {
	ctx := context.Background()

	exp, err := opr.newTraceExporter(ctx)
	if err != nil {
		return err
	}

	tp := opr.newTraceProvider(exp)

	otel.SetTracerProvider(tp)
	opr.wg.Add(1)
	go func() {
		defer opr.wg.Done()
		<-opr.done
		err = tp.Shutdown(ctx)
		if err != nil {
			logger.L().Error(fmt.Sprintf("failed to shutdown tracer: %v", err))
		}
	}()
	otel.SetTracerProvider(tp)
	opr.tracerProvider = tp

	return nil
}

func NewTracer(name string) trace.Tracer {
	if !op.isConfigure {
		panic("You must configure adapter before calling NewTracer")
	}

	return op.tracerProvider.Tracer(name)
}
