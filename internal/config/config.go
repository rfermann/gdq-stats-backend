package config

import (
	"fmt"
	"os"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"

	"github.com/knadh/koanf/v2"
)

type Config struct {
	Database_Url string
	Environment  string
	Port         int32
}

const (
	DevEnv = "development"
)

const (
	APPLICATION_PORT = "APPLICATION_PORT"
	DATABASE_URL     = "DATABASE_URL"
)

func New() *Config {
	var k = koanf.New(".")
	runtimeEnv := os.Getenv("RUNTIME_ENV")
	if len(runtimeEnv) == 0 {
		runtimeEnv = DevEnv
	}

	if runtimeEnv == DevEnv {
		err := k.Load(file.Provider(fmt.Sprintf(".env.%s", runtimeEnv)), dotenv.Parser())
		if err != nil {
			panic(err)
		}
	} else {
		k.Load(env.Provider("", ".", func(s string) string {
			return s
		}), nil)
	}

	return &Config{
		Database_Url: k.String(DATABASE_URL),
		Environment:  runtimeEnv,
		Port:         int32(k.Int64(APPLICATION_PORT)),
	}
}
