package repositories

import (
	"heisei/internal/server/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) GetByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) GetByThread(threadID uint) ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Where("thread_id = ?", threadID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepository) Delete(id uint) error {
	return r.db.Delete(&models.Post{}, id).Error
}

func (r *PostRepository) SoftDelete(id uint) error {
	return r.db.Model(&models.Post{}).Where("id = ?", id).
		Update("is_deleted", true).Error
}

func (r *PostRepository) GetPostCountByThread(threadID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Post{}).Where("thread_id = ?", threadID).Count(&count).Error
	return count, err
}

func (r *PostRepository) GetLatestPostByThread(threadID uint) (*models.Post, error) {
	var post models.Post
	if err := r.db.Where("thread_id = ?", threadID).Order("created_at DESC").First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
