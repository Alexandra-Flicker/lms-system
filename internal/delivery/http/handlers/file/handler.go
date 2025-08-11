package file

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"lms_system/internal/domain"
)

type Handler struct {
	fileService domain.FileServiceInterface
}

func NewHandler(fileService domain.FileServiceInterface) *Handler {
	return &Handler{
		fileService: fileService,
	}
}

// UploadFile handles generic file upload
func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form (max 10MB)
	err := r.ParseMultipartForm(10 << 20)
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

	// Get file type from form (optional, defaults to "general")
	fileType := r.FormValue("type")
	if fileType == "" {
		fileType = "general"
	}

	// Upload file
	filePath, err := h.fileService.UploadFile(r.Context(), fileType, file, fileHeader)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to upload file: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response
	response := map[string]string{
		"file_path": filePath,
		"message":   "File uploaded successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// UploadLessonFile handles file upload for a specific lesson
func (h *Handler) UploadLessonFile(w http.ResponseWriter, r *http.Request) {
	// Get lesson ID from URL
	lessonIDStr := chi.URLParam(r, "lessonId")
	lessonID, err := strconv.ParseUint(lessonIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	// Parse multipart form (max 10MB)
	err = r.ParseMultipartForm(10 << 20)
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

	// Upload file
	filePath, err := h.fileService.UploadLessonFile(r.Context(), uint(lessonID), file, fileHeader)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to upload file: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response
	response := map[string]interface{}{
		"file_path": filePath,
		"lesson_id": lessonID,
		"message":   "File uploaded successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// UploadCourseFile handles file upload for a specific course
func (h *Handler) UploadCourseFile(w http.ResponseWriter, r *http.Request) {
	// Get course ID from URL
	courseIDStr := chi.URLParam(r, "courseId")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid course ID", http.StatusBadRequest)
		return
	}

	// Parse multipart form (max 10MB)
	err = r.ParseMultipartForm(10 << 20)
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

	// Upload file
	filePath, err := h.fileService.UploadCourseFile(r.Context(), uint(courseID), file, fileHeader)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to upload file: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response
	response := map[string]interface{}{
		"file_path": filePath,
		"course_id": courseID,
		"message":   "File uploaded successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// DownloadFile handles file download
func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	// Get file path from query parameter
	filePath := r.URL.Query().Get("path")
	if filePath == "" {
		http.Error(w, "File path is required", http.StatusBadRequest)
		return
	}

	// Get file from storage
	fileReader, err := h.fileService.GetFile(r.Context(), filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get file: %v", err), http.StatusNotFound)
		return
	}
	defer fileReader.Close()

	// Set appropriate headers
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filePath))
	w.Header().Set("Content-Type", "application/octet-stream")

	// Copy file to response
	_, err = io.Copy(w, fileReader)
	if err != nil {
		http.Error(w, "Failed to send file", http.StatusInternalServerError)
		return
	}
}

// GetFileURL returns a presigned URL for file access
func (h *Handler) GetFileURL(w http.ResponseWriter, r *http.Request) {
	// Get file path from query parameter
	filePath := r.URL.Query().Get("path")
	if filePath == "" {
		http.Error(w, "File path is required", http.StatusBadRequest)
		return
	}

	// Get presigned URL
	url, err := h.fileService.GetFileURL(r.Context(), filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get file URL: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response
	response := map[string]string{
		"url":       url,
		"file_path": filePath,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteFile handles file deletion
func (h *Handler) DeleteFile(w http.ResponseWriter, r *http.Request) {
	// Get file path from request body
	var req struct {
		FilePath string `json:"file_path"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.FilePath == "" {
		http.Error(w, "File path is required", http.StatusBadRequest)
		return
	}

	// Delete file
	err := h.fileService.DeleteFile(r.Context(), req.FilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete file: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response
	response := map[string]string{
		"message": "File deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}