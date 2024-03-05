package trace

import (
	"context"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var Trace trace

type trace struct {
}

var tracerExp *otlptrace.Exporter

func (t *trace) RetryInitTracer() func() {
	var shutdown func()
	go func() {
		for {
			// otel will reconnected and re-send spans when otel col recover. so, we don't need to re-init tracer exporter.
			if tracerExp == nil {
				shutdown = t.InitTracer()
			} else {
				break
			}
			time.Sleep(time.Minute * 5)
		}
	}()
	return shutdown
}

func (t *trace) InitTracer() func() {
	// temporarily set timeout to 10s
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serviceName, ok := os.LookupEnv("OTEL_SERVICE_NAME")
	if !ok {
		serviceName = "gin-otel-demo"
		os.Setenv("OTEL_SERVICE_NAME", serviceName)
	}
	otelAgentAddr, ok := os.LookupEnv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if !ok {
		otelAgentAddr = "tempo.monitoring:4317"
		os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", otelAgentAddr)
	}
	zap.S().Infof("OTLP Trace connect to: %s with service name: %s", otelAgentAddr, serviceName)

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithDialOption(grpc.WithBlock()))
	if err != nil {
		t.HandleErr(err, "OTLP Trace gRPC Creation")
		return nil
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL)))

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	tracerExp = traceExporter
	return func() {
		// Shutdown will flush any remaining spans and shut down the exporter.
		t.HandleErr(tracerProvider.Shutdown(ctx), "failed to shutdown TracerProvider")
	}
}

func (t *trace) HandleErr(err error, message string) {
	if err != nil {
		zap.S().Errorf("%s: %v", message, err)
	}
}
