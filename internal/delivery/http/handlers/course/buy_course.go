package course

import (
	"encoding/json"
	"lms_system/utils"
	"net/http"
	"strconv"

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
	id, err := strconv.ParseUint(userCtx.UserID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	request.UserId = uint(id)

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
