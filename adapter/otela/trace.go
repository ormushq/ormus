package otela

import (
	"context"
	"fmt"

	"github.com/ormushq/ormus/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"go.opentelemetry.io/otel/trace"
)

func (opr *otelProvider) newTraceExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	switch opr.exporter {
	case ExporterGrpc:
		return otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithGRPCConn(opr.conn))
	case ExporterConsole:
		return stdouttrace.New()
	default:
		panic("unsupported")
	}
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

func NewTracer(name string, options ...trace.TracerOption) trace.Tracer {
	if !op.isConfigure {
		panic("You must configure adapter before calling NewTracer")
	}

	return op.tracerProvider.Tracer(name, options...)
}

func GetCarrierFromContext(ctx context.Context) map[string]string {
	propgator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})

	carrier := propagation.MapCarrier{}
	propgator.Inject(ctx, carrier)

	return carrier
}

func GetContextFromCarrier(carrier map[string]string) context.Context {
	propgator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	c := propagation.MapCarrier{}
	for k, v := range carrier {
		c.Set(k, v)
	}
	parentCtx := propgator.Extract(context.Background(), c)

	return parentCtx
}
