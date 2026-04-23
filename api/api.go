package api

import (
	"log/slog"
	"net/http"
)

type Server struct {
	mux    *http.ServeMux
	logger *slog.Logger
}

func NewServer(logger *slog.Logger) *Server {
	mux := http.NewServeMux()

	s := &Server{
		mux:    mux,
		logger: logger,
	}

	mux.HandleFunc("/health", s.HandleHealth)

	return s
}
func (s *Server) Router() http.Handler {
	return s.mux
}
