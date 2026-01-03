package server

import "net/http"

func (s Server) createRoutes() *http.ServeMux {
	// Define routes
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/connect", s.handleConnect)
	mux.HandleFunc("GET /api/keys", s.handleListKeys)
	mux.HandleFunc("GET /api/keys/{key}", s.handleGetKey)

	return mux
}
