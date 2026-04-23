package shopify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/mattrowley10/the_faywood_adapter/pkg/shopify/types"
)

type Client struct {
	shopdomain  string
	accesstoken string
	version     string
	client      *http.Client
	transport   *http.Transport
}

type ClientError struct {
	Message string `json:"message"`
}

type Shopifyer interface {
}

var (
	ErrEmptyQuery  = errors.New("query cannot be empty")
	ErrNilResult   = errors.New("result cannot be nil")
	ErrHTTPRequest = errors.New("failed to make HTTP request")
	ErrHTTPStatus  = errors.New("shopify returned an error")
)

func NewClient(shopdomain string, accesstoken string, client *http.Client) *Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxConnsPerHost:       20,
		Proxy:                 http.ProxyFromEnvironment,
	}
	return &Client{
		shopdomain:  shopdomain,
		accesstoken: accesstoken,
		version:     string(types.Version),
		client:      client,
		transport:   transport,
	}
}

func (c *Client) doRequest(
	ctx context.Context,
	method string,
	query string,
	vars map[string]any,
	result any,
) error {
	if c.shopdomain == "" || c.accesstoken == "" {
		return errors.New("shopdomain or access token not found")
	}
	url := fmt.Sprintf("https://%s.myshopify.com/admin/api/%s/graphql.json", c.shopdomain, c.version)

	req := types.ShopReq{
		Query:     query,
		Variables: vars,
	}

	bodybytes, err := json.Marshal(req)
	if err != nil {
		return errors.New("failed to marshal req query")
	}

	clientreq, err := http.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewReader(bodybytes),
	)
	if err != nil {
		return errors.New("failed to build new HTTP request with context")
	}

	clientreq.Header.Set("Content-Type", "Application/json")
	clientreq.Header.Set("x-shopify-access-token", c.accesstoken)

	resp, err := c.client.Do(clientreq)
	if err != nil {
		return ErrHTTPRequest
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("failed to read response body")
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return ErrHTTPStatus
	}

	var decoded types.ShopResp[any]
	if err := json.Unmarshal(body, &decoded); err != nil {
		return fmt.Errorf("decode graphql response: %w", err)
	}

	// GraphQL-level errors (still HTTP 200)
	if len(decoded.Errors) > 0 {
		return fmt.Errorf("graphql error: %s", decoded.Errors[0].Message)
	}

	result = decoded.Data
	return nil
}

func (c *Client) Post(ctx context.Context, req types.ShopReq, result any) error {
	if req.Query == "" {
		return ErrEmptyQuery
	}

	if result == nil {
		return ErrNilResult
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	return c.doRequest(ctx,
		http.MethodPost,
		req.Query,
		req.Variables,
		result,
	)
}
