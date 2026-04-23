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
)
