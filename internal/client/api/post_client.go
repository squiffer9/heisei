package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"heisei/internal/common/models"
	"net/http"
)

type PostClient struct {
	baseURL string
	client  *http.Client
}

func NewPostClient(baseURL string, client *http.Client) *PostClient {
	return &PostClient{
		baseURL: baseURL,
		client:  client,
	}
}

func (c *PostClient) GetPostsByThread(threadID uint) ([]models.PostDTO, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/api/threads/%d/posts", c.baseURL, threadID))
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var posts []models.PostDTO
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, fmt.Errorf("failed to decode posts: %w", err)
	}

	return posts, nil
}

func (c *PostClient) GetPostByID(id uint) (*models.PostDTO, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/api/posts/%d", c.baseURL, id))
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var post models.PostDTO
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		return nil, fmt.Errorf("failed to decode post: %w", err)
	}

	return &post, nil
}

func (c *PostClient) CreatePost(post models.PostDTO) (*models.PostDTO, error) {
	body, err := json.Marshal(post)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal post: %w", err)
	}

	resp, err := c.client.Post(fmt.Sprintf("%s/api/posts", c.baseURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var createdPost models.PostDTO
	if err := json.NewDecoder(resp.Body).Decode(&createdPost); err != nil {
		return nil, fmt.Errorf("failed to decode created post: %w", err)
	}

	return &createdPost, nil
}

// Additional methods like UpdatePost, DeletePost can be added here
