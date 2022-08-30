package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	GetServiceHost string
	GetServicePort int

	CRUDServiceHost string
	CRUDServicePort int

	HTTPPort string
}

func Load() Config {
	c := Config{}

	c.HTTPPort = cast.ToString(look(`HTTP_PORT`, `:9999`))

	c.GetServiceHost = cast.ToString(look(`GET_SERVICE_HOST`, `127.0.0.1`))
	c.GetServicePort = cast.ToInt(look(`GET_SERVICE_PORT`, 8000))

	c.CRUDServiceHost = cast.ToString(look(`CRUD_SERVICE_HOST`, `127.0.0.1`))
	c.CRUDServicePort = cast.ToInt(look(`CRUD_SERVICE_PORT`, 9000))

	return c
}

func look(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
