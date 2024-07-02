package otela

import (
	"context"
	"fmt"

	"github.com/ormushq/ormus/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/semconv/v1.25.0"
)

func (opr *otelProvider) newMetricExporter(ctx context.Context) (metric.Exporter, error) {
	return otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure(), otlpmetricgrpc.WithGRPCConn(opr.conn))
	// Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
}

func (opr *otelProvider) newMetricProvider(exp metric.Exporter) *metric.MeterProvider {
	// Ensure default SDK resources and the required service name are set.
	return metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(exp)),
		metric.WithResource(resource.NewWithAttributes(
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
