package config

type HTTPServer struct {
	Port    int    `koanf:"port"`
	Network string `koanf:"network"`
}

type Config struct {
	HTTPServer HTTPServer `koanf:"http_server"`
	// TODO - add source, auth and etc configurations
}