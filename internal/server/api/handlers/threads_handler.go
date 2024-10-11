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

type ThreadHandler struct {
	service *services.ThreadService
	logger  *zap.Logger
}

func NewThreadHandler(service *services.ThreadService, logger *zap.Logger) *ThreadHandler {
	return &ThreadHandler{
		service: service,
		logger:  logger,
	}
}

func (h *ThreadHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/threads", h.GetThreads).Methods("GET")
	r.HandleFunc("/threads", h.CreateThread).Methods("POST")
	r.HandleFunc("/threads/{id}", h.GetThread).Methods("GET")
	r.HandleFunc("/threads/{id}", h.UpdateThread).Methods("PUT")
	r.HandleFunc("/threads/{id}", h.DeleteThread).Methods("DELETE")
}

func (h *ThreadHandler) GetThreads(w http.ResponseWriter, r *http.Request) {
	categoryID := r.URL.Query().Get("category_id")
	var threads []models.ThreadDTO
	var err error

	if categoryID != "" {
		id, err := strconv.Atoi(categoryID)
		if err != nil {
			h.logger.Error("Invalid category ID", zap.Error(err))
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}
		threads, err = h.service.GetThreadsByCategory(uint(id))
	} else {
		threads, err = h.service.GetAllThreads()
	}

	if err != nil {
		h.logger.Error("Failed to get threads", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(threads)
}

func (h *ThreadHandler) CreateThread(w http.ResponseWriter, r *http.Request) {
	var thread models.ThreadDTO
	if err := json.NewDecoder(r.Body).Decode(&thread); err != nil {
		h.logger.Error("Failed to decode thread", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdThread, err := h.service.CreateThread(thread)
	if err != nil {
		h.logger.Error("Failed to create thread", zap.Error(err))
		http.Error(w, "Failed to create thread", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdThread)
}

func (h *ThreadHandler) GetThread(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid thread ID", zap.Error(err))
		http.Error(w, "Invalid thread ID", http.StatusBadRequest)
		return
	}

	thread, err := h.service.GetThreadByID(uint(id))
	if err != nil {
		h.logger.Error("Failed to get thread", zap.Error(err))
		http.Error(w, "Thread not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(thread)
}

func (h *ThreadHandler) UpdateThread(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid thread ID", zap.Error(err))
		http.Error(w, "Invalid thread ID", http.StatusBadRequest)
		return
	}

	var thread models.ThreadDTO
	if err := json.NewDecoder(r.Body).Decode(&thread); err != nil {
		h.logger.Error("Failed to decode thread", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedThread, err := h.service.UpdateThread(uint(id), thread)
	if err != nil {
		h.logger.Error("Failed to update thread", zap.Error(err))
		http.Error(w, "Failed to update thread", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedThread)
}

func (h *ThreadHandler) DeleteThread(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid thread ID", zap.Error(err))
		http.Error(w, "Invalid thread ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteThread(uint(id)); err != nil {
		h.logger.Error("Failed to delete thread", zap.Error(err))
		http.Error(w, "Failed to delete thread", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
