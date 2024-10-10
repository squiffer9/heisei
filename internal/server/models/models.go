package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"size:50;not null"`
	Slug string `gorm:"size:50;not null;unique"`
}

type Thread struct {
	gorm.Model
	CategoryID uint      `gorm:"not null"`
	Title      string    `gorm:"size:200;not null"`
	LastPostAt time.Time `gorm:"not null"`
	PostCount  int       `gorm:"not null;default:0"`
}

type Post struct {
	gorm.Model
	ThreadID  uint   `gorm:"not null"`
	Content   string `gorm:"type:text;not null"`
	AuthorIP  string `gorm:"type:inet;not null"`
	IsDeleted bool   `gorm:"not null;default:false"`
}
