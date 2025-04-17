package web

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"error"`
}

func errorResponse(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	jsonResponse(w, ErrorResponse{msg})
}

func jsonResponse(w http.ResponseWriter, val any) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(val)
	if err != nil {
		slog.Error(fmt.Sprintf("json.NewEncoder: %v", err))
	}
}
