package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	ThreadID  uint           `gorm:"not null" json:"thread_id"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	AuthorIP  string         `gorm:"type:inet;not null" json:"author_ip"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	IsDeleted bool           `gorm:"not null;default:false" json:"is_deleted"`
	Thread    Thread         `gorm:"foreignKey:ThreadID" json:"-"`
}

func (Post) TableName() string {
	return "posts"
}

type PostDTO struct {
	ID        uint      `json:"id"`
	ThreadID  uint      `json:"thread_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

func (p *Post) ToDTO() *PostDTO {
	return &PostDTO{
		ID:        p.ID,
		ThreadID:  p.ThreadID,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		IsDeleted: p.IsDeleted,
	}
}

func (dto *PostDTO) ToModel() *Post {
	return &Post{
		ID:        dto.ID,
		ThreadID:  dto.ThreadID,
		Content:   dto.Content,
		CreatedAt: dto.CreatedAt,
		IsDeleted: dto.IsDeleted,
	}
}
