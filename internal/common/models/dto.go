package models

import "time"

// CategoryDTO represents the data transfer object for a category
type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ThreadDTO represents the data transfer object for a thread
type ThreadDTO struct {
	ID         uint      `json:"id"`
	CategoryID uint      `json:"category_id"`
	Title      string    `json:"title"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LastPostAt time.Time `json:"last_post_at"`
	PostCount  int       `json:"post_count"`
}

// PostDTO represents the data transfer object for a post
type PostDTO struct {
	ID        uint      `json:"id"`
	ThreadID  uint      `json:"thread_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	AuthorIP  string    `json:"author_ip,omitempty"` // オプショナル、管理者のみ表示
}

// CreateThreadRequest represents the request body for creating a new thread
type CreateThreadRequest struct {
	CategoryID uint   `json:"category_id"`
	Title      string `json:"title"`
}

// CreatePostRequest represents the request body for creating a new post
type CreatePostRequest struct {
	ThreadID uint   `json:"thread_id"`
	Content  string `json:"content"`
}

// PaginatedResponse represents a generic paginated response
type PaginatedResponse struct {
	TotalCount  int         `json:"total_count"`
	PageCount   int         `json:"page_count"`
	CurrentPage int         `json:"current_page"`
	PerPage     int         `json:"per_page"`
	Data        interface{} `json:"data"`
}
