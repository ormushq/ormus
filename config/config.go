package config

import (
	"github.com/ormushq/ormus/adapter/etcd"
	rabbitmq "github.com/ormushq/ormus/adapter/rabbitmq"
	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/adapter/scylladb"
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/manager"
	"github.com/ormushq/ormus/source"
)

type Config struct {
	Redis       redis.Config    `koanf:"redis"`
	RabbitMQ    rabbitmq.Config `koanf:"rabbitmq"`
	Etcd        etcd.Config     `koanf:"etcd"`
	Manager     manager.Config  `koanf:"manager"`
	Source      source.Config   `koanf:"source"`
	Destination dconfig.Config  `koanf:"destination"`
	Scylladb    scylladb.Config `koanf:"scylladb"`
	Swagger     Swagger         `koanf:"swagger"`
}
