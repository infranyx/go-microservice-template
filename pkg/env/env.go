package env

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type EVar struct {
	key        string
	defaultVal interface{}
}

func init() {
	LoadEnv()
}

func LoadEnv() {
	_, callerDir, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Error generating env dir")
	}

	dir := filepath.Join(filepath.Dir(callerDir), "../..", "envs/.env")
	err := godotenv.Load(dir)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func New(key string, defaultVal interface{}) *EVar {
	return &EVar{key: key, defaultVal: defaultVal}
}

func (eVar EVar) GetEnv() interface{} {
	if val, exists := os.LookupEnv(eVar.key); exists {
		return val
	}

	if eVar.defaultVal == nil {
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
		log.Fatalf("could not convert eVar to Int %v", eVar.key)
	}

	return val
}

func (eVar EVar) AsBool() bool {
	val, err := strconv.ParseBool(eVar.AsString())
	if err != nil {
		log.Fatalf("could not convert eVar to bool %v", eVar.key)
	}

	return val
}

func (eVar EVar) AsStringSlice(sep string) []string {
	valStr := eVar.AsString()

	return strings.Split(valStr, sep)
}
