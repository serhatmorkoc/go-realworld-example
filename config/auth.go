package config

import "os"

type AuthConfig struct {
	AccessSecret  string
	RefreshSecret string
}

func LoadAuthConfig() AuthConfig{
	return AuthConfig{
		AccessSecret:  os.Getenv("JWT_ACCESS_SECRET"),
		RefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
	}
}