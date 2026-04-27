package api

import (
	"encoding/json"
	"net/http"

	"github.com/mattrowley10/the_faywood_adapter/internal/core/domain/types"
)

func (s *Server) GetCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	cartID := r.PathValue("cartID")
	if cartID == "" {
		s.errorResponse(w, http.StatusBadRequest, "invalid cart ID")
		return
	}

	s.logger.Info("fetching cart", "cart_id", cartID)

	mockCart := &types.Cart{
		CartID: cartID,
		Items: []types.Item{
			{
				ProductID: "prod_1",
				VariantID: "var_1",
				Title:     "Product One",
				Quantity:  2,
				Price: types.Money{
					Amount:   9999,
					Currency: "USD",
				},
			},
		},
		Total: types.Money{
			Amount:   19998,
			Currency: "USD",
		},
	}

	s.jsonResponse(w, http.StatusOK, mockCart)
}

func (s *Server) AddToCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	cartID := r.PathValue("cartID")
	if cartID == "" {
		s.errorResponse(w, http.StatusBadRequest, "invalid cart ID")
		return
	}

	var item types.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		s.errorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if item.ProductID == "" {
		s.errorResponse(w, http.StatusBadRequest, "product_id required")
		return
	}

	s.logger.Info("adding item to cart", "cart_id", cartID, "product_id", item.ProductID)

	mockCart := &types.Cart{
		CartID: cartID,
		Items:  []types.Item{item},
		Total:  item.Price,
	}

	s.jsonResponse(w, http.StatusOK, mockCart)
}

func (s *Server) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	cartID := r.PathValue("cartID")
	itemID := r.PathValue("itemID")
	if cartID == "" || itemID == "" {
		s.errorResponse(w, http.StatusBadRequest, "invalid cart ID or item ID")
		return
	}

	s.logger.Info("removing item from cart", "cart_id", cartID, "item_id", itemID)

	mockCart := &types.Cart{
		CartID: cartID,
		Items:  []types.Item{},
		Total: types.Money{
			Amount:   0,
			Currency: "USD",
		},
	}

	s.jsonResponse(w, http.StatusOK, mockCart)
}
