package repositories

import (
	"gorm.io/gorm"
)

type Repositories struct {
	DB *gorm.DB
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		DB: db,
	}
}

// Add specific repository methods here
