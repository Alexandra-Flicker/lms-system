package lms

import (
	"context"
	"fmt"
)

func (s *Service) DownloadAttachment(ctx context.Context, attachmentId uint, userId uint) (string, error) {
	// Get attachment
	attachment, err := s.repo.Attachment().GetAttachmentById(ctx, attachmentId)
	if err != nil {
		return "", fmt.Errorf("failed to get attachment: %w", err)
	}
	if attachment == nil {
		return "", fmt.Errorf("attachment not found")
	}

	// Check user access to lesson
	hasAccess, err := s.CheckUserAccessToLesson(ctx, userId, attachment.LessonID)
	if err != nil {
		return "", fmt.Errorf("failed to check access: %w", err)
	}
	if !hasAccess {
		return "", fmt.Errorf("access denied")
	}

	// Get presigned URL for download
	url, err := s.fileService.GetFileURL(ctx, attachment.URL)
	if err != nil {
		return "", fmt.Errorf("failed to get download URL: %w", err)
	}

	return url, nil
}