package domain

import (
	"context"
	"io"
	"mime/multipart"
)

type FileServiceInterface interface {
	// UploadFile uploads a file and returns the file path/key
	UploadFile(ctx context.Context, fileType string, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
	
	// GetFile retrieves a file by its path/key
	GetFile(ctx context.Context, filePath string) (io.ReadCloser, error)
	
	// DeleteFile removes a file by its path/key
	DeleteFile(ctx context.Context, filePath string) error
	
	// GetFileURL returns a presigned URL for file access
	GetFileURL(ctx context.Context, filePath string) (string, error)
	
	// UploadLessonFile uploads a file associated with a lesson
	UploadLessonFile(ctx context.Context, lessonID uint, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
	
	// UploadCourseFile uploads a file associated with a course
	UploadCourseFile(ctx context.Context, courseID uint, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}