package course

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.service.GetAllCourses(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(courses); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}