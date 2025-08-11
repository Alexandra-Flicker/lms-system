package lesson

import (
	"encoding/json"
	"net/http"
	
	"lms_system/internal/domain/entity"
)

// CreateLessonRequest представляет запрос на создание урока
type CreateLessonRequest struct {
	ChapterID     uint   `json:"chapter_id" validate:"required"`
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description"`
	Content       string `json:"content"`
	OrderPosition int    `json:"order_position"`
}

// CreateLessonStandalone создает урок с указанием chapter_id в теле запроса
func (h *Handler) CreateLessonStandalone(w http.ResponseWriter, r *http.Request) {
	var req CreateLessonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Валидация
	if req.ChapterID == 0 {
		http.Error(w, "chapter_id is required", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	// Создаем entity.Lesson из запроса
	lesson := entity.Lesson{
		Title:         req.Title,
		Description:   req.Description,
		Content:       req.Content,
		OrderPosition: req.OrderPosition,
	}

	id, err := h.service.CreateLesson(r.Context(), req.ChapterID, lesson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":         id,
		"chapter_id": req.ChapterID,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}