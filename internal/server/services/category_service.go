package services

import (
	"heisei/internal/common/models"
	"heisei/internal/server/repositories"

	"go.uber.org/zap"
)

type CategoryService struct {
	repo   *repositories.CategoryRepository
	logger *zap.Logger
}

func NewCategoryService(repo *repositories.CategoryRepository, logger *zap.Logger) *CategoryService {
	return &CategoryService{
		repo:   repo,
		logger: logger,
	}
}

func (s *CategoryService) CreateCategory(dto models.CategoryDTO) (*models.CategoryDTO, error) {
	category := dto.ToModel()
	err := s.repo.Create(category)
	if err != nil {
		s.logger.Error("Failed to create category", zap.Error(err))
		return nil, err
	}
	return category.ToDTO(), nil
}

func (s *CategoryService) GetAllCategories() ([]models.CategoryDTO, error) {
	categories, err := s.repo.GetAll()
	if err != nil {
		s.logger.Error("Failed to get all categories", zap.Error(err))
		return nil, err
	}
	categoryDTOs := make([]models.CategoryDTO, len(categories))
	for i, category := range categories {
		categoryDTOs[i] = *category.ToDTO()
	}
	return categoryDTOs, nil
}

func (s *CategoryService) GetCategoryByID(id uint) (*models.CategoryDTO, error) {
	category, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("Failed to get category by ID", zap.Error(err), zap.Uint("id", id))
		return nil, err
	}
	return category.ToDTO(), nil
}

func (s *CategoryService) UpdateCategory(id uint, dto models.CategoryDTO) (*models.CategoryDTO, error) {
	category, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("Failed to get category for update", zap.Error(err), zap.Uint("id", id))
		return nil, err
	}
	category.Name = dto.Name
	category.Slug = dto.Slug
	err = s.repo.Update(category)
	if err != nil {
		s.logger.Error("Failed to update category", zap.Error(err), zap.Uint("id", id))
		return nil, err
	}
	return category.ToDTO(), nil
}

func (s *CategoryService) DeleteCategory(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		s.logger.Error("Failed to delete category", zap.Error(err), zap.Uint("id", id))
		return err
	}
	return nil
}
