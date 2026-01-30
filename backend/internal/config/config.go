package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigValues struct {
	Conn_String string
	JwtSecret   string
	JwtIssuer   string
	JwtTTLMin   int
}

func LoadConfig() (ConfigValues, error) {
	_ = godotenv.Load()

	connStr := os.Getenv("CONN_STR")
	if connStr == "" {
		err1 := errors.New("Connection string is empty")
		return ConfigValues{}, err1
	}

	ttlMinutesConfig := os.Getenv("JWT_TTL_MINUTES")
	ttlMinutes, err := strconv.ParseInt(ttlMinutesConfig, 10, 64)
	if err != nil {
		ttlMinutes = 60 // Default to 60 minutes
	}

	config := ConfigValues{
		Conn_String: os.Getenv("CONN_STR"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
		JwtIssuer:   os.Getenv("JWT_ISSUER"),
		JwtTTLMin:   int(ttlMinutes), // Default to 60 minutes or load from env if needed
	}

	return config, nil
}
