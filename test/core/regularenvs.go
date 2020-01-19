package core

import "omega/config"

func getRegularEnvs() config.Environment {
	server := config.Server{
		Port: "8070",
		ADDR: "localhost",
	}

	setting := config.Setting{
		PasswordSalt:  "",
		JWTSecretKey:  "kz84HcnwKSn0k2vk6Ddw03kdck9k9SKedWFdGkwe70",
		JWTExpiration: 1000000,
	}

	database := config.Database{
		Data: config.Data{
			DSN:  "root@tcp(127.0.0.1:3306)/omega_test?charset=utf8&parseTime=True&loc=Local",
			Type: "mysql",
		},
		Activity: config.Activity{
			DSN:  "root@tcp(127.0.0.1:3306)/omega_test?charset=utf8&parseTime=True&loc=Local",
			Type: "mysql",
		},
	}

	log := config.Log{
		ServerLog: config.ServerLog{
			Format:     "json",
			Output:     "stdout",
			Level:      "trace",
			JSONIndent: true,
		},
		ApiLog: config.ApiLog{
			Format:     "json",
			Output:     "stdout",
			Level:      "trace",
			JSONIndent: true,
		},
	}

	return config.Environment{
		Server:   server,
		Setting:  setting,
		Database: database,
		Log:      log,
	}
}
