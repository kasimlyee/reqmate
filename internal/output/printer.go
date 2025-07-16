package output

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
)

// PrintResponse prints a formatted HTTP response to stdout.
//
// The response status line is always printed, colored by the status code: green
// for 200-299, yellow for 400-499, and red for 500-599. If showHeaders is true,
// the response headers are printed in cyan after the status line. The response
// body is always printed, in white if showHeaders is true, and in the default
// color otherwise.
func PrintResponse(resp *http.Response, showHeaders bool) error {
	// Print status line
	statusCode := resp.StatusCode
	var statusColor *color.Color
	switch {
	case statusCode >= 200 && statusCode < 300:
		statusColor = color.New(color.FgGreen, color.Bold)
	case statusCode >= 400 && statusCode < 500:
		statusColor = color.New(color.FgYellow, color.Bold)
	case statusCode >= 500:
		statusColor = color.New(color.FgRed, color.Bold)
	}
	statusColor.Printf("%s %s\n", resp.Proto, resp.Status)

	// Print headers if requested
	if showHeaders {
		headerColor := color.New(color.FgCyan)
		for k, v := range resp.Header {
			headerColor.Printf("%s: %v\n", k, v)
		}
		fmt.Println()
	}

	// Print body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var bodyColor *color.Color
	if showHeaders {
		bodyColor = color.New(color.FgWhite)
	}
	bodyColor.Println(string(body))

	return nil
}

// PrintError prints the given error to stderr with a red bold color.
func PrintError(err error) {
	color.New(color.FgRed, color.Bold).Fprintf(os.Stderr, "Error: %v\n", err)
}