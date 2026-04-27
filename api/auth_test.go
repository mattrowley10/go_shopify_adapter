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

func TestRegister(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	t.Run("valid registration", func(t *testing.T) {
		loginReq := types.LoginReq{
			Email:    "test@example.com",
			Password: "password123",
		}
		body, _ := json.Marshal(loginReq)

		resp, err := http.Post(ts.URL+"/api/v1/auth/register", "application/json", bytes.NewBuffer(body))
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

	t.Run("missing email", func(t *testing.T) {
		loginReq := types.LoginReq{
			Password: "password123",
		}
		body, _ := json.Marshal(loginReq)

		resp, err := http.Post(ts.URL+"/api/v1/auth/register", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", resp.StatusCode)
		}
	})

	t.Run("invalid method", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/api/v1/auth/register")
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, got %d", resp.StatusCode)
		}
	})
}

func TestLogin(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	t.Run("valid login", func(t *testing.T) {
		loginReq := types.LoginReq{
			Email:    "test@example.com",
			Password: "password123",
		}
		body, _ := json.Marshal(loginReq)

		resp, err := http.Post(ts.URL+"/api/v1/auth/login", "application/json", bytes.NewBuffer(body))
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

	t.Run("missing password", func(t *testing.T) {
		loginReq := types.LoginReq{
			Email: "test@example.com",
		}
		body, _ := json.Marshal(loginReq)

		resp, err := http.Post(ts.URL+"/api/v1/auth/login", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", resp.StatusCode)
		}
	})
}

func TestLogout(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := NewServer(logger)
	ts := httptest.NewServer(server.Router())
	defer ts.Close()

	t.Run("valid logout", func(t *testing.T) {
		resp, err := http.Post(ts.URL+"/api/v1/auth/logout", "application/json", bytes.NewBuffer([]byte("")))
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}
	})

	t.Run("invalid method", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/api/v1/auth/logout")
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, got %d", resp.StatusCode)
		}
	})
}
