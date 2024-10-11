package repositories

import (
	"heisei/internal/server/models"

	"gorm.io/gorm"
)

type ThreadRepository struct {
	db *gorm.DB
}

func NewThreadRepository(db *gorm.DB) *ThreadRepository {
	return &ThreadRepository{db: db}
}

func (r *ThreadRepository) Create(thread *models.Thread) error {
	return r.db.Create(thread).Error
}

func (r *ThreadRepository) GetByID(id uint) (*models.Thread, error) {
	var thread models.Thread
	if err := r.db.First(&thread, id).Error; err != nil {
		return nil, err
	}
	return &thread, nil
}

func (r *ThreadRepository) GetAll() ([]models.Thread, error) {
	var threads []models.Thread
	if err := r.db.Find(&threads).Error; err != nil {
		return nil, err
	}
	return threads, nil
}

func (r *ThreadRepository) GetByCategory(categoryID uint) ([]models.Thread, error) {
	var threads []models.Thread
	if err := r.db.Where("category_id = ?", categoryID).Find(&threads).Error; err != nil {
		return nil, err
	}
	return threads, nil
}

func (r *ThreadRepository) Update(thread *models.Thread) error {
	return r.db.Save(thread).Error
}

func (r *ThreadRepository) Delete(id uint) error {
	return r.db.Delete(&models.Thread{}, id).Error
}

func (r *ThreadRepository) IncrementPostCount(threadID uint) error {
	return r.db.Model(&models.Thread{}).Where("id = ?", threadID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

func (r *ThreadRepository) UpdateLastPostAt(threadID uint) error {
	return r.db.Model(&models.Thread{}).Where("id = ?", threadID).
		UpdateColumn("last_post_at", gorm.Expr("CURRENT_TIMESTAMP")).Error
}
