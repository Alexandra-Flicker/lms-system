package lms

import (
	"context"
	"fmt"
	"lms_system/internal/domain/entity"
)

func (s *Service) GetAttachment(ctx context.Context, attachmentId uint, userId uint) (*entity.Attachment, error) {
	// Get attachment
	attachment, err := s.repo.Attachment().GetAttachmentById(ctx, attachmentId)
	if err != nil {
		return nil, fmt.Errorf("failed to get attachment: %w", err)
	}
	if attachment == nil {
		return nil, fmt.Errorf("attachment not found")
	}

	// Check user access to lesson
	hasAccess, err := s.CheckUserAccessToLesson(ctx, userId, attachment.LessonID)
	if err != nil {
		return nil, fmt.Errorf("failed to check access: %w", err)
	}
	if !hasAccess {
		return nil, fmt.Errorf("access denied")
	}

	return attachment, nil
}