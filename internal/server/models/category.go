package models

import (
	"gorm.io/gorm"
)

type Category struct {
	BaseModel
	Name    string   `gorm:"size:50;not null;index" json:"name" validate:"required,max=50"`
	Slug    string   `gorm:"size:50;not null;uniqueIndex" json:"slug" validate:"required,max=50,alphanum"`
	Threads []Thread `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE" json:"threads,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ToDTO converts the category model to a category DTO.
func (c *Category) ToDTO() *CategoryDTO {
	return &CategoryDTO{
		ID:   c.ID,
		Name: c.Name,
		Slug: c.Slug,
	}
}

// ToModel converts the category DTO to a category model.
func (dto *CategoryDTO) ToModel() *Category {
	return &Category{
		BaseModel: BaseModel{ID: dto.ID},
		Name:      dto.Name,
		Slug:      dto.Slug,
	}
}

// Validate validates the category.
func (c *Category) Validate() error {
	return ValidateStruct(c)
}

// IsEmpty checks if the category has no threads.
func (c *Category) IsEmpty(db *gorm.DB) (bool, error) {
	var count int64
	err := db.Model(&Thread{}).Where("category_id = ?", c.ID).Count(&count).Error
	return count == 0, err
}

// GetThreadCount returns the number of threads in the category.
func (c *Category) GetThreadCount(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Thread{}).Where("category_id = ?", c.ID).Count(&count).Error
	return count, err
}

// GetLatestThread returns the latest thread in the category.
func (c *Category) GetLatestThread(db *gorm.DB) (*Thread, error) {
	var thread Thread
	err := db.Where("category_id = ?", c.ID).Order("last_post_at DESC").First(&thread).Error
	if err != nil {
		return nil, err
	}
	return &thread, nil
}
