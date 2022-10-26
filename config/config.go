package config

import (
	"fmt"

	"github.com/infranyx/go-grpc-template/utils"
	"github.com/joho/godotenv"
)

var Conf *Config

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Port int
}

// init is invoked before main()
func init() {
	fmt.Println("env")
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

// New returns a new Config struct
func New() *Config {
	Conf = &Config{
		App: AppConfig{
			Port: utils.GetEnvAsInt("PORT", 3000),
		},
	}
	return Conf
}
