package rabbitmq

type Config struct {
	UserName        string `koanf:"username"`
	Password        string `koanf:"password"`
	Host            string `koanf:"host"`
	Port            int    `koanf:"port"`
	Vhost           string `koanf:"vhost"`
	ReconnectSecond int    `koanf:"reconnect_second"`
}
