package chapter

import (
	"encoding/json"
	"net/http"
	
	"lms_system/internal/domain/entity"
)

// CreateChapterRequest представляет запрос на создание главы
type CreateChapterRequest struct {
	CourseID      uint   `json:"course_id" validate:"required"`
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description"`
	OrderPosition int    `json:"order_position"`
}

// CreateChapterStandalone создает главу с указанием course_id в теле запроса
func (h *Handler) CreateChapterStandalone(w http.ResponseWriter, r *http.Request) {
	var req CreateChapterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Валидация
	if req.CourseID == 0 {
		http.Error(w, "course_id is required", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	// Создаем entity.Chapter из запроса
	chapter := entity.Chapter{
		Title:         req.Title,
		Description:   req.Description,
		OrderPosition: req.OrderPosition,
	}

	id, err := h.service.CreateChapter(r.Context(), req.CourseID, chapter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":        id,
		"course_id": req.CourseID,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}