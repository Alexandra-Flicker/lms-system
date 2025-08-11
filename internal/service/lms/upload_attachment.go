package lms

import (
	"context"
	"fmt"
	"mime/multipart"

	"lms_system/internal/domain/entity"
)

func (s *Service) UploadAttachment(ctx context.Context, lessonId uint, file multipart.File, fileHeader *multipart.FileHeader) (*entity.Attachment, error) {
	// Check if lesson exists
	lesson, err := s.repo.Lesson().GetLessonById(ctx, lessonId)
	if err != nil {
		return nil, fmt.Errorf("failed to get lesson: %w", err)
	}
	if lesson == nil {
		return nil, fmt.Errorf("lesson not found")
	}

	// Upload file to MinIO
	filePath, err := s.fileService.UploadLessonFile(ctx, lessonId, file, fileHeader)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}

	// Create attachment record
	attachment := &entity.Attachment{
		Name:     fileHeader.Filename,
		URL:      filePath,
		LessonID: lessonId,
	}

	if err := s.repo.Attachment().CreateAttachment(ctx, attachment); err != nil {
		// Try to delete uploaded file if DB save fails
		_ = s.fileService.DeleteFile(ctx, filePath)
		return nil, fmt.Errorf("failed to save attachment: %w", err)
	}

	return attachment, nil
}