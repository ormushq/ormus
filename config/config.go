package config

import (
	"github.com/ormushq/ormus/adapter/etcd"
	"github.com/ormushq/ormus/adapter/rabbitmq"
	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/adapter/scylladb"
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/manager"
	"github.com/ormushq/ormus/source"
)

type Config struct {
	Redis       redis.Config    `koanf:"redis"`
	Etcd        etcd.Config     `koanf:"etcd"`
	RabbitMq    rabbitmq.Config `koanf:"rabbitmq"`
	Manager     manager.Config  `koanf:"manager"`
	Source      source.Config   `koanf:"source"`
	Destination dconfig.Config  `koanf:"destination"`
	Scylladb    scylladb.Config `koanf:"scylladb"`
	Swagger     Swagger         `koanf:"swagger"`
}
