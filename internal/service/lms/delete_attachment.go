package lms

import (
	"context"
	"fmt"
)

func (s *Service) DeleteAttachment(ctx context.Context, attachmentId uint) error {
	// Get attachment to get file URL
	attachment, err := s.repo.Attachment().GetAttachmentById(ctx, attachmentId)
	if err != nil {
		return fmt.Errorf("failed to get attachment: %w", err)
	}
	if attachment == nil {
		return fmt.Errorf("attachment not found")
	}

	// Delete from database
	if err := s.repo.Attachment().DeleteAttachment(ctx, attachmentId); err != nil {
		return fmt.Errorf("failed to delete attachment record: %w", err)
	}

	// Delete file from MinIO
	if err := s.fileService.DeleteFile(ctx, attachment.URL); err != nil {
		// Log error but don't fail the operation
		s.logger.WithError(err).Errorf("Failed to delete file from storage: %s", attachment.URL)
	}

	return nil
}