package main

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	HTTP     serverConfig `envconfig:"HTTP"`
	Postgres dbConfig     `envconfig:"POSTGRES"`
	GCS      gcsConfig    `envconfig:"GCS"`
}

func configFromEnv() (*config, error) {
	var c config
	err := envconfig.Process("accountsd", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

type dbConfig struct {
	Host         string `envconfig:"HOST" default:"localhost"`
	Port         int    `envconfig:"PORT" default:"5433"`
	User         string `envconfig:"USER" default:"playj"`
	Password     string `envconfig:"PASSWORD" default:"playj1307"`
	DBName       string `envconfig:"DB_NAME" default:"playj-accounts-db"`
	MaxIdleConns int    `envconfig:"MAX_IDLE_CONNS" default:"16"`
	MaxOpenConns int    `envconfig:"MAX_OPEN_CONNS" default:"32"`
}

type serverConfig struct {
	Host string `envconfig:"HOST" default:"localhost"`
	Port int    `envconfig:"PORT" default:"8091"`
}

type gcsConfig struct {
	BucketURL string `envconfig:"GCS_BUCKET_URL"`
}
