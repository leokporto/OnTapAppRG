package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type ConfigValues struct {
	Conn_String string
}

func LoadConfig() (ConfigValues, error) {
	_ = godotenv.Load()

	connStr := os.Getenv("CONN_STR")
	if connStr == "" {
		err1 := errors.New("Connection string is empty")
		return ConfigValues{}, err1
	}

	config := ConfigValues{
		Conn_String: os.Getenv("CONN_STR"),
	}

	return config, nil
}
