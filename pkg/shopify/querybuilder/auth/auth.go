package auth

import (
	"net/http"

	"github.com/mattrowley10/the_faywood_adapter/pkg/shopify"
	shopify_types "github.com/mattrowley10/the_faywood_adapter/pkg/shopify/types"
)

type AuthBuilder struct {
	client    *shopify.Client
	baseUrl   string
	tokenReq  shopify_types.TokenReq
	transport *http.Transport
}

func NewAuthBuiler(client *shopify.Client, tokenReq shopify_types.TokenReq) *AuthBuilder {
	return &AuthBuilder{
		client:   client,
		tokenReq: tokenReq,
	}
}
