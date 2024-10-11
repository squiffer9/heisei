package models

import (
	"gorm.io/gorm"
	"time"
)

type Thread struct {
	BaseModel
	CategoryID uint      `gorm:"not null;index" json:"category_id" validate:"required"`
	Title      string    `gorm:"size:200;not null;index" json:"title" validate:"required,max=200"`
	LastPostAt time.Time `gorm:"not null;index" json:"last_post_at"`
	PostCount  int       `gorm:"not null;default:0" json:"post_count" validate:"min=0"`
	Category   Category  `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE" json:"category,omitempty"`
	Posts      []Post    `gorm:"foreignKey:ThreadID;constraint:OnDelete:CASCADE" json:"posts,omitempty"`
}

func (Thread) TableName() string {
	return "threads"
}

type ThreadDTO struct {
	ID         uint      `json:"id"`
	CategoryID uint      `json:"category_id"`
	Title      string    `json:"title"`
	CreatedAt  time.Time `json:"created_at"`
	LastPostAt time.Time `json:"last_post_at"`
	PostCount  int       `json:"post_count"`
}

// ToDTO converts the thread model to a thread DTO.
func (t *Thread) ToDTO() *ThreadDTO {
	return &ThreadDTO{
		ID:         t.ID,
		CategoryID: t.CategoryID,
		Title:      t.Title,
		CreatedAt:  t.CreatedAt,
		LastPostAt: t.LastPostAt,
		PostCount:  t.PostCount,
	}
}

// ToModel converts the thread DTO to a thread model.
func (dto *ThreadDTO) ToModel() *Thread {
	return &Thread{
		BaseModel: BaseModel{
			ID:        dto.ID,
			CreatedAt: dto.CreatedAt,
		},
		CategoryID: dto.CategoryID,
		Title:      dto.Title,
		LastPostAt: dto.LastPostAt,
		PostCount:  dto.PostCount,
	}
}

// Validate validates the thread.
func (t *Thread) Validate() error {
	return ValidateStruct(t)
}

// IncrementPostCount increments the post count of the thread.
func (t *Thread) IncrementPostCount(db *gorm.DB) error {
	return db.Model(t).Update("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// UpdateLastPostAt updates the last post time of the thread.
func (t *Thread) UpdateLastPostAt(db *gorm.DB) error {
	return db.Model(t).Update("last_post_at", time.Now()).Error
}

// GetLatestPosts returns the latest posts of the thread.
func (t *Thread) GetLatestPosts(db *gorm.DB, n int) ([]Post, error) {
	var posts []Post
	err := db.Where("thread_id = ?", t.ID).Order("created_at DESC").Limit(n).Find(&posts).Error
	return posts, err
}
