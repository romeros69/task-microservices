package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

const (
	host     = "rc1b-szajh704ic3nig49.mdb.yandexcloud.net"
	port     = 6432
	user     = "evgen"
	password = "romahach"
	dbname   = "tasks"
)

type Config struct {
	AppPort      string `env:"APP_PORT" envDefault:"9000"`
	PostgresUrl  string
	SecretKey    string `env:"SECRET_KEY" envDefault:"SOME_SECRET_KEY"`
	GrpcProtocol string `env:"GRPC_PROT" envDefault:"tcp"`
	GrpcURL      string `env:"GRPC_URL" envDefault:"auth:9011"`
}

func NewConfig() (*Config, error) {
	connstring := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=verify-full target_session_attrs=read-write",
		host, port, dbname, user, password)
	cfg := &Config{}
	cfg.PostgresUrl = connstring
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
