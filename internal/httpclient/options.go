package httpclient

import "time"

// WithHeaders sets the headers to send with the request.
//
// The headers are sent as part of the request, and are used by the server to
// interpret the request.
//
// Example:
//   opts := []RequestOption{
//       WithHeaders(map[string]string{
//           "Content-Type": "application/json",
//       }),
//   }
func WithHeaders(headers map[string]string) RequestOption {
	return func(opts *RequestOptions) {
		opts.Headers = headers
	}
}

// WithBody sets the body of the request.
//
// The body is sent as part of the request, and is used by the server to
// process the request.
//
// Example:
//   opts := []RequestOption{
//       WithBody([]byte(`{"key": "value"}`)),
//   }
func WithBody(body []byte) RequestOption {
	return func(opts *RequestOptions) {
		opts.Body = body
	}
}

// WithTimeout sets the timeout for the request.
//
// The timeout is the maximum amount of time the request waits for a response
// from the server. If the timeout is exceeded, the request will be canceled and
// an error will be returned.
//
// Example:
//   opts := []RequestOption{
//       WithTimeout(10 * time.Second),
//   }
func WithTimeout(timeout time.Duration) RequestOption {
	return func(opts *RequestOptions) {
		opts.Timeout = timeout
	}
}

// WithRetryConfig sets the retry configuration for the request.
//
// The retry configuration is used to configure the retry behavior of the
// request. If the request fails, the retry configuration is used to determine
// if the request should be retried, and how many times it should be retried.
//
// Example:
//   opts := []RequestOption{
//       WithRetryConfig(&RetryConfig{
//           MaxAttempts: 3,
//           WaitTime:    100 * time.Millisecond,
//           MaxWaitTime: 2 * time.Second,
//           RetryOn:     []int{500, 502, 503, 504},
//       }),
//   }
func WithRetryConfig(config *RetryConfig) RequestOption {
	return func(opts *RequestOptions) {
		opts.Retry = config
	}
}