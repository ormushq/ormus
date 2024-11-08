package source

import "github.com/ormushq/ormus/adapter/otela"

// HTTPServer is the main object for http configurations.
type HTTPServer struct {
	Port    int    `koanf:"port"`
	Network string `koanf:"network"`
}

// Config is the main object for managing source configuration.
type Config struct {
	HTTPServer HTTPServer `koanf:"http_server"`
	// TODO - add source, auth and etc configurations
	Otel                    otela.Otel `koanf:"otel"`
	WriteKeyRedisExpiration uint       `koanf:"write_key_expiration"`
	NewSourceEventName      string     `koanf:"new_source_event_name"`
	BufferSize              int        `koanf:"buffersize"`
	MaxRetry                int        `koanf:"maxretry"`
}
