package services

import (
	"heisei/internal/common/models"
	"heisei/internal/server/repositories"

	"go.uber.org/zap"
)

type ThreadService struct {
	repo   *repositories.ThreadRepository
	logger *zap.Logger
}

func NewThreadService(repo *repositories.ThreadRepository, logger *zap.Logger) *ThreadService {
	return &ThreadService{
		repo:   repo,
		logger: logger,
	}
}

func (s *ThreadService) CreateThread(dto models.ThreadDTO) (*models.ThreadDTO, error) {
	thread := dto.ToModel()
	err := s.repo.Create(thread)
	if err != nil {
		s.logger.Error("Failed to create thread", zap.Error(err))
		return nil, err
	}
	return thread.ToDTO(), nil
}

func (s *ThreadService) GetAllThreads() ([]models.ThreadDTO, error) {
	threads, err := s.repo.GetAll()
	if err != nil {
		s.logger.Error("Failed to get all threads", zap.Error(err))
		return nil, err
	}
	threadDTOs := make([]models.ThreadDTO, len(threads))
	for i, thread := range threads {
		threadDTOs[i] = *thread.ToDTO()
	}
	return threadDTOs, nil
}

func (s *ThreadService) GetThreadByID(id uint) (*models.ThreadDTO, error) {
	thread, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("Failed to get thread by ID", zap.Error(err), zap.Uint("id", id))
		return nil, err
	}
	return thread.ToDTO(), nil
}

func (s *ThreadService) GetThreadsByCategory(categoryID uint) ([]models.ThreadDTO, error) {
	threads, err := s.repo.GetByCategory(categoryID)
	if err != nil {
		s.logger.Error("Failed to get threads by category", zap.Error(err), zap.Uint("categoryID", categoryID))
		return nil, err
	}
	threadDTOs := make([]models.ThreadDTO, len(threads))
	for i, thread := range threads {
		threadDTOs[i] = *thread.ToDTO()
	}
	return threadDTOs, nil
}

func (s *ThreadService) UpdateThread(id uint, dto models.ThreadDTO) (*models.ThreadDTO, error) {
	thread, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("Failed to get thread for update", zap.Error(err), zap.Uint("id", id))
		return nil, err
	}
	thread.Title = dto.Title
	thread.CategoryID = dto.CategoryID
	err = s.repo.Update(thread)
	if err != nil {
		s.logger.Error("Failed to update thread", zap.Error(err), zap.Uint("id", id))
		return nil, err
	}
	return thread.ToDTO(), nil
}

func (s *ThreadService) DeleteThread(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		s.logger.Error("Failed to delete thread", zap.Error(err), zap.Uint("id", id))
		return err
	}
	return nil
}

func (s *ThreadService) IncrementPostCount(threadID uint) error {
	err := s.repo.IncrementPostCount(threadID)
	if err != nil {
		s.logger.Error("Failed to increment post count", zap.Error(err), zap.Uint("threadID", threadID))
		return err
	}
	return nil
}

func (s *ThreadService) UpdateLastPostAt(threadID uint) error {
	err := s.repo.UpdateLastPostAt(threadID)
	if err != nil {
		s.logger.Error("Failed to update last post time", zap.Error(err), zap.Uint("threadID", threadID))
		return err
	}
	return nil
}
