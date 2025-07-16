package cmd

import (
	"fmt"
	"net/http"
	"context"
	"strings"

	"github.com/spf13/cobra"
	"github.com/kasimlyee/reqmate/internal/httpclient"
	"github.com/kasimlyee/reqmate/internal/output"
)

var (
	envName    string
	showHeaders bool
)

var runCmd = &cobra.Command{
	Use:   "run [url]",
	Short: "Execute API request or test suite",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			return executeSingleRequest(args[0])
		}
		return executeTestSuite()
	},
}

func init() {
	runCmd.Flags().StringVarP(&envName, "env", "e", "dev", "Environment to use")
	runCmd.Flags().BoolVar(&showHeaders, "headers", false, "Show response headers")
	rootCmd.AddCommand(runCmd)
}

func executeSingleRequest(url string) error {
	client := httpclient.NewClient()

	env, err := cfg.GetEnvironment(envName)
	if err != nil {
		return fmt.Errorf("failed to get active environment: %w", err)
	}

	//Use environment base URL if relative path provided
	fullURL := buildFullURL(url, env.BaseURL)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Get(context.Background(), fullURL,
		httpclient.WithHeaders(env.Headers),)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return output.PrintResponse(resp, showHeaders)
}

// buildFullURL returns the full URL given a request URL and a base URL.
//
// If the base URL is empty, or the request URL is absolute (starts with "http"),
// the request URL is returned unmodified.
//
// Otherwise, the base URL is trimmed of any trailing slash, and the request URL
// is prefixed to it after removing any leading slash from the request URL.
func buildFullURL(requestURL, baseURL string) string {
	if baseURL == "" || strings.HasPrefix(requestURL, "http") {
		return requestURL
	}
	return strings.TrimSuffix(baseURL, "/") + "/" + strings.TrimPrefix(requestURL, "/")
}

func executeTestSuite() error {
	// Will be implemented in later phases
	fmt.Println("Test suite execution coming soon...")
	return nil
}