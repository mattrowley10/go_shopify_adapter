package auth

import (
	"net/http"

	shopify_types "github.com/mattrowley10/the_faywood_adapter/pkg/shopify/types"
)

type AuthBuilder struct {
	client    *http.Client
	tokenReq  shopify_types.TokenReq
	transport *http.Transport
}

func NewAuthBuiler(client *http.Client, tokenReq shopify_types.TokenReq) *AuthBuilder {
	return &AuthBuilder{
		client:   client,
		tokenReq: tokenReq,
	}
}
