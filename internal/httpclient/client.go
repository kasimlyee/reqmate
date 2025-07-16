package httpclient

import (
	"context"
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