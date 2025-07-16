package httpclient

import (
	"context"
	"net/http"
	"time"
)

type RetryConfig struct {
	MaxAttempts int
	WaitTime    time.Duration
	MaxWaitTime time.Duration
	RetryOn     []int // HTTP status codes to retry on
}

type retryingClient struct {
	client      *http.Client
	retryConfig *RetryConfig
}

// Do wraps the http.Client's Do method with retry logic.
//
// The retry logic is as follows:
//   - It will retry a request up to MaxAttempts times.
//   - It will wait for WaitTime between each retry. The wait time is
//     exponentially increased up to MaxWaitTime.
//   - It will only retry on the HTTP status codes specified in RetryOn.
func (rc *retryingClient) Do(req *http.Request) (*http.Response, error) {
	var response *http.Response
	var err error

	for attempt := 0; attempt <= rc.retryConfig.MaxAttempts; attempt++ {
		response, err = rc.client.Do(req)
		if err == nil && !shouldRetry(response.StatusCode, rc.retryConfig.RetryOn) {
			break
		}

		if attempt < rc.retryConfig.MaxAttempts {
			sleepDuration := exponentialBackoff(attempt, rc.retryConfig.WaitTime, rc.retryConfig.MaxWaitTime)
			time.Sleep(sleepDuration)
		}
	}

	return response, err
}

// shouldRetry returns true if the given HTTP status code is in the given slice of status codes to retry on.
func shouldRetry(statusCode int, retryStatusCodes []int) bool {
	for _, code := range retryStatusCodes {
		if statusCode == code {
			return true
		}
	}
	return false
}

// exponentialBackoff calculates the duration to sleep between retries given the attempt number, base duration, and maximum duration.
//
// The backoff duration is calculated as base * 2^attempt, capped at the maximum duration.
func exponentialBackoff(attempt int, base, max time.Duration) time.Duration {
	backoff := base << uint(attempt)
	if backoff > max {
		return max
	}
	return backoff
}
