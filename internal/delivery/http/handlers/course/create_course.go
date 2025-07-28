package course

import (
	"encoding/json"
	"net/http"

	"lms_system/internal/domain/entity"
)

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course entity.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateCourse(r.Context(), course)
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