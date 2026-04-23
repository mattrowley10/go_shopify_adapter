package api

import "net/http"

func (s *Server) HandleHealth(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("api is healthy")
	w.Write([]byte("api is healthy"))
}
