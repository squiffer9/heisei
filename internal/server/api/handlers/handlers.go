package handlers

import (
	"net/http"

	"go.uber.org/zap"
	"heisei/internal/server/services"
)

type Handlers struct {
	Services *services.Services
	Logger   *zap.Logger
}

func NewHandlers(services *services.Services, logger *zap.Logger) *Handlers {
	return &Handlers{
		Services: services,
		Logger:   logger,
	}
}

// SetupRoutes sets up the routes for the API
func (h *Handlers) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Add your routes here
	// Example:
	// mux.HandleFunc("/api/categories", h.GetCategories)

	return mux
}
