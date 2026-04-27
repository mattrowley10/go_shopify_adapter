package api

import (
	"net/http"
	"time"

	"github.com/mattrowley10/the_faywood_adapter/internal/core/domain/types"
)

func (s *Server) ListOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	s.logger.Info("listing orders")

	mockOrders := []types.Order{
		{
			OrderID: "ord_1",
			Status:  "completed",
			Total: types.Money{
				Amount:   29999,
				Currency: "USD",
			},
			CreatedAt: time.Now().Add(-24 * time.Hour),
		},
		{
			OrderID: "ord_2",
			Status:  "pending",
			Total: types.Money{
				Amount:   19999,
				Currency: "USD",
			},
			CreatedAt: time.Now(),
		},
	}

	s.jsonResponse(w, http.StatusOK, mockOrders)
}

func (s *Server) GetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	orderID := r.PathValue("orderID")
	if orderID == "" {
		s.errorResponse(w, http.StatusBadRequest, "invalid order ID")
		return
	}

	s.logger.Info("fetching order", "order_id", orderID)

	mockOrder := &types.Order{
		OrderID: orderID,
		Status:  "completed",
		Total: types.Money{
			Amount:   29999,
			Currency: "USD",
		},
		CreatedAt: time.Now(),
	}

	s.jsonResponse(w, http.StatusOK, mockOrder)
}
