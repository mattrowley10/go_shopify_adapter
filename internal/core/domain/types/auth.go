package types

type (
	LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginRes struct {
		AccessToken string   `json:"access_token"`
		Customer    Customer `json:"customer"`
	}
)
