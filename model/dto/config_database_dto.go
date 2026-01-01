package dto

type DatabaseConfig struct {
	Name     string `env:"DB_SCHEMA" default:"early_setup"`
	Adapter  string `env:"DB_DRIVER" default:"mysql"`
	Host     string `env:"DB_HOST" default:"localhost"`
	Port     string `env:"DB_PORT" default:"8889"`
	User     string `env:"DB_USER" default:"root"`
	Password string `env:"DB_PASSWORD" default:"root"`
	SslMode  string `env:"DB_SSL_MODE"`
}
