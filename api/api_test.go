package api

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServer(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)

	if server == nil {
		t.Error("expected server to be created")
	}

	if server.mux == nil {
		t.Error("expected mux to be initialized")
	}

	if server.logger == nil {
		t.Error("expected logger to be set")
	}
}

func TestServerRouter(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)

	router := server.Router()
	if router == nil {
		t.Error("expected router to be returned")
	}

}

func TestHealthEndpoint(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestAllRoutesRegistered(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	testRoutes := []struct {
		method string
		path   string
	}{
		{"POST", "/api/v1/auth/register"},
		{"POST", "/api/v1/auth/login"},
		{"POST", "/api/v1/auth/logout"},
		{"GET", "/api/v1/products"},
		{"GET", "/api/v1/products/123"},
		{"GET", "/api/v1/customers"},
		{"GET", "/api/v1/customers/456"},
		{"GET", "/api/v1/orders"},
		{"GET", "/api/v1/orders/789"},
		{"GET", "/api/v1/cart/cart1"},
		{"POST", "/api/v1/cart/cart1/items"},
		{"DELETE", "/api/v1/cart/cart1/items/item1"},
		{"POST", "/api/v1/checkout"},
		{"GET", "/api/v1/collections"},
		{"GET", "/api/v1/collections/col1"},
	}

	for _, route := range testRoutes {
		t.Run(route.method+" "+route.path, func(t *testing.T) {
			req, err := http.NewRequest(route.method, ts.URL+route.path, nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("failed to make request: %v", err)
			}
			defer resp.Body.Close()

			// Just verify we get a response (not 404)
			if resp.StatusCode == http.StatusNotFound {
				t.Errorf("route %s %s returned 404", route.method, route.path)
			}
		})
	}
}
