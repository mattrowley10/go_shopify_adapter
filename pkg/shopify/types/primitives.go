package types

type ShopifyAPIVersion string

const Version ShopifyAPIVersion = "2026-04"

type (
	ShopReq struct {
		Query     string `json:"query"`
		Variables map[string]any
	}

	ShopResp[T any] struct {
		Data   T `json:"data"`
		Errors []ShopErr
	}

	ShopErr struct {
		Message string `json:"message"`
	}

	// Query response types
	CustomerQueryResp struct {
		Customer *CustomerNode `json:"customer"`
	}

	CustomerNode struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}

	ProductQueryResp struct {
		Product *ProductNode `json:"product"`
	}

	ProductNode struct {
		ID       string            `json:"id"`
		Title    string            `json:"title"`
		Handle   string            `json:"handle"`
		Variants VariantConnection `json:"variants"`
		Images   ImageConnection   `json:"images"`
	}

	VariantConnection struct {
		Edges []VariantEdge `json:"edges"`
	}

	VariantEdge struct {
		Node VariantNode `json:"node"`
	}

	VariantNode struct {
		ID        string `json:"id"`
		Title     string `json:"title"`
		Price     string `json:"price"`
		Available bool   `json:"available"`
	}

	ImageConnection struct {
		Edges []ImageEdge `json:"edges"`
	}

	ImageEdge struct {
		Node ImageNode `json:"node"`
	}

	ImageNode struct {
		URL string `json:"url"`
	}

	CollectionQueryResp struct {
		Collection *CollectionNode `json:"collection"`
	}

	CollectionNode struct {
		ID     string `json:"id"`
		Title  string `json:"title"`
		Handle string `json:"handle"`
	}

	CartQueryResp struct {
		Cart *CartNode `json:"cart"`
	}

	CartNode struct {
		ID            string             `json:"id"`
		Lines         CartLineConnection `json:"lines"`
		EstimatedCost CartCostNode       `json:"estimatedCost"`
	}

	CartLineConnection struct {
		Edges []CartLineEdge `json:"edges"`
	}

	CartLineEdge struct {
		Node CartLineNode `json:"node"`
	}

	CartLineNode struct {
		ID          string               `json:"id"`
		Quantity    int                  `json:"quantity"`
		Merchandise MerchandiseInterface `json:"merchandise"`
	}

	MerchandiseInterface struct {
		ID    string    `json:"id"`
		Title string    `json:"title"`
		Price PriceNode `json:"price"`
	}

	PriceNode struct {
		Amount string `json:"amount"`
	}

	CartCostNode struct {
		TotalAmount MoneyNode `json:"totalAmount"`
	}

	MoneyNode struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currencyCode"`
	}

	CartMutateResp struct {
		CartLinesAdd    *CartLinesAddPayload    `json:"cartLinesAdd,omitempty"`
		CartLinesUpdate *CartLinesUpdatePayload `json:"cartLinesUpdate,omitempty"`
		CartLinesRemove *CartLinesRemovePayload `json:"cartLinesRemove,omitempty"`
	}

	CartLinesAddPayload struct {
		Cart *CartNode `json:"cart"`
	}

	CartLinesUpdatePayload struct {
		Cart *CartNode `json:"cart"`
	}

	CartLinesRemovePayload struct {
		Cart *CartNode `json:"cart"`
	}

	CheckoutMutateResp struct {
		CartCheckoutUrlCreate *CheckoutUrlCreate `json:"cartCheckoutUrlCreate,omitempty"`
	}

	CheckoutUrlCreate struct {
		CheckoutURL string `json:"checkoutUrl"`
	}

	OrderQueryResp struct {
		Order *OrderNode `json:"order"`
	}

	OrderNode struct {
		ID            string    `json:"id"`
		OrderNumber   string    `json:"orderNumber"`
		StatusPageURL string    `json:"statusPageUrl"`
		TotalPrice    MoneyNode `json:"totalPrice"`
		CreatedAt     string    `json:"createdAt"`
	}
)
