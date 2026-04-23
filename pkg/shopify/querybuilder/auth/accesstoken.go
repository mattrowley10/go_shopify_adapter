package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	shopify_types "github.com/mattrowley10/the_faywood_adapter/pkg/shopify/types"
)

func (b *AuthBuilder) GetToken(ctx context.Context, tokenReq shopify_types.TokenReq) (*shopify_types.TokenResp, error) {
	url := "/admin/oauth/access_token"
	endpoint := b.baseUrl + url

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := b.client.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var tokenResp shopify_types.TokenResp
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &tokenResp, nil
}
