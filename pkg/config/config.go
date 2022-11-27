package config

import (
	"github.com/infranyx/go-grpc-template/pkg/constant"
	"github.com/infranyx/go-grpc-template/pkg/env"
)

var Conf *Config

type Config struct {
	App      AppConfig
	Grpc     GrpcConfig
	Postgres PostgresConfig
}

type AppConfig struct {
	AppEnv string
}

type PostgresConfig struct {
	Host   string
	Port   int
	User   string
	Pass   string
	DBName string
}
type GrpcConfig struct {
	Port int
	Host string
}

func init() {
	Conf = NewConfig()
}

func NewConfig() *Config {
	return &Config{
		App: AppConfig{
			AppEnv: env.New("APP_ENV", constant.AppEnvDev).AsString(),
		},
		Grpc: GrpcConfig{
			Port: env.New("GRPC_PORT", constant.GrpcPort).AsInt(),
			Host: env.New("GRPC_HOST", constant.GrpcHost).AsString(),
		},
		Postgres: PostgresConfig{
			Host:   env.New("PG_HOST", constant.PgHost).AsString(),
			Port:   env.New("PG_PORT", constant.PgPort).AsInt(),
			User:   env.New("PG_USER", constant.PgUser).AsString(),
			Pass:   env.New("PG_PASS", constant.PgPass).AsString(),
			DBName: env.New("PG_DB", constant.PgDb).AsString(),
		},
	}
}

func IsDevEnv() bool {
	return Conf.App.AppEnv == constant.AppEnvDev
}

func IsProdEnv() bool {
	return Conf.App.AppEnv == constant.AppEnvProd
}

func IsTestEnv() bool {
	return Conf.App.AppEnv == constant.AppEnvTest
}
