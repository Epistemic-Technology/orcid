package orcid

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	MemberSandboxHost = "https://api.sandbox.orcid.org/v3.0"
	PublicSandboxHost = "https://pub.sandbox.orcid.org/v3.0"
	MemberHost        = "https://api.orcid.org/v3.0"
	PublicHost        = "https://pub.orcid.org/v3.0"
)

const (
	DefaultAPIURL     = PublicHost
	DefaultTimeout    = 30 * time.Second
	DefaultMaxRetries = 3
	DefaultRateLimit  = 10
)

type ContentType string

const (
	ContentTypeJSON ContentType = "application/json"
	ContentTypeXML  ContentType = "application/vnd.orcid+xml"
)

type Client struct {
	httpClient  *http.Client
	apiURL      string
	timeout     time.Duration
	maxRetries  int
	rateLimit   int
	userAgent   string
	contentType ContentType
	rateLimiter *time.Ticker
	bearerToken string
}

type ClientOption func(*Client)

func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		httpClient:  &http.Client{Timeout: DefaultTimeout},
		apiURL:      DefaultAPIURL,
		timeout:     DefaultTimeout,
		maxRetries:  DefaultMaxRetries,
		rateLimit:   DefaultRateLimit,
		contentType: ContentTypeJSON,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.rateLimit > 0 {
		c.rateLimiter = time.NewTicker(time.Second / time.Duration(c.rateLimit))
	}

	return c
}

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

func WithAPIURL(url string) ClientOption {
	return func(c *Client) {
		// Remove trailing slash if present
		if len(url) > 0 && url[len(url)-1] == '/' {
			c.apiURL = url[:len(url)-1]
		} else {
			c.apiURL = url
		}
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.timeout = timeout
		c.httpClient.Timeout = timeout
	}
}

func WithMaxRetries(maxRetries int) ClientOption {
	return func(c *Client) {
		c.maxRetries = maxRetries
	}
}

func WithRateLimit(requestsPerSecond int) ClientOption {
	return func(c *Client) {
		c.rateLimit = requestsPerSecond
		if requestsPerSecond > 0 {
			c.rateLimiter = time.NewTicker(time.Second / time.Duration(requestsPerSecond))
		}
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

func WithContentType(contentType ContentType) ClientOption {
	return func(c *Client) {
		c.contentType = contentType
	}
}

func WithBearerToken(token string) ClientOption {
	return func(c *Client) {
		c.bearerToken = token
	}
}

func (c *Client) doRequest(ctx context.Context, method, url string, body io.Reader) (*http.Response, error) {
	// ORCID API requires bearer token authentication for all requests
	if c.bearerToken == "" {
		return nil, fmt.Errorf("bearer token is required for ORCID API requests. Use WithBearerToken() when creating the client")
	}

	if c.rateLimiter != nil {
		select {
		case <-c.rateLimiter.C:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(attempt*attempt) * time.Second
			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}

		req, err := http.NewRequestWithContext(ctx, method, url, body)
		if err != nil {
			return nil, err
		}

		req.Header.Set("User-Agent", c.userAgent)
		req.Header.Set("Accept", string(c.contentType))
		req.Header.Set("Authorization", "Bearer "+c.bearerToken)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			continue
		}

		if resp.StatusCode == http.StatusOK {
			return resp, nil
		}

		if resp.StatusCode == http.StatusTooManyRequests ||
			resp.StatusCode == http.StatusRequestTimeout ||
			(resp.StatusCode >= 500 && resp.StatusCode < 600) {
			resp.Body.Close()
			lastErr = fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
			continue
		}

		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("HTTP %d: %s - %s", resp.StatusCode, resp.Status, string(bodyBytes))
	}

	return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}

func (c *Client) unmarshalResponse(data []byte, v interface{}) error {
	switch c.contentType {
	case ContentTypeJSON:
		return json.Unmarshal(data, v)
	case ContentTypeXML:
		return xml.Unmarshal(data, v)
	default:
		return fmt.Errorf("unsupported content type: %s", c.contentType)
	}
}

func (c *Client) buildSearchURL(params SearchParams) string {
	baseURL := c.apiURL + "/search"

	queryParams := url.Values{}
	queryParams.Set("q", params.Query)

	if params.Start > 0 {
		queryParams.Set("start", fmt.Sprintf("%d", params.Start))
	}

	if params.Rows > 0 {
		queryParams.Set("rows", fmt.Sprintf("%d", params.Rows))
	} else {
		queryParams.Set("rows", "10")
	}

	return fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
}
