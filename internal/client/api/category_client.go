package api

import (
	"encoding/json"
	"fmt"
	"heisei/internal/common/models"
	"net/http"
)

type CategoryClient struct {
	baseURL string
	client  *http.Client
}

func NewCategoryClient(baseURL string, client *http.Client) *CategoryClient {
	return &CategoryClient{
		baseURL: baseURL,
		client:  client,
	}
}

func (c *CategoryClient) GetCategories() ([]models.CategoryDTO, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/api/categories", c.baseURL))
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var categories []models.CategoryDTO
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, fmt.Errorf("failed to decode categories: %w", err)
	}

	return categories, nil
}

func (c *CategoryClient) GetCategoryByID(id uint) (*models.CategoryDTO, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/api/categories/%d", c.baseURL, id))
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var category models.CategoryDTO
	if err := json.NewDecoder(resp.Body).Decode(&category); err != nil {
		return nil, fmt.Errorf("failed to decode category: %w", err)
	}

	return &category, nil
}

// Additional methods like CreateCategory, UpdateCategory, DeleteCategory can be added here
