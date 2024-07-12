package etcd

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/ormushq/ormus/adapter/otela"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Config struct {
	Host               string `koanf:"host"`
	Port               int    `koanf:"port"`
	DialTimeoutSeconds uint8  `koanf:"dial_timeout"`
}

type Adapter struct {
	client *clientv3.Client
}

func New(config Config) (Adapter, error) {
	return NewWithContext(context.Background(), config)
}

func NewWithContext(ctx context.Context, config Config) (Adapter, error) {
	tracer := otela.NewTracer("etcd")
	_, span := tracer.Start(ctx, "etcd@NewWithContext", trace.WithAttributes(
		attribute.String("config", fmt.Sprintf("%+v", config))))
	defer span.End()

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", config.Host, config.Port)},
		DialTimeout: time.Duration(config.DialTimeoutSeconds) * time.Second,
	})
	if err != nil {
		slog.Error("Error creating etcd client: ", err)
		span.AddEvent("error-on-connect-to-etcd", trace.WithAttributes(
			attribute.String("error", err.Error())))

		return Adapter{}, err
	}

	span.AddEvent("connected-to-etcd-successfully")

	return Adapter{
		client: etcdClient,
	}, nil
}

func (a Adapter) Client() *clientv3.Client {
	return a.client
}

func (a Adapter) Close() error {
	return a.client.Close()
}

func (a Adapter) Lock(ctx context.Context, key string, ttl int64) (unlock func() error, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(ttl))

	tracer := otela.NewTracer("etcd")
	ctx, span := tracer.Start(ctx, "etcd@Lock", trace.WithAttributes(
		attribute.String("key", key)))
	defer span.End()

	session, err := concurrency.NewSession(a.client, concurrency.WithTTL(int(ttl)), concurrency.WithContext(ctx))
	if err != nil {
		span.AddEvent("error-on-new-session", trace.WithAttributes(
			attribute.String("error", err.Error())))
		cancel()

		return nil, err
	}

	mutex := concurrency.NewMutex(session, key)
	if err := mutex.Lock(ctx); err != nil {
		span.AddEvent("error-on-new-mutex", trace.WithAttributes(
			attribute.String("error", err.Error())))
		cancel()

		return nil, err
	}

	span.AddEvent("key-locked")

	return func() error {
		err = mutex.Unlock(ctx)
		cancel()

		return err
	}, nil
}
