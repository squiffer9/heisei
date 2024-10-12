package repositories

import (
	"context"
	"errors"
	"heisei/internal/server/models"

	"gorm.io/gorm"
)

var (
	ErrPostNotFound = errors.New("post not found")
	ErrPostExists   = errors.New("post already exists")
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

// Create adds a new post to the database
func (r *PostRepository) Create(ctx context.Context, post *models.Post) error {
	result := r.db.WithContext(ctx).Create(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetByID retrieves a post by its ID
func (r *PostRepository) GetByID(ctx context.Context, id uint) (*models.Post, error) {
	var post models.Post
	result := r.db.WithContext(ctx).First(&post, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrPostNotFound
		}
		return nil, result.Error
	}
	return &post, nil
}

// GetByThread retrieves posts by thread ID
func (r *PostRepository) GetByThread(ctx context.Context, threadID uint) ([]models.Post, error) {
	var posts []models.Post
	result := r.db.WithContext(ctx).Where("thread_id = ?", threadID).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

// Update updates an existing post
func (r *PostRepository) Update(ctx context.Context, post *models.Post) error {
	result := r.db.WithContext(ctx).Save(post)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPostNotFound
	}
	return nil
}

// Delete removes a post by its ID
func (r *PostRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPostNotFound
	}
	return nil
}

// SoftDelete marks a post as deleted without actually deleting it
func (r *PostRepository) SoftDelete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Model(&models.Post{}).Where("id = ?", id).
		Update("is_deleted", true)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPostNotFound
	}
	return nil
}

// GetPostCountByThread returns the number of posts in a thread
func (r *PostRepository) GetPostCountByThread(ctx context.Context, threadID uint) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&models.Post{}).Where("thread_id = ?", threadID).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// GetLatestPostByThread returns the latest post in a thread
func (r *PostRepository) GetLatestPostByThread(ctx context.Context, threadID uint) (*models.Post, error) {
	var post models.Post
	result := r.db.WithContext(ctx).Where("thread_id = ?", threadID).Order("created_at DESC").First(&post)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrPostNotFound
		}
		return nil, result.Error
	}
	return &post, nil
}

// CreateWithTx adds a new post to the database within a transaction
func (r *PostRepository) CreateWithTx(ctx context.Context, tx *gorm.DB, post *models.Post) error {
	result := tx.WithContext(ctx).Create(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateWithTx updates an existing post within a transaction
func (r *PostRepository) UpdateWithTx(ctx context.Context, tx *gorm.DB, post *models.Post) error {
	result := tx.WithContext(ctx).Save(post)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPostNotFound
	}
	return nil
}

// DeleteWithTx removes a post by its ID within a transaction
func (r *PostRepository) DeleteWithTx(ctx context.Context, tx *gorm.DB, id uint) error {
	result := tx.WithContext(ctx).Delete(&models.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPostNotFound
	}
	return nil
}

// SoftDeleteWithTx marks a post as deleted without actually deleting it within a transaction
func (r *PostRepository) SoftDeleteWithTx(ctx context.Context, tx *gorm.DB, id uint) error {
	result := tx.WithContext(ctx).Model(&models.Post{}).Where("id = ?", id).
		Update("is_deleted", true)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPostNotFound
	}
	return nil
}

// GetByThreadPaginated retrieves posts by thread ID with pagination
func (r *PostRepository) GetByThreadPaginated(ctx context.Context, threadID uint, pagination *models.Pagination) ([]models.Post, error) {
	var posts []models.Post
	var total int64

	if err := r.db.Model(&models.Post{}).Where("thread_id = ?", threadID).Count(&total).Error; err != nil {
		return nil, err
	}

	result := r.db.WithContext(ctx).Where("thread_id = ?", threadID).Offset(pagination.Offset).Limit(pagination.Limit).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	pagination.SetTotal(total)
	return posts, nil
}
