package types

type (
	TokenReq struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		ShopURL      string `json:"shop_url"`
	}

	TokenResp struct {
		AccessToken string `json:"access_token"`
		Scope       string `json:"scope"`
		EcpiresIn   int32  `json:"expires_in"`
	}
)
