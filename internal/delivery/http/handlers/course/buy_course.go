package course

import (
	"encoding/json"
	"lms_system/utils"
	"net/http"

	"lms_system/internal/domain/dto"
)

func (h *Handler) BuyCourse(w http.ResponseWriter, r *http.Request) {
	userCtx := utils.GetUserFromContext(r.Context())

	var request dto.BuyCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set user ID from context
	request.UserId = userCtx.UserID

	if err := h.service.BuyCourse(r.Context(), request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "Course purchased successfully"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
