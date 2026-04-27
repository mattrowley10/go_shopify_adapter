package api

import (
	"encoding/json"
	"net/http"

	"github.com/mattrowley10/the_faywood_adapter/internal/core/domain/types"
)

func (s *Server) CreateCheckout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var checkout types.Checkout
	if err := json.NewDecoder(r.Body).Decode(&checkout); err != nil {
		s.errorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if checkout.CartID == "" {
		s.errorResponse(w, http.StatusBadRequest, "cart_id required")
		return
	}

	s.logger.Info("creating checkout", "cart_id", checkout.CartID)

	mockCheckout := &types.Checkout{
		CartID: checkout.CartID,
	}

	s.jsonResponse(w, http.StatusCreated, mockCheckout)
}
