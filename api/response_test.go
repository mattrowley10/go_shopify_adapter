package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonResponse(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := &Server{
		mux:    http.NewServeMux(),
		logger: logger,
	}

	w := httptest.NewRecorder()
	testData := map[string]string{"key": "value"}
	server.jsonResponse(w, http.StatusOK, testData)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	if ct := w.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", ct)
	}

	var result APIResponse
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !result.Success {
		t.Error("expected success to be true")
	}
}

func TestErrorResponse(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := &Server{
		mux:    http.NewServeMux(),
		logger: logger,
	}

	w := httptest.NewRecorder()
	server.errorResponse(w, http.StatusBadRequest, "test error")

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}

	var result APIResponse
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result.Success {
		t.Error("expected success to be false")
	}

	if result.Message != "test error" {
		t.Errorf("expected message 'test error', got %s", result.Message)
	}
}

func TestParseID(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := &Server{
		mux:    http.NewServeMux(),
		logger: logger,
	}

	tests := []struct {
		path   string
		prefix string
		want   string
	}{
		{"/api/v1/products/123", "/api/v1/products/", "123"},
		{"/api/v1/products/123/variants", "/api/v1/products/", "123"},
		{"/api/v1/products/", "/api/v1/products/", ""},
		{"/api/v1/customers/cust_456", "/api/v1/customers/", "cust_456"},
		{"/wrong/path", "/api/v1/products/", ""},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			result := server.parseID(tt.path, tt.prefix)
			if result != tt.want {
				t.Errorf("parseID(%q, %q) = %q, want %q", tt.path, tt.prefix, result, tt.want)
			}
		})
	}
}
