package config

type Swagger struct {
	Port    int           `koanf:"port"`
	Expose  bool          `koanf:"expose"`
	Manager SwaggerConfig `koanf:"manager"`
	Source  SwaggerConfig `koanf:"source"`
}
type SwaggerConfig struct {
	Version          string `koanf:"version"`
	Host             string `koanf:"host"`
	BasePath         string `koanf:"base_path"`
	Title            string `koanf:"title"`
	Description      string `koanf:"description"`
	InfoInstanceName string `koanf:"info_instance_name"`
}
