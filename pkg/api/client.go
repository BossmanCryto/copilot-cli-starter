package api

import (
	"fmt"
	"net/http"
)

// Client handles API communication
type Client struct {
	BaseURL string
	Token   string
	HTTP    *http.Client
}

// NewClient creates a new API client
func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL: baseURL,
		Token:   token,
		HTTP:    &http.Client{},
	}
}

// Get performs a GET request
func (c *Client) Get(endpoint string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return nil, nil
}
