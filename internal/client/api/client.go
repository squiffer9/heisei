package api

import (
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// TODO: Add methods for API calls here
// Example:
// func (c *Client) GetCategories() ([]models.CategoryDTO, error) {
//     // Implementation
// }
