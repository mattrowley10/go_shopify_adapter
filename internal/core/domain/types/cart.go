package types

type (
	Cart struct {
		CartID string `json:"cart_id"`
		Items  []Item `json:"items"`
		Total  Money  `json:"total"`
	}

	Item struct {
		ProductID string `json:"product_id"`
		VariantID string `json:"variant_id"`
		Title     string `json:"title"`
		Quantity  int    `json:"quantity"`
		Price     Money  `json:"price"`
	}
)
