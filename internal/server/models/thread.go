package models

import (
	"time"

	"gorm.io/gorm"
)

type Thread struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CategoryID uint           `gorm:"not null" json:"category_id"`
	Title      string         `gorm:"size:200;not null" json:"title"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	LastPostAt time.Time      `gorm:"not null" json:"last_post_at"`
	PostCount  int            `gorm:"not null;default:0" json:"post_count"`
	Category   Category       `gorm:"foreignKey:CategoryID" json:"-"`
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

func (dto *ThreadDTO) ToModel() *Thread {
	return &Thread{
		ID:         dto.ID,
		CategoryID: dto.CategoryID,
		Title:      dto.Title,
		CreatedAt:  dto.CreatedAt,
		LastPostAt: dto.LastPostAt,
		PostCount:  dto.PostCount,
	}
}
