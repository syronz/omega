package config

// Environment directly fetch os envs with getting help from env
type Environment struct {
	Server struct {
		Port string `env:"OMEGA_SERVER_PORT"`
		ADDR string `env:"OMEGA_SERVER_ADDR"`
	}
	Database struct {
		URL  string `env:"OMEGA_DATABASE_URL,required"`
		Type string `env:"OMEGA_DATABASE_TYPE,required"`
	}
	Log struct {
		Format string `env:"OMEGA_LOG_FORMAT,required"`
		Output string `env:"OMEGA_LOG_OUTPUT,required"`
		Level  string `env:"OMEGA_LOG_LEVEL,required"`
	}
	Logapi struct {
		Format string `env:"OMEGA_LOGAPI_FORMAT,required"`
		Output string `env:"OMEGA_LOGAPI_OUTPUT,required"`
		Level  string `env:"OMEGA_LOGAPI_LEVEL,required"`
	}
	Setting struct {
		PasswordSalt string `env:"OMEGA_PASSWORD_SALT"`
		AutoMigrate  string `env:"OMEGA_AUTOMIGRATE"`
		JWTSecretKey string `env:"OMEGA_JWT_SECRET_KEY,required"`
	}
}
