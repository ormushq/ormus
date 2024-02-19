package rabbitmq

// AMQPConfig holds configuration options for connecting to RabbitMQ using AMQP URI.
type AMQPConfig struct {
	Username    string
	Password    string
	Hostname    string
	Port        int
	VirtualHost string
}

// DefaultAMQPConfig creates a new instance of AMQPConfig with default values.
func DefaultAMQPConfig() *AMQPConfig {
	return &AMQPConfig{
		Username:    "guest",
		Password:    "guest",
		Hostname:    "localhost",
		Port:        5672,
		VirtualHost: "/",
	}
}

// NEWAMQPConfig generates the AMQPConfig from the AMQPConfig.
func NEWAMQPConfig(config *AMQPConfig) *AMQPConfig {
	return &AMQPConfig{
		Username:    config.Username,
		Password:    config.Password,
		Hostname:    config.Hostname,
		Port:        config.Port,
		VirtualHost: config.VirtualHost,
	}
}
