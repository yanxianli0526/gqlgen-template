package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	Env struct {
		HTTPPort int  `envconfig:"HTTP_PORT"    default:"4000"`
		Debug    bool `envconfig:"DEBUG"`
	}

	Database struct {
		DBSSLMode  string `envconfig:"DB_SSL_MODE" default:"disable"`
		DBHost     string `envconfig:"DB_HOST" default:"localhost"`
		DBPort     string `envconfig:"DB_PORT" default:"5432"`
		DBname     string `envconfig:"DB_NAME" default:"inventory-toll"`
		DBPassword string `envconfig:"DB_PASSWORD" default:""`
		DBUser     string `envconfig:"DB_USER" default:"postgres"`
	}
	EnvConfig struct {
		Env      Env
		Database Database
	}
)

func Process(env *EnvConfig) (err error) {
	err = envconfig.Process("", env)
	return
}
