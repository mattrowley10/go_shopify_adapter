package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListProducts(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	t.Run("valid list", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/api/v1/products")
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}

		var result APIResponse
		json.NewDecoder(resp.Body).Decode(&result)
		if !result.Success {
			t.Error("expected success to be true")
		}
	})

	t.Run("invalid method", func(t *testing.T) {
		req, _ := http.NewRequest("POST", ts.URL+"/api/v1/products", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, got %d", resp.StatusCode)
		}
	})
}

func TestGetProduct(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	t.Run("valid product ID", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/api/v1/products/prod_123")
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}

		var result APIResponse
		json.NewDecoder(resp.Body).Decode(&result)
		if !result.Success {
			t.Error("expected success to be true")
		}
	})

	t.Run("missing product ID", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/api/v1/products/")
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// This might be a 404 since path doesn't match the pattern
		if resp.StatusCode == http.StatusOK {
			t.Error("expected non-200 status for missing ID")
		}
	})

	t.Run("invalid method", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", ts.URL+"/api/v1/products/prod_123", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, got %d", resp.StatusCode)
		}
	})
}
