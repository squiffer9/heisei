package services

import (
	"heisei/internal/common/models"
	"heisei/internal/server/repositories"

	"go.uber.org/zap"
)

type PostService struct {
	repo          *repositories.PostRepository
	threadRepo    *repositories.ThreadRepository
	threadService *ThreadService
	logger        *zap.Logger
}

func NewPostService(repo *repositories.PostRepository, threadRepo *repositories.ThreadRepository, threadService *ThreadService, logger *zap.Logger) *PostService {
	return &PostService{
		repo:          repo,
		threadRepo:    threadRepo,
		threadService: threadService,
		logger:        logger,
	}
}

func (s *PostService) CreatePost(dto models.PostDTO) (*models.PostDTO, error) {
	post := dto.ToModel()
	err := s.repo.Create(post)
	if err != nil {
		s.logger.Error("Failed to create post", zap.Error(err))
		return nil, err
	}

	// Update thread's post count and last post time
	err = s.threadService.IncrementPostCount(post.ThreadID)
	if err != nil {
		s.logger.Error("Failed to increment thread post count", zap.Error(err), zap.Uint("threadID", post.ThreadID))
	}
	err = s.threadService.UpdateLastPostAt(post.ThreadID)
	if err != nil {
		s.logger.Error("Failed to update thread last post time", zap.Error(err), zap.Uint("threadID", post.ThreadID))
	}

	return post.ToDTO(), nil
}

func (s *PostService) GetPostByID(id uint) (*models.PostDTO, error) {
	post, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("Failed to get post by ID", zap.Error(err), zap.Uint("id", id))
		return nil, err
	}
	return post.ToDTO(), nil
}

func (s *PostService) GetPostsByThread(threadID uint) ([]models.PostDTO, error) {
	posts, err := s.repo.GetByThread(threadID)
	if err != nil {
		s.logger.Error("Failed to get posts by thread", zap.Error(err), zap.Uint("threadID", threadID))
		return nil, err
	}
	postDTOs := make([]models.PostDTO, len(posts))
	for i, post := range posts {
		postDTOs[i] = *post.ToDTO()
	}
	return postDTOs, nil
}

func (s *PostService) UpdatePost(id uint, dto models.PostDTO) (*models.PostDTO, error) {
	post, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("Failed to get post for update", zap.Error(err), zap.Uint("id", id))
		return nil, err
	}
	post.Content = dto.Content
	err = s.repo.Update(post)
	if err != nil {
		s.logger.Error("Failed to update post", zap.Error(err), zap.Uint("id", id))
		return nil, err
	}
	return post.ToDTO(), nil
}

func (s *PostService) DeletePost(id uint) error {
	err := s.repo.SoftDelete(id)
	if err != nil {
		s.logger.Error("Failed to soft delete post", zap.Error(err), zap.Uint("id", id))
		return err
	}
	return nil
}

func (s *PostService) GetPostCountByThread(threadID uint) (int64, error) {
	count, err := s.repo.GetPostCountByThread(threadID)
	if err != nil {
		s.logger.Error("Failed to get post count by thread", zap.Error(err), zap.Uint("threadID", threadID))
		return 0, err
	}
	return count, nil
}

func (s *PostService) GetLatestPostByThread(threadID uint) (*models.PostDTO, error) {
	post, err := s.repo.GetLatestPostByThread(threadID)
	if err != nil {
		s.logger.Error("Failed to get latest post by thread", zap.Error(err), zap.Uint("threadID", threadID))
		return nil, err
	}
	return post.ToDTO(), nil
}
