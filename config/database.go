package config

import "os"

type DatabaseConfig struct {
	User     string
	Password string
	Driver   string
	Name     string
	Host     string
	Port     string
}

func LoadDatabaseConfig() DatabaseConfig{
	return DatabaseConfig{
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
		Name:     os.Getenv("DB_DATABASE"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
}
