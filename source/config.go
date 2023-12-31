package source

// HTTPServer is the main object for http configurations.
type HTTPServer struct {
	Port    int    `koanf:"port"`
	Network string `koanf:"network"`
}

// Config is the main object for managing source configuration.
type Config struct {
	HTTPServer HTTPServer `koanf:"http_server"`
	// TODO - add source, auth and etc configurations
}
