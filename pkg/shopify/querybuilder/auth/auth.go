package auth

import (
	"context"
	"net/http"

	shopify_types "github.com/mattrowley10/the_faywood_adapter/pkg/shopify/types"
)

type AuthBuilder struct {
	client    http.RoundTripper
	baseUrl   string
	tokenReq  shopify_types.TokenReq
	transport *http.Transport
}

type Auther interface {
	GetToken(ctx context.Context, tokenReq shopify_types.TokenReq) (*shopify_types.TokenResp, error)
}

func NewAuthBuiler(client http.RoundTripper, tokenReq shopify_types.TokenReq) *AuthBuilder {
	return &AuthBuilder{
		client:   client,
		tokenReq: tokenReq,
	}
}
