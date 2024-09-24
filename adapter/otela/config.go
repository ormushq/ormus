package otela

type Otel struct {
	Endpoint           string `koanf:"endpoint"`
	ServiceName        string `koanf:"service_name"`
	EnableMetricExpose bool   `koanf:"enable_metric_expose"`
	MetricExposePort   int    `koanf:"metric_expose_port"`
	MetricExposePath   string `koanf:"metric_expose_path"`
}
