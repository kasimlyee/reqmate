package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
	Get(ctx context.Context, url string, opts ...RequestOption) (*http.Response, error)
	Post(ctx context.Context, url string, opts ...RequestOption) (*http.Response, error)
	Put(ctx context.Context, url string, opts ...RequestOption) (*http.Response, error)
	Delete(ctx context.Context, url string, opts ...RequestOption) (*http.Response, error)
}

type RequestOptions struct {
	Headers   map[string]string
	Body      []byte
	Timeout   time.Duration
	Retry     *RetryConfig
}

type RequestOption func(*RequestOptions)


type httpClient struct {
	client *http.Client
}

func NewClient() Client {
	return &httpClient{
		client: &http.Client{},
	}
}

func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

func (c *httpClient) Get(ctx context.Context, url string, options ...RequestOption) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}
	return c.doWithOptions(request, options...)
}

func (c *httpClient) Post(ctx context.Context, url string, options ...RequestOption) (*http.Response, error) {
	return c.doWithBody(ctx, "POST", url, options...)
}

func (c *httpClient) Put(ctx context.Context, url string, options ...RequestOption) (*http.Response, error) {
	return c.doWithBody(ctx, "PUT", url, options...)
}

func (c *httpClient) Delete(ctx context.Context, url string, options ...RequestOption) (*http.Response, error) {
	return c.doWithBody(ctx, "DELETE", url, options...)
}


// doWithBody sends an HTTP request with the given method, URL, and body.
//
// It marshals the given body to JSON and encodes it to the request body.
// If the body is nil, it sends an empty request body.
func (c *httpClient) doWithBody(ctx context.Context, method, url string, opts ...RequestOption) (*http.Response, error) {
	options := &RequestOptions{}
	for _, opt := range opts {
		opt(options)
	}

	var body bytes.Buffer
	if options.Body != nil {
		if err := json.NewEncoder(&body).Encode(options.Body); err != nil {
			return nil, fmt.Errorf("failed to encode body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url, &body)
	if err != nil {
		return nil, err
	}

	return c.doWithOptions(req, opts...)
}

func (c *httpClient) doWithOptions(req *http.Request, opts ...RequestOption) (*http.Response, error) {
	options := &RequestOptions{}
	for _, opt := range opts {
		opt(options)
	}

	// Apply headers
	for k, v := range options.Headers {
		req.Header.Set(k, v)
	}

	return c.client.Do(req)
}