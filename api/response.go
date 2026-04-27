package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func (s *Server) jsonResponse(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(&APIResponse{
		Success: true,
		Data:    data,
	})
}

func (s *Server) errorResponse(w http.ResponseWriter, status int, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(&APIResponse{
		Success: false,
		Error:   http.StatusText(status),
		Message: message,
	})
}

func (s *Server) parseID(path, prefix string) string {
	if !strings.HasPrefix(path, prefix) {
		return ""
	}

	id := strings.TrimPrefix(path, prefix)
	// Remove any trailing path segments (in case of nested routes)
	if idx := strings.Index(id, "/"); idx != -1 {
		id = id[:idx]
	}

	id = strings.TrimSpace(id)
	if id == "" {
		return ""
	}

	return id
}
