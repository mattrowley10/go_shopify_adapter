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

func TestGetCart(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	t.Run("valid cart ID", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/api/v1/cart/cart_123")
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
		req, _ := http.NewRequest("DELETE", ts.URL+"/api/v1/cart/cart_123", nil)
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

func TestAddToCart(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	t.Run("valid add to cart", func(t *testing.T) {
		item := types.Item{
			ProductID: "prod_123",
			VariantID: "var_123",
			Title:     "Test Product",
			Quantity:  1,
			Price: types.Money{
				Amount:   9999,
				Currency: "USD",
			},
		}
		body, _ := json.Marshal(item)

		resp, err := http.Post(ts.URL+"/api/v1/cart/cart_123/items", "application/json", bytes.NewBuffer(body))
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
		item := types.Item{
			Quantity: 1,
		}
		body, _ := json.Marshal(item)

		resp, err := http.Post(ts.URL+"/api/v1/cart/cart_123/items", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", resp.StatusCode)
		}
	})

	t.Run("invalid method", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/api/v1/cart/cart_123/items")
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, got %d", resp.StatusCode)
		}
	})
}

func TestRemoveFromCart(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	t.Run("valid remove from cart", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", ts.URL+"/api/v1/cart/cart_123/items/item_456", nil)
		resp, err := http.DefaultClient.Do(req)
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
		resp, err := http.Post(ts.URL+"/api/v1/cart/cart_123/items/item_456", "application/json", bytes.NewBuffer([]byte("")))
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, got %d", resp.StatusCode)
		}
	})
}
