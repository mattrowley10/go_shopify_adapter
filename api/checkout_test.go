package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mattrowley10/the_faywood_adapter/internal/core/domain/types"
)

func TestCreateCheckout(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	t.Run("valid checkout", func(t *testing.T) {
		checkout := types.Checkout{
			CartID: "cart_123",
		}
		body, _ := json.Marshal(checkout)

		resp, err := http.Post(ts.URL+"/api/v1/checkout", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected status 201, got %d", resp.StatusCode)
		}

		var result APIResponse
		json.NewDecoder(resp.Body).Decode(&result)
		if !result.Success {
			t.Error("expected success to be true")
		}
	})

	t.Run("missing cart ID", func(t *testing.T) {
		checkout := types.Checkout{}
		body, _ := json.Marshal(checkout)

		resp, err := http.Post(ts.URL+"/api/v1/checkout", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", resp.StatusCode)
		}
	})

	t.Run("invalid method", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/api/v1/checkout")
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, got %d", resp.StatusCode)
		}
	})
}
