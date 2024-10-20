package otela

import (
	"context"
	"fmt"

	"github.com/ormushq/ormus/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric"
	sdkMetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"

	"go.opentelemetry.io/otel/trace"
)

func (opr *otelProvider) newMetricExporter(ctx context.Context) (sdkMetric.Exporter, error) {
	switch opr.exporter {
	case ExporterGrpc:
		return otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure(), otlpmetricgrpc.WithGRPCConn(opr.conn))
	case ExporterConsole:
		return stdoutmetric.New()
	default:
		panic("unsupported")
	}
}

func (opr *otelProvider) newMetricProvider(exp sdkMetric.Exporter) *sdkMetric.MeterProvider {
	// Ensure default SDK resources and the required service name are set.
	return sdkMetric.NewMeterProvider(
		sdkMetric.WithReader(sdkMetric.NewPeriodicReader(exp)),
		sdkMetric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(op.serviceName),
		)),
	)
}

func (opr *otelProvider) initMetric() error {
	ctx := context.Background()
	exp, err := opr.newMetricExporter(ctx)
	if err != nil {
		return err
	}
	mp := opr.newMetricProvider(exp)
	otel.SetMeterProvider(mp)

	op.wg.Add(1)
	go func() {
		defer opr.wg.Done()
		<-opr.done
		err = mp.Shutdown(ctx)
		if err != nil {
			logger.L().Error(fmt.Sprintf("failed to shutdown tracer: %v", err))
		}
	}()
	opr.metricProvider = mp

	return nil
}

func AddFloat64Counter(ctx context.Context, meter metric.Meter, name, desc string, cv float64, options ...metric.Float64CounterOption) {
	tracer := NewTracer("otela")
	ctx, span := tracer.Start(ctx, "otela@metric")
	options = append(options, metric.WithDescription(desc))

	processedEventCounter, err := meter.Float64Counter(name, options...)
	if err != nil {
		span.AddEvent("error on create counter", trace.WithAttributes(
			attribute.String("error", err.Error())))
		logger.L().Error("error on create counter", "error", err.Error())
	} else {
		processedEventCounter.Add(ctx, cv)
	}
}

func IncrementFloat64Counter(ctx context.Context, meter metric.Meter, name, desc string, options ...metric.Float64CounterOption) {
	cv := 1.0
	AddFloat64Counter(ctx, meter, name, desc, cv, options...)
}
