package server

import (
	"fmt"
	"os"
)

type Config struct {
	ValkeyAddr string
}

func LoadConfig() (Config, error) {
	// Read from all configuration paths: env vars and files
	addr, present := os.LookupEnv("VALKEY_ADDR")
	if !present {
		return Config{}, fmt.Errorf("VALKEY_ADDR not set")
	}

	return Config{
		ValkeyAddr: addr,
	}, nil
}
