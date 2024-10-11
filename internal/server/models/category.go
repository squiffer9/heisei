package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:50;not null" json:"name" validate:"required,max=50"`
	Slug      string         `gorm:"size:50;not null;uniqueIndex" json:"slug" validate:"required,max=50,alphanum"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *Category) ToDTO() *CategoryDTO {
	return &CategoryDTO{
		ID:   c.ID,
		Name: c.Name,
		Slug: c.Slug,
	}
}

func (dto *CategoryDTO) ToModel() *Category {
	return &Category{
		ID:   dto.ID,
		Name: dto.Name,
		Slug: dto.Slug,
	}
}
