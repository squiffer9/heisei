package repositories

import (
	"context"
	"errors"
	"heisei/internal/server/models"

	"gorm.io/gorm"
)

var (
	ErrThreadNotFound = errors.New("thread not found")
	ErrThreadExists   = errors.New("thread already exists")
)

type ThreadRepository struct {
	db *gorm.DB
}

func NewThreadRepository(db *gorm.DB) *ThreadRepository {
	return &ThreadRepository{db: db}
}

// Create adds a new thread to the database
func (r *ThreadRepository) Create(ctx context.Context, thread *models.Thread) error {
	result := r.db.WithContext(ctx).Create(thread)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetByID retrieves a thread by its ID
func (r *ThreadRepository) GetByID(ctx context.Context, id uint) (*models.Thread, error) {
	var thread models.Thread
	result := r.db.WithContext(ctx).First(&thread, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrThreadNotFound
		}
		return nil, result.Error
	}
	return &thread, nil
}

// GetAll retrieves all threads
func (r *ThreadRepository) GetAll(ctx context.Context) ([]models.Thread, error) {
	var threads []models.Thread
	result := r.db.WithContext(ctx).Find(&threads)
	if result.Error != nil {
		return nil, result.Error
	}
	return threads, nil
}

// GetByCategory retrieves threads by category ID
func (r *ThreadRepository) GetByCategory(ctx context.Context, categoryID uint) ([]models.Thread, error) {
	var threads []models.Thread
	result := r.db.WithContext(ctx).Where("category_id = ?", categoryID).Find(&threads)
	if result.Error != nil {
		return nil, result.Error
	}
	return threads, nil
}

// Update updates an existing thread
func (r *ThreadRepository) Update(ctx context.Context, thread *models.Thread) error {
	result := r.db.WithContext(ctx).Save(thread)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrThreadNotFound
	}
	return nil
}

// Delete removes a thread by its ID
func (r *ThreadRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Thread{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrThreadNotFound
	}
	return nil
}

// IncrementPostCount increments the post count of a thread
func (r *ThreadRepository) IncrementPostCount(ctx context.Context, threadID uint) error {
	result := r.db.WithContext(ctx).Model(&models.Thread{}).Where("id = ?", threadID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrThreadNotFound
	}
	return nil
}

// UpdateLastPostAt updates the last post time of a thread
func (r *ThreadRepository) UpdateLastPostAt(ctx context.Context, threadID uint) error {
	result := r.db.WithContext(ctx).Model(&models.Thread{}).Where("id = ?", threadID).
		UpdateColumn("last_post_at", gorm.Expr("CURRENT_TIMESTAMP"))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrThreadNotFound
	}
	return nil
}

// CreateWithTx creates a new thread within a transaction
func (r *ThreadRepository) CreateWithTx(ctx context.Context, tx *gorm.DB, thread *models.Thread) error {
	result := tx.WithContext(ctx).Create(thread)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateWithTx updates an existing thread within a transaction
func (r *ThreadRepository) UpdateWithTx(ctx context.Context, tx *gorm.DB, thread *models.Thread) error {
	result := tx.WithContext(ctx).Save(thread)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrThreadNotFound
	}
	return nil
}

// DeleteWithTx removes a thread by its ID within a transaction
func (r *ThreadRepository) DeleteWithTx(ctx context.Context, tx *gorm.DB, id uint) error {
	result := tx.WithContext(ctx).Delete(&models.Thread{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrThreadNotFound
	}
	return nil
}

// IncrementPostCountWithTx increments the post count of a thread within a transaction
func (r *ThreadRepository) IncrementPostCountWithTx(ctx context.Context, tx *gorm.DB, threadID uint) error {
	result := tx.WithContext(ctx).Model(&models.Thread{}).Where("id = ?", threadID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrThreadNotFound
	}
	return nil
}

// UpdateLastPostAtWithTx updates the last post time of a thread within a transaction
func (r *ThreadRepository) UpdateLastPostAtWithTx(ctx context.Context, tx *gorm.DB, threadID uint) error {
	result := tx.WithContext(ctx).Model(&models.Thread{}).Where("id = ?", threadID).
		UpdateColumn("last_post_at", gorm.Expr("CURRENT_TIMESTAMP"))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrThreadNotFound
	}
	return nil
}
