package handlers

import (
	"encoding/json"
	"net/http"
)

// Helper function for easily writing response messages
func writeResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
