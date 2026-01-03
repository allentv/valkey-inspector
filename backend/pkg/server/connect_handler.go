package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s Server) handleConnect(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := s.getDefaultContext()
	defer cancel()

	err := s.deps.ValkeyClient.Do(ctx, s.deps.ValkeyClient.B().Ping().Build()).Error()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "connected"})
}
