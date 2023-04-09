package config

import (
	"flag"
	"fmt"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"

	"github.com/knadh/koanf/v2"
)

type Config struct {
	Port int32
}

const (
	DevEnv = "development"
)

const (
	APPLICATION_PORT = "APPLICATION_PORT"
)

func New() *Config {
	var k = koanf.New(".")
	var runtimeEnv string
	flag.StringVar(&runtimeEnv, "env", DevEnv, "Environment")
	flag.Parse()

	err := k.Load(file.Provider(fmt.Sprintf(".env.%s", runtimeEnv)), dotenv.Parser())
	if err != nil {
		panic(err)
	}

	return &Config{
		Port: int32(k.Int64(APPLICATION_PORT)),
	}
}
