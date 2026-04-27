package api

import (
	"encoding/json"
	"net/http"

	"github.com/mattrowley10/the_faywood_adapter/internal/core/domain/types"
)

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var loginReq types.LoginReq
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		s.errorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if loginReq.Email == "" || loginReq.Password == "" {
		s.errorResponse(w, http.StatusBadRequest, "email and password required")
		return
	}

	s.logger.Info("registering user", "email", loginReq.Email)

	mockCustomer := &types.Customer{
		CustomerID: "cust_123",
		Email:      loginReq.Email,
	}

	s.jsonResponse(w, http.StatusCreated, mockCustomer)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var loginReq types.LoginReq
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		s.errorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if loginReq.Email == "" || loginReq.Password == "" {
		s.errorResponse(w, http.StatusBadRequest, "email and password required")
		return
	}

	s.logger.Info("logging in user", "email", loginReq.Email)

	mockRes := &types.LoginRes{
		AccessToken: "token_abc123",
		Customer: types.Customer{
			CustomerID: "cust_456",
			Email:      loginReq.Email,
		},
	}

	s.jsonResponse(w, http.StatusOK, mockRes)
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	s.logger.Info("logging out user")

	s.jsonResponse(w, http.StatusOK, map[string]string{
		"message": "logged out successfully",
	})
}
