package lms

import (
	"context"
	"fmt"
	"lms_system/internal/domain/entity"
)

func (s *Service) GetLessonAttachments(ctx context.Context, lessonId uint) ([]entity.Attachment, error) {
	// Check if lesson exists
	lesson, err := s.repo.Lesson().GetLessonById(ctx, lessonId)
	if err != nil {
		return nil, fmt.Errorf("failed to get lesson: %w", err)
	}
	if lesson == nil {
		return nil, fmt.Errorf("lesson not found")
	}

	// Get attachments
	attachments, err := s.repo.Attachment().GetAttachmentsByLessonId(ctx, lessonId)
	if err != nil {
		return nil, fmt.Errorf("failed to get attachments: %w", err)
	}

	return attachments, nil
}