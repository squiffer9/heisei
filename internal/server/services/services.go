package services

import (
	"heisei/internal/server/repositories"

	"go.uber.org/zap"
)

type Services struct {
	Repos  *repositories.Repositories
	Logger *zap.Logger
}

func NewServices(repos *repositories.Repositories, logger *zap.Logger) *Services {
	return &Services{
		Repos:  repos,
		Logger: logger,
	}
}

// Add specific service methods here
