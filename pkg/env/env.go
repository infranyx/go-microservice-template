package env

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const defaultEnvPath = "./envs/.env"

type EVar struct {
	key        string
	defaultVal interface{}
}

func New(key string, defaultVal interface{}) *EVar {
	return &EVar{key: key, defaultVal: defaultVal}
}

func init() {
	LoadEnv()
}

func LoadEnv() {
	// TODO : Log + Err
	err := godotenv.Load(defaultEnvPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func (eVar EVar) GetEnv() interface{} {
	if val, exists := os.LookupEnv(eVar.key); exists {
		return val
	}

	if eVar.defaultVal == nil {
		// TODO : Log + Err
		log.Fatalf("Env variable is required %v", eVar.key)
	}

	return eVar.defaultVal
}

func (eVar EVar) AsString() string {
	return fmt.Sprintf("%v", eVar.GetEnv())
}

func (eVar EVar) AsInt() int {
	val, err := strconv.Atoi(eVar.AsString())
	if err != nil {
		// TODO : Log + Err
		log.Fatalf("could not convert eVar to Int %v", eVar.key)
	}

	return val
}

func (eVar EVar) AsBool() bool {
	val, err := strconv.ParseBool(eVar.AsString())
	if err != nil {
		// TODO : Log + Err
		log.Fatalf("could not convert eVar to bool %v", eVar.key)
	}

	return val
}

func (eVar EVar) AsStringSlice(sep string) []string {
	valStr := eVar.AsString()

	val := strings.Split(valStr, sep)

	return val
}
