package config

// Environment directly fetch os envs with getting help from env
type Environment struct {
	Server
	Setting
	Database
	Log
}

// Server hold gin and tls configuration
type Server struct {
	Port     string `env:"OMEGA_SERVER_PORT"`
	ADDR     string `env:"OMEGA_SERVER_ADDR"`
	TLSKey   string `env:"OMEGA_TLS_KEY"`
	TLSCert  string `env:"OMEGA_TLS_CERT"`
	TimeZone string `env:"OMEGA_TIME_ZONE"`
}

// Setting hold pass-keys and JWT, used for security
type Setting struct {
	PasswordSalt  string `env:"OMEGA_PASSWORD_SALT"`
	AutoMigrate   bool   `env:"OMEGA_AUTO_MIGRATE"`
	JWTSecretKey  string `env:"OMEGA_JWT_SECRET_KEY,required"`
	JWTExpiration int    `env:"OMEGA_JWT_EXPIRATION,required"`
}

// Database hold DB connections, in case we just have one database use same DSN for both
type Database struct {
	Data struct {
		DSN  string `env:"OMEGA_DATABASE_DATA_URL,required"`
		Type string `env:"OMEGA_DATABASE_DATA_TYPE,required"`
	}
	Activity struct {
		DSN  string `env:"OMEGA_DATABASE_ACTIVITY_URL,required"`
		Type string `env:"OMEGA_DATABASE_ACTIVITY_TYPE,required"`
	}
}

// Log configuration terms hold here
type Log struct {
	ServerLog struct {
		Format     string `env:"OMEGA_SERVER_LOG_FORMAT,required"`
		Output     string `env:"OMEGA_SERVER_LOG_OUTPUT,required"`
		Level      string `env:"OMEGA_SERVER_LOG_LEVEL,required"`
		JSONIndent bool   `env:"OMEGA_SERVER_LOG_JSON_INDENT,required"`
	}

	ApiLog struct {
		Format     string `env:"OMEGA_API_LOG_FORMAT,required"`
		Output     string `env:"OMEGA_API_LOG_OUTPUT,required"`
		Level      string `env:"OMEGA_API_LOG_LEVEL,required"`
		JSONIndent bool   `env:"OMEGA_API_LOG_JSON_INDENT,required"`
	}
}
