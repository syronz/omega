package config

type Environment struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT"`
		Host string `envconfig:"SERVER_HOST"`
	}
	Database struct {
		URL  string `envconfig:"DATABASE_URL"`
		Type string `envconfig:"DATABASE_Type"`
	}
	Log struct {
		Format string `envconfig:"LOG_FORMAT"`
		Output string `envconfig:"LOG_OUTPUT"`
		Level  string `envconfig:"LOG_Level"`
	}
}
