package lesson

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) DeleteLessonById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteLessonById(r.Context(), uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}