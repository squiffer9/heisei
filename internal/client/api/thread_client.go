package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"heisei/internal/common/models"
	"net/http"
)

type ThreadClient struct {
	baseURL string
	client  *http.Client
}

func NewThreadClient(baseURL string, client *http.Client) *ThreadClient {
	return &ThreadClient{
		baseURL: baseURL,
		client:  client,
	}
}

func (c *ThreadClient) GetThreadsByCategory(categoryID uint) ([]models.ThreadDTO, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/api/categories/%d/threads", c.baseURL, categoryID))
	if err != nil {
		return nil, fmt.Errorf("failed to get threads: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var threads []models.ThreadDTO
	if err := json.NewDecoder(resp.Body).Decode(&threads); err != nil {
		return nil, fmt.Errorf("failed to decode threads: %w", err)
	}

	return threads, nil
}

func (c *ThreadClient) GetThreadByID(id uint) (*models.ThreadDTO, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/api/threads/%d", c.baseURL, id))
	if err != nil {
		return nil, fmt.Errorf("failed to get thread: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var thread models.ThreadDTO
	if err := json.NewDecoder(resp.Body).Decode(&thread); err != nil {
		return nil, fmt.Errorf("failed to decode thread: %w", err)
	}

	return &thread, nil
}

func (c *ThreadClient) CreateThread(thread models.ThreadDTO) (*models.ThreadDTO, error) {
	body, err := json.Marshal(thread)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal thread: %w", err)
	}

	resp, err := c.client.Post(fmt.Sprintf("%s/api/threads", c.baseURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create thread: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var createdThread models.ThreadDTO
	if err := json.NewDecoder(resp.Body).Decode(&createdThread); err != nil {
		return nil, fmt.Errorf("failed to decode created thread: %w", err)
	}

	return &createdThread, nil
}

// Additional methods like UpdateThread, DeleteThread can be added here
