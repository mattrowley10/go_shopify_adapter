package types

type (
	Product struct {
		ProductID   string  `json:"product_id"`
		Slug        string  `json:"slug"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       Money   `json:"price"`
		Images      []Image `json:"images"`
	}

	Variant struct {
		VariantID string `json:"variant_id"`
		Name      string `json:"name"`
		InStock   bool   `json:"in_stock"`
		Price     Money  `json:"price"`
	}

	Image string
)
