package api

import (
	"net/http"

	"github.com/mattrowley10/the_faywood_adapter/internal/core/domain/types"
)

func (s *Server) ListProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	s.logger.Info("listing products")

	mockProducts := []types.Product{
		{
			ProductID:   "prod_1",
			Slug:        "product-one",
			Title:       "Product One",
			Description: "A mock product",
			Price: types.Money{
				Amount:   9999,
				Currency: "USD",
			},
			Images: []types.Image{"image1.jpg"},
		},
		{
			ProductID:   "prod_2",
			Slug:        "product-two",
			Title:       "Product Two",
			Description: "Another mock product",
			Price: types.Money{
				Amount:   19999,
				Currency: "USD",
			},
			Images: []types.Image{"image2.jpg"},
		},
	}

	s.jsonResponse(w, http.StatusOK, mockProducts)
}

func (s *Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	productID := r.PathValue("productID")
	if productID == "" {
		s.errorResponse(w, http.StatusBadRequest, "invalid product ID")
		return
	}

	s.logger.Info("fetching product", "product_id", productID)

	mockProduct := &types.Product{
		ProductID:   productID,
		Slug:        "product-slug",
		Title:       "Mock Product",
		Description: "A mock product for API testing",
		Price: types.Money{
			Amount:   9999,
			Currency: "USD",
		},
		Images: []types.Image{"mock_image.jpg"},
	}

	s.jsonResponse(w, http.StatusOK, mockProduct)
}
