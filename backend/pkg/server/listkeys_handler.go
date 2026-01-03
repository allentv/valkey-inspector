package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s Server) handleListKeys(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := s.getDefaultContext()
	defer cancel()

	var keys []string
	var cursor uint64
	for {
		scan, err := s.deps.ValkeyClient.Do(ctx, s.deps.ValkeyClient.B().Scan().Cursor(cursor).Match("*").Build()).AsScanEntry()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning keys: %v", err), http.StatusInternalServerError)
			return
		}
		keys = append(keys, scan.Elements...)
		cursor = scan.Cursor
		if cursor == 0 {
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(keys)
}
