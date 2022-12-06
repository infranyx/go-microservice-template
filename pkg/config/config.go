package config

import (
	"github.com/infranyx/go-grpc-template/pkg/constant"
	"github.com/infranyx/go-grpc-template/pkg/env"
)

type Config struct {
	App                  AppConfig
	Grpc                 GrpcConfig
	Postgres             PostgresConfig
	GoTemplateGrpcClient GrpcConfig
}

var Conf *Config

type AppConfig struct {
	AppEnv string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Pass            string
	DBName          string
	MaxConn         int
	MaxIdleConn     int
	MaxLifeTimeConn int
	SslMode         string
}
type GrpcConfig struct {
	Port int
	Host string
}

func init() {
	Conf = newConfig()
}

func newConfig() *Config {
	return &Config{
		App: AppConfig{
			AppEnv: env.New("APP_ENV", constant.AppEnvDev).AsString(),
		},
		Grpc: GrpcConfig{
			Port: env.New("GRPC_PORT", constant.GrpcPort).AsInt(),
			Host: env.New("GRPC_HOST", constant.GrpcHost).AsString(),
		},
		Postgres: PostgresConfig{
			Host:            env.New("PG_HOST", constant.PgHost).AsString(),
			Port:            env.New("PG_PORT", constant.PgPort).AsString(),
			User:            env.New("PG_USER", constant.PgUser).AsString(),
			Pass:            env.New("PG_PASS", constant.PgPass).AsString(),
			DBName:          env.New("PG_DB", constant.PgDb).AsString(),
			MaxConn:         env.New("PG_MAX_CONNECTIONS", constant.PgMaxConn).AsInt(),
			MaxIdleConn:     env.New("PG_MAX_IDLE_CONNECTIONS", constant.PgMaxIdleConn).AsInt(),
			MaxLifeTimeConn: env.New("PG_MAX_LIFETIME_CONNECTIONS", constant.PgMaxLifeTimeConn).AsInt(),
			SslMode:         env.New("PG_SSL_MODE", constant.PgSslMode).AsString(),
		},
		GoTemplateGrpcClient: GrpcConfig{
			Port: env.New("GO_TEMPLATE_GRPC_PORT", constant.GrpcPort).AsInt(),
			Host: env.New("GO_TEMPLATE_GRPC_HOST", constant.GrpcHost).AsString(),
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
