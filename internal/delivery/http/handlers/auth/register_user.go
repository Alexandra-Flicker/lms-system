package auth

import (
	"encoding/json"
	"net/http"

	"lms_system/internal/domain/dto"
)

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var request dto.UserRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if request.Username == "" || request.Email == "" || request.Password == "" || len(request.Roles) == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	response, err := h.service.RegisterUser(r.Context(), &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}