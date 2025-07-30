package auth

import (
	"encoding/json"
	"net/http"

	"lms_system/internal/domain/dto"
	"lms_system/utils"
)

func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by auth middleware)
	user := utils.GetUserFromContext(r.Context())
	if user == nil || user.UserID == "" {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	var request dto.ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if request.CurrentPassword == "" || request.NewPassword == "" {
		http.Error(w, "Current password and new password are required", http.StatusBadRequest)
		return
	}

	if len(request.NewPassword) < 6 {
		http.Error(w, "New password must be at least 6 characters long", http.StatusBadRequest)
		return
	}

	response, err := h.service.ChangePassword(r.Context(), user.UserID, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}