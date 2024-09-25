package otela

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type TracerOptions struct {
	ctx           context.Context
	tracerOptions []trace.TracerOption
	spanOptions   []trace.SpanStartOption
}

func TraceBuilder(packageName, functionName string, options ...TracerOptions) (context.Context, trace.Span) {
	traceOpt := TracerOptions{}
	for _, option := range options {
		if option.ctx != nil {
			traceOpt.ctx = option.ctx
		}
		if option.spanOptions != nil {
			traceOpt.spanOptions = option.spanOptions
		}
		if option.tracerOptions != nil {
			traceOpt.tracerOptions = option.tracerOptions
		}
	}
	if traceOpt.ctx == nil {
		traceOpt.ctx = context.Background()
	}
	tracer := NewTracer(packageName, traceOpt.tracerOptions...)

	return tracer.Start(traceOpt.ctx, packageName+"@"+functionName, traceOpt.spanOptions...)
}

func WithCarrier(carrier map[string]string) TracerOptions {
	return TracerOptions{
		ctx: GetContextFromCarrier(carrier),
	}
}

func WithContext(ctx context.Context) TracerOptions {
	return TracerOptions{
		ctx: ctx,
	}
}

func WithTracerOptions(options ...trace.TracerOption) TracerOptions {
	return TracerOptions{
		tracerOptions: options,
	}
}

func WithSpanOptions(options ...trace.SpanStartOption) TracerOptions {
	return TracerOptions{
		spanOptions: options,
	}
}

func WithSpanOptionAttributes(attributes ...attribute.KeyValue) TracerOptions {
	return TracerOptions{
		spanOptions: []trace.SpanStartOption{
			trace.WithAttributes(attributes...),
		},
	}
}
