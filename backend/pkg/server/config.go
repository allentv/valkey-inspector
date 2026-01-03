package server

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	ValkeyAddr string
}

func LoadConfig() (Config, error) {
	// Read from all configuration paths: env vars and files
	valkeyAddr, present := os.LookupEnv("VALKEY_ADDR")
	if !present {
		return Config{}, fmt.Errorf("VALKEY_ADDR not set")
	}

	serverAddr := os.Getenv("SERVER_PORT")
	if serverAddr == "" {
		serverAddr = ":8080"
		log.Println("Server port not set. Using default of :8080")
	}

	return Config{
		ValkeyAddr: valkeyAddr,
	}, nil
}
