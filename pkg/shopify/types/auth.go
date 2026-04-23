package types

type (
	TokenReq struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		ShopURL      string `json:"shop_url"`
	}
)
