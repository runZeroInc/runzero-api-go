package client

import (
	"fmt"
	"os"
)

const (
	APIKey = "RUNZERO_API_KEY"
	APIURL = "RUNZERO_API_URL"

	defaultAPIURL = "https://console.runzero.com/api/v1.0"
)

// Config provides overrides the API host and key
type Config struct {
	// URL of the runZero API; otherwise, defaultAPIURL is used
	URL string
	// Key is the API key used for the session; otherwise, access is restricted to publicly accessible endpoints
	Key string
}

// NewClient returns a ready to use runZero API Client
func NewClient(c *Config) *runzeroapi.APIClient {
	// Configure the endpoint host
	apiURL := defaultAPIURL
	if envHost := os.Getenv(APIURL); envHost != "" {
		apiURL = envHost
	}
	if c.URL != "" {
		apiURL = c.URL
	}

	// Configure the authorization header
	headers := make(map[string]string)
	apiKey := os.Getenv(APIKey)
	if c.Key != "" {
		apiKey = c.Key
	}
	if apiKey != "" {
		headers["Authorization"] = fmt.Sprintf("Bearer %s", apiKey)
	}

	// Create the client
	config := runzeroapi.NewConfiguration()
	config.DefaultHeader = headers
	config.BasePath = apiURL
	return runzeroapi.NewAPIClient(config)
}
