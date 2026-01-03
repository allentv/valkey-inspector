package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	cfg  Config
	deps Dependencies
}

// New will create a new server instance with configuration and dependencies
func New() Server {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	deps, err := NewDependencies(cfg)
	if err != nil {
		panic(err)
	}
	return Server{
		cfg:  cfg,
		deps: deps,
	}
}

func (s Server) Start() {
	mux := s.createRoutes()

	// Middleware for CORS
	handler := s.enableCORS(mux)

	log.Println("Starting Valkey Inspector Backend on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func (s Server) getDefaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
