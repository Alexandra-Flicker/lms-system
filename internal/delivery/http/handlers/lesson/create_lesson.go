package lesson

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"lms_system/internal/domain/entity"
)

func (h *Handler) CreateLesson(w http.ResponseWriter, r *http.Request) {
	chapterIdStr := chi.URLParam(r, "chapterId")
	chapterId, err := strconv.ParseUint(chapterIdStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid chapter ID", http.StatusBadRequest)
		return
	}

	var lesson entity.Lesson
	if err := json.NewDecoder(r.Body).Decode(&lesson); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateLesson(r.Context(), uint(chapterId), lesson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]uint{"id": id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}