package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// Config is config adpater redis.
type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Password string `koanf:"password"`
	DB       int    `koanf:"db"`
}

type Adapter struct {
	client *redis.Client
}

// New is Constructor redis Adapter.
func New(config Config) (Adapter, error) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
		/*	TODO -Is it necessary to add other Options? for example:
			DialTimeout:  config.DialTimeout,
			ReadTimeout:  config.ReadTimeout,
			WriteTimeout: config.WriteTimeout,
		*/
	})

	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return Adapter{}, err
	}

	return Adapter{client: redisClient}, nil

}

// Client is a method that returns the redis.Client.
func (a Adapter) Client() *redis.Client {
	return a.client
}
