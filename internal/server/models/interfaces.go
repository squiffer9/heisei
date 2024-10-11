package models

import (
	"gorm.io/gorm"
	"time"
)

// Model is the interface that wraps the basic methods of a model.
type Model interface {
	GetID() uint
	SetID(id uint)
}

// Timestamps is the interface that wraps the methods for handling timestamps.
type Timestamps interface {
	BeforeCreate(*gorm.DB) error
	BeforeUpdate(*gorm.DB) error
}

// SoftDelete is the interface that wraps the methods for soft deleting and restoring a model.
type SoftDelete interface {
	SoftDelete(*gorm.DB) error
	Restore(*gorm.DB) error
}

// BaseModel is the base model that includes the common fields for all models.
type BaseModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time      `gorm:"index" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Validate is the interface that wraps the validation method.
type Validate interface {
	Validate() error
}

// GetID returns the ID of the model.
func (bm *BaseModel) GetID() uint {
	return bm.ID
}

// SetID sets the ID of the model.
func (bm *BaseModel) SetID(id uint) {
	bm.ID = id
}

// BeforeCreate sets the created and updated timestamps before creating a model.
func (bm *BaseModel) BeforeCreate(tx *gorm.DB) error {
	bm.CreatedAt = time.Now()
	bm.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate sets the updated timestamp before updating a model.
func (bm *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	bm.UpdatedAt = time.Now()
	return nil
}

// Validate validates the model.
func (bm *BaseModel) Validate() error {
	return nil
}
