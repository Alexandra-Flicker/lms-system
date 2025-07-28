package chapter

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"lms_system/internal/domain/entity"
)

func (h *Handler) UpdateChapterById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid chapter ID", http.StatusBadRequest)
		return
	}

	var chapter entity.Chapter
	if err := json.NewDecoder(r.Body).Decode(&chapter); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	chapter.ID = uint(id)
	if err := h.service.UpdateChapterById(r.Context(), chapter); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}