package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListOrders(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	t.Run("valid list", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/api/v1/orders")
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
		req, _ := http.NewRequest("PUT", ts.URL+"/api/v1/orders", nil)
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

func TestGetOrder(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	t.Run("valid order ID", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/api/v1/orders/ord_123")
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
		req, _ := http.NewRequest("DELETE", ts.URL+"/api/v1/orders/ord_123", nil)
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
