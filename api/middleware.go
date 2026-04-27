package api

import (
	"net/http"
	"time"
)

func (s *Server) withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		s.logger.Info("handling request",
			"method", r.Method,
			"path", r.URL.Path,
		)

		next.ServeHTTP(w, r)

		s.logger.Info("request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}
