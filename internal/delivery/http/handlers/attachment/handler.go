package attachment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"lms_system/internal/domain"
	"lms_system/internal/domain/common"
	"lms_system/internal/domain/entity"
	"lms_system/internal/utils"
)

type Handler struct {
	service domain.ServiceInterface
}

func NewHandler(service domain.ServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}

// UploadAttachment handles file upload for a lesson
// Only ROLE_ADMIN and ROLE_TEACHER can upload
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	// Get lesson ID from URL
	lessonIDStr := chi.URLParam(r, "lessonId")
	lessonID, err := strconv.ParseUint(lessonIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	// Parse multipart form (max 50MB)
	err = r.ParseMultipartForm(50 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get file from form
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Upload attachment
	attachment, err := h.service.UploadAttachment(r.Context(), uint(lessonID), file, fileHeader)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to upload attachment: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(attachment)
}

// DownloadAttachment handles file download
// Only users with access to the lesson can download
func (h *Handler) DownloadAttachment(w http.ResponseWriter, r *http.Request) {
	// Get attachment ID from URL
	attachmentIDStr := chi.URLParam(r, "attachmentId")
	attachmentID, err := strconv.ParseUint(attachmentIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid attachment ID", http.StatusBadRequest)
		return
	}

	// Get user context from auth middleware
	userCtx, ok := r.Context().Value(common.UserContextKey).(*entity.UserContext)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Convert Keycloak user ID to uint
	userID := utils.ConvertKeycloakIDToUint(userCtx.UserID)

	// Get download URL with access check
	downloadURL, err := h.service.DownloadAttachment(r.Context(), uint(attachmentID), userID)
	if err != nil {
		if err.Error() == "access denied" {
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}
		if err.Error() == "attachment not found" {
			http.Error(w, "Attachment not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to get download URL: %v", err), http.StatusInternalServerError)
		return
	}

	// Redirect to download URL
	http.Redirect(w, r, downloadURL, http.StatusTemporaryRedirect)
}

// GetLessonAttachments returns all attachments for a lesson
func (h *Handler) GetLessonAttachments(w http.ResponseWriter, r *http.Request) {
	// Get lesson ID from URL
	lessonIDStr := chi.URLParam(r, "lessonId")
	lessonID, err := strconv.ParseUint(lessonIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	// Get attachments
	attachments, err := h.service.GetLessonAttachments(r.Context(), uint(lessonID))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get attachments: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attachments)
}

// DeleteAttachment deletes an attachment
// Only ROLE_ADMIN can delete
func (h *Handler) DeleteAttachment(w http.ResponseWriter, r *http.Request) {
	// Get attachment ID from URL
	attachmentIDStr := chi.URLParam(r, "attachmentId")
	attachmentID, err := strconv.ParseUint(attachmentIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid attachment ID", http.StatusBadRequest)
		return
	}

	// Delete attachment
	err = h.service.DeleteAttachment(r.Context(), uint(attachmentID))
	if err != nil {
		if err.Error() == "attachment not found" {
			http.Error(w, "Attachment not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to delete attachment: %v", err), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Attachment deleted successfully",
	})
}