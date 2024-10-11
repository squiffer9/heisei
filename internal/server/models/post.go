package models

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type Post struct {
	BaseModel
	ThreadID  uint   `gorm:"not null;index" json:"thread_id" validate:"required"`
	Content   string `gorm:"type:text;not null" json:"content" validate:"required,min=1,max=10000"`
	AuthorIP  string `gorm:"type:inet;not null" json:"author_ip" validate:"required,ip"`
	IsDeleted bool   `gorm:"not null;default:false;index" json:"is_deleted"`
	Thread    Thread `gorm:"foreignKey:ThreadID;constraint:OnDelete:CASCADE" json:"thread,omitempty"`
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
		BaseModel: BaseModel{
			ID:        dto.ID,
			CreatedAt: dto.CreatedAt,
		},
		ThreadID:  dto.ThreadID,
		Content:   dto.Content,
		IsDeleted: dto.IsDeleted,
	}
}

// Validate validates the post.
func (p *Post) Validate() error {
	return ValidateStruct(p)
}

// SoftDelete marks the post as deleted without actually deleting it.
func (p *Post) SoftDelete(db *gorm.DB) error {
	return db.Model(p).Update("is_deleted", true).Error
}

// Restore restores the post.
func (p *Post) Restore(db *gorm.DB) error {
	return db.Model(p).Update("is_deleted", false).Error
}

// GetAuthorIPMasked returns the author IP with the last octet masked.
func (p *Post) GetAuthorIPMasked() string {
	parts := strings.Split(p.AuthorIP, ".")
	if len(parts) == 4 {
		parts[3] = "xxx"
		return strings.Join(parts, ".")
	}
	return p.AuthorIP
}

// TODO: Implement the IsEditable()
