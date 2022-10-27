package config

import (
	"fmt"

	const_app_env "github.com/infranyx/go-grpc-template/constant/app_env"
	"github.com/infranyx/go-grpc-template/utils"
	"github.com/joho/godotenv"
)

var Conf *Config

type Config struct {
	App  AppConfig
	Grpc GrpcConfig
}

type AppConfig struct {
	AppEnv string
}
type GrpcConfig struct {
	Port int
	Host string
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
			AppEnv: (utils.GetEnv("APP_ENV", const_app_env.DEV)),
		},
		Grpc: GrpcConfig{
			Port: utils.GetEnvAsInt("PORT", 3000),
			Host: utils.GetEnv("HOST", "0.0.0.0"),
		},
	}
	return Conf
}
