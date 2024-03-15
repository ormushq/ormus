package dconfig

import "time"

type RedisTaskIdempotency struct {
	Host       string        `koanf:"host"`
	Port       int           `koanf:"port"`
	Password   string        `koanf:"password"`
	DB         int           `koanf:"db"`
	Prefix     string        `koanf:"prefix"`
	Expiration time.Duration `koanf:"expiration"`
}
