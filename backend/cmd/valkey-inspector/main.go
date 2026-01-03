package main

import (
	"valkey-inspector/backend/pkg/server"
)

func main() {
	s := server.New()
	s.Start()
}
