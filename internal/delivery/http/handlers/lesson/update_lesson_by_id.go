package lesson

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"lms_system/internal/domain/entity"
)

func (h *Handler) UpdateLessonById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	var lesson entity.Lesson
	if err := json.NewDecoder(r.Body).Decode(&lesson); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	lesson.ID = uint(id)
	if err := h.service.UpdateLessonById(r.Context(), lesson); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}