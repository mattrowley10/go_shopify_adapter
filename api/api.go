package api

import (
	"log/slog"
	"net/http"
)

type Server struct {
	mux    *http.ServeMux
	logger *slog.Logger
}

func NewServer(logger *slog.Logger) *Server {
	mux := http.NewServeMux()

	s := &Server{
		mux:    mux,
		logger: logger,
	}

	// Health
	mux.HandleFunc("/health", s.HandleHealth)

	// Auth
	mux.HandleFunc("POST /api/v1/auth/register", s.Register)
	mux.HandleFunc("POST /api/v1/auth/login", s.Login)
	mux.HandleFunc("POST /api/v1/auth/logout", s.Logout)

	// Products
	mux.HandleFunc("GET /api/v1/products", s.ListProducts)
	mux.HandleFunc("GET /api/v1/products/{productID}", s.GetProduct)

	// Customers
	mux.HandleFunc("GET /api/v1/customers", s.ListCustomers)
	mux.HandleFunc("GET /api/v1/customers/{customerID}", s.GetCustomer)

	// Orders
	mux.HandleFunc("GET /api/v1/orders", s.ListOrders)
	mux.HandleFunc("GET /api/v1/orders/{orderID}", s.GetOrder)

	// Cart
	mux.HandleFunc("GET /api/v1/cart/{cartID}", s.GetCart)
	mux.HandleFunc("POST /api/v1/cart/{cartID}/items", s.AddToCart)
	mux.HandleFunc("DELETE /api/v1/cart/{cartID}/items/{itemID}", s.RemoveFromCart)

	// Checkout
	mux.HandleFunc("POST /api/v1/checkout", s.CreateCheckout)

	// Collections
	mux.HandleFunc("GET /api/v1/collections", s.ListCollections)
	mux.HandleFunc("GET /api/v1/collections/{collectionID}", s.GetCollection)

	return s
}

func (s *Server) Router() http.Handler {
	return s.mux
}
