package api

import (
	"net/http"

	"github.com/mattrowley10/the_faywood_adapter/internal/core/domain/types"
)

func (s *Server) ListCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	s.logger.Info("listing customers")

	mockCustomers := []types.Customer{
		{
			CustomerID: "cust_1",
			Email:      "customer1@example.com",
		},
		{
			CustomerID: "cust_2",
			Email:      "customer2@example.com",
		},
	}

	s.jsonResponse(w, http.StatusOK, mockCustomers)
}

func (s *Server) GetCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	customerID := r.PathValue("customerID")
	if customerID == "" {
		s.errorResponse(w, http.StatusBadRequest, "invalid customer ID")
		return
	}

	s.logger.Info("fetching customer", "customer_id", customerID)

	mockCustomer := &types.Customer{
		CustomerID: customerID,
		Email:      "mock@example.com",
	}

	s.jsonResponse(w, http.StatusOK, mockCustomer)
}
