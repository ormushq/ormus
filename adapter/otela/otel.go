package otela

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var op otelProvider

func Configure(wg *sync.WaitGroup, done <-chan bool, cfg Config) error {
	if cfg.EnableMetricExpose && !op.isConfigure {
		serveMetric(cfg, wg, done)
	}

	conn, err := initConn(cfg)
	if err != nil {
		return err
	}

	op = otelProvider{
		wg:             wg,
		done:           done,
		conn:           conn,
		isConfigure:    true,
		tracerProvider: nil,
		metricProvider: nil,
		serviceName:    cfg.ServiceName,
		exporter:       cfg.Exporter,
	}
	err = op.initTrace()
	if err != nil {
		return err
	}

	err = op.initMetric()
	if err != nil {
		return err
	}

	return nil
}

func serveMetric(cfg Config, wg *sync.WaitGroup, done <-chan bool) {
	fmt.Printf("Metric enabled on port %d and path: %s \n", cfg.MetricExposePort, cfg.MetricExposePath)
	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", cfg.MetricExposePort),
		ReadTimeout: time.Second,
	}

	http.Handle(fmt.Sprintf("/%s", cfg.MetricExposePath), promhttp.Handler())

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-done
		timeout := 5
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
}

func initConn(cfg Config) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(cfg.Endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	return conn, err
}

type Exporter string

const (
	EXPORTER_GRPC    = Exporter("grpc")
	EXPORTER_CONSOLE = Exporter("console")
)

type otelProvider struct {
	wg             *sync.WaitGroup
	done           <-chan bool
	conn           *grpc.ClientConn
	isConfigure    bool
	tracerProvider trace.TracerProvider
	metricProvider metric.MeterProvider
	serviceName    string
	exporter       Exporter
}

type Config struct {
	Endpoint           string
	ServiceName        string
	EnableMetricExpose bool
	MetricExposePath   string
	MetricExposePort   int
	Exporter           Exporter
}
