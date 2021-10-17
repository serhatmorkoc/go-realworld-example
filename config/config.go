package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Auth     AuthConfig
	Database DatabaseConfig
	HTTP     HTTPConfig
	Redis    RedisConfig
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	return &Config{
		Auth:     LoadAuthConfig(),
		Database: LoadDatabaseConfig(),
		HTTP:     LoadHTTPConfig(),
		Redis:    LoadRedisConfig(),
	}
}
