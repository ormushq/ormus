package etcd

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
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
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", config.Host, config.Port)},
		DialTimeout: time.Duration(config.DialTimeoutSeconds) * time.Second,
	})
	if err != nil {
		fmt.Println("Error creating client:", err)

		return Adapter{}, err
	}

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
	session, err := concurrency.NewSession(a.client, concurrency.WithTTL(int(ttl)))
	if err != nil {
		return nil, err
	}

	mutex := concurrency.NewMutex(session, key)
	if err := mutex.Lock(ctx); err != nil {
		return nil, err
	}

	return func() error {
		return mutex.Unlock(ctx)
	}, nil
}
