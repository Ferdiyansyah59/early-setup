package dto

type ServerConfig struct {
	GinMode     string `env:"GIN_MODE"`
	ServiceHost string `env:"SERVICE_HOST"`
	ServicePort string `env:"SERVICE_PORT"`
	Database    DatabaseConfig
}
  