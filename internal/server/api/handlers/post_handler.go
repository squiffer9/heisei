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

type PostHandler struct {
	service *services.PostService
	logger  *zap.Logger
}

func NewPostHandler(service *services.PostService, logger *zap.Logger) *PostHandler {
	return &PostHandler{
		service: service,
		logger:  logger,
	}
}

func (h *PostHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/posts", h.CreatePost).Methods("POST")
	r.HandleFunc("/threads/{threadId}/posts", h.GetPostsByThread).Methods("GET")
	r.HandleFunc("/posts/{id}", h.GetPost).Methods("GET")
	r.HandleFunc("/posts/{id}", h.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", h.DeletePost).Methods("DELETE")
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.PostDTO
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		h.logger.Error("Failed to decode post", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdPost, err := h.service.CreatePost(post)
	if err != nil {
		h.logger.Error("Failed to create post", zap.Error(err))
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)
}

func (h *PostHandler) GetPostsByThread(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	threadID, err := strconv.Atoi(vars["threadId"])
	if err != nil {
		h.logger.Error("Invalid thread ID", zap.Error(err))
		http.Error(w, "Invalid thread ID", http.StatusBadRequest)
		return
	}

	posts, err := h.service.GetPostsByThread(uint(threadID))
	if err != nil {
		h.logger.Error("Failed to get posts", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid post ID", zap.Error(err))
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := h.service.GetPostByID(uint(id))
	if err != nil {
		h.logger.Error("Failed to get post", zap.Error(err))
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid post ID", zap.Error(err))
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var post models.PostDTO
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		h.logger.Error("Failed to decode post", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedPost, err := h.service.UpdatePost(uint(id), post)
	if err != nil {
		h.logger.Error("Failed to update post", zap.Error(err))
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("Invalid post ID", zap.Error(err))
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeletePost(uint(id)); err != nil {
		h.logger.Error("Failed to delete post", zap.Error(err))
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
