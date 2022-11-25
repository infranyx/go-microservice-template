package config

import (
	"fmt"

	const_app_env "github.com/infranyx/go-grpc-template/constant/app_env"
	"github.com/infranyx/go-grpc-template/utils"
	"github.com/joho/godotenv"
)

var Conf *Config

type Config struct {
	App      AppConfig
	Grpc     GrpcConfig
	Postgres PostgresConfig
	Kafka    KafkaConfig
}

type AppConfig struct {
	AppEnv string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	SSLMode  string
	Password string
}
type GrpcConfig struct {
	Port int
	Host string
}

type KafkaConfig struct {
	Network string
	Address string
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

// New returns a new Config struct
func NewConfig() *Config {
	Conf = &Config{
		App: AppConfig{
			AppEnv: utils.GetEnv("APP_ENV", const_app_env.DEV),
		},
		Grpc: GrpcConfig{
			Port: utils.GetEnvAsInt("PORT", 3000),
			Host: utils.GetEnv("HOST", "0.0.0.0"),
		},
		Postgres: PostgresConfig{
			Host:     utils.GetEnv("PG_HOST", "localhost"),
			Port:     utils.GetEnv("PG_PORT", "5432"),
			User:     utils.GetEnv("PG_USER", "postgres"),
			DBName:   utils.GetEnv("PG_DB", "postgres"),
			SSLMode:  utils.GetEnv("PG_SSL", "disable"),
			Password: utils.GetEnv("PG_PASS", "postgrespw"),
		},
		Kafka: KafkaConfig{
			Network: utils.GetEnv("KAFKA_NETWORK", "tcp"),
			Address: utils.GetEnv("KAFKA_ADDRESS", "localhost:9092"),
		},
	}
	return Conf
}

func IsDevelopment() bool {
	return Conf.App.AppEnv == const_app_env.DEV
}

func IsProduction() bool {
	return Conf.App.AppEnv == const_app_env.PROD
}
