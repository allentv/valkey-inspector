package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"valkey-inspector/backend/pkg/keys"
)

func (s Server) handleGetKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := s.getDefaultContext()
	defer cancel()

	key := r.PathValue("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	keyType, err := s.deps.ValkeyClient.Do(ctx, s.deps.ValkeyClient.B().Type().Key(key).Build()).ToString()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting type: %v", err), http.StatusInternalServerError)
		return
	}

	var value interface{}
	var valErr error

	switch keyType {
	case "string":
		value, valErr = s.deps.ValkeyClient.Do(ctx, s.deps.ValkeyClient.B().Get().Key(key).Build()).ToString()
	case "list":
		value, valErr = s.deps.ValkeyClient.Do(ctx, s.deps.ValkeyClient.B().Lrange().Key(key).Start(0).Stop(-1).Build()).AsStrSlice()
	case "set":
		value, valErr = s.deps.ValkeyClient.Do(ctx, s.deps.ValkeyClient.B().Smembers().Key(key).Build()).AsStrSlice()
	case "hash":
		value, valErr = s.deps.ValkeyClient.Do(ctx, s.deps.ValkeyClient.B().Hgetall().Key(key).Build()).AsStrMap()
	case "zset":
		// Return with scores for completeness
		value, valErr = s.deps.ValkeyClient.Do(ctx, s.deps.ValkeyClient.B().Zrange().Key(key).Min("0").Max("-1").Withscores().Build()).AsZScores()
	case "none":
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	default:
		value = fmt.Sprintf("Unsupported type: %s", keyType)
	}

	if valErr != nil {
		http.Error(w, fmt.Sprintf("Error getting value: %v", valErr), http.StatusInternalServerError)
		return
	}

	response := keys.KeyDetail{
		Key:   key,
		Type:  keyType,
		Value: value,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
