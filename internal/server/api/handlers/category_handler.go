package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"heisei/internal/common/models"
	"heisei/internal/server/services"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type CategoryHandler struct {
	service *services.CategoryService
	logger  *zap.Logger
}

func NewCategoryHandler(service *services.CategoryService, logger *zap.Logger) *CategoryHandler {
	return &CategoryHandler{
		service: service,
		logger:  logger,
	}
}

func (h *CategoryHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/categories", h.GetCategories).Methods("GET")
	r.HandleFunc("/categories", h.CreateCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", h.GetCategory).Methods("GET")
	r.HandleFunc("/categories/{id}", h.UpdateCategory).Methods("PUT")
	r.HandleFunc("/categories/{id}", h.DeleteCategory).Methods("DELETE")
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		h.logger.Error("Failed to get categories", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.CategoryDTO
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		h.logger.Error("Failed to decode category", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdCategory, err := h.service.CreateCategory(category)
	if err != nil {
		h.logger.Error("Failed to create category", zap.Error(err))
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCategory)
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid category ID", zap.Error(err))
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	category, err := h.service.GetCategoryByID(uint(id))
	if err != nil {
		h.logger.Error("Failed to get category", zap.Error(err))
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid category ID", zap.Error(err))
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category models.CategoryDTO
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		h.logger.Error("Failed to decode category", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedCategory, err := h.service.UpdateCategory(uint(id), category)
	if err != nil {
		h.logger.Error("Failed to update category", zap.Error(err))
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCategory)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid category ID", zap.Error(err))
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteCategory(uint(id)); err != nil {
		h.logger.Error("Failed to delete category", zap.Error(err))
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
