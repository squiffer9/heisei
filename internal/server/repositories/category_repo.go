package repositories

import (
	"context"
	"errors"
	"heisei/internal/server/models"

	"gorm.io/gorm"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
	ErrCategoryExists   = errors.New("category already exists")
)

type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new category repository
func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// Create adds a new category to the database
func (r *CategoryRepository) Create(ctx context.Context, category *models.Category) error {
	result := r.db.WithContext(ctx).Create(category)
	if result.Error != nil {
		if r.db.WithContext(ctx).Where("slug = ?", category.Slug).First(&models.Category{}).Error == nil {
			return ErrCategoryExists
		}
		return result.Error
	}
	return nil
}

// GetByID retrieves a category by its ID
func (r *CategoryRepository) GetByID(ctx context.Context, id uint) (*models.Category, error) {
	var category models.Category
	result := r.db.WithContext(ctx).First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, result.Error
	}
	return &category, nil
}

// GetAll retrieves all categories
func (r *CategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	result := r.db.WithContext(ctx).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

// Update updates an existing category
func (r *CategoryRepository) Update(ctx context.Context, category *models.Category) error {
	result := r.db.WithContext(ctx).Save(category)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrCategoryNotFound
	}
	return nil
}

// Delete removes a category by its ID
func (r *CategoryRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrCategoryNotFound
	}
	return nil
}

// GetBySlug retrieves a category by its slug
func (r *CategoryRepository) GetBySlug(ctx context.Context, slug string) (*models.Category, error) {
	var category models.Category
	result := r.db.WithContext(ctx).Where("slug = ?", slug).First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, result.Error
	}
	return &category, nil
}

// CreateWithTx adds a new category to the database with a transaction
func (r *CategoryRepository) CreateWithTx(ctx context.Context, tx *gorm.DB, category *models.Category) error {
	result := tx.WithContext(ctx).Create(category)
	if result.Error != nil {
		if tx.WithContext(ctx).Where("slug = ?", category.Slug).First(&models.Category{}).Error == nil {
			return ErrCategoryExists
		}
		return result.Error
	}
	return nil
}

// UpdateWithTx updates an existing category with a transaction
func (r *CategoryRepository) UpdateWithTx(ctx context.Context, tx *gorm.DB, category *models.Category) error {
	result := tx.WithContext(ctx).Save(category)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrCategoryNotFound
	}
	return nil
}

// DeleteWithTx removes a category by its ID with a transaction
func (r *CategoryRepository) DeleteWithTx(ctx context.Context, tx *gorm.DB, id uint) error {
	result := tx.WithContext(ctx).Delete(&models.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrCategoryNotFound
	}
	return nil
}

// GetAllPaginated retrieves all categories with pagination
func (r *CategoryRepository) GetAllPaginated(ctx context.Context, pagination *models.Pagination) ([]models.Category, error) {
	var categories []models.Category
	var total int64

	if err := r.db.Model(&models.Category{}).Count(&total).Error; err != nil {
		return nil, err
	}

	result := r.db.WithContext(ctx).Offset(pagination.Offset).Limit(pagination.Limit).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	pagination.SetTotal(total)
	return categories, nil
}
