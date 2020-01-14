package config

// Environment directly fetch os envs with getting help from env
type Environment struct {
	Server struct {
		Port string `env:"SERVER_PORT"`
		ADDR string `env:"SERVER_ADDR"`
	}
	Database struct {
		URL  string `env:"DATABASE_URL,required"`
		Type string `env:"DATABASE_TYPE,required"`
	}
	Log struct {
		Format string `env:"LOG_FORMAT,required"`
		Output string `env:"LOG_OUTPUT,required"`
		Level  string `env:"LOG_LEVEL,required"`
	}
	Logapi struct {
		Format string `env:"LOGAPI_FORMAT,required"`
		Output string `env:"LOGAPI_OUTPUT,required"`
		Level  string `env:"LOGAPI_LEVEL,required"`
	}
}
