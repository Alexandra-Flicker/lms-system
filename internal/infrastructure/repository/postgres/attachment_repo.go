package postgres

import (
	"context"
	"lms_system/internal/domain"
	"lms_system/internal/domain/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AttachmentRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewAttachmentRepository(db *gorm.DB, logger *logrus.Logger) domain.AttachmentRepositoryInterface {
	return &AttachmentRepository{
		db:     db,
		logger: logger,
	}
}

func (r *AttachmentRepository) CreateAttachment(ctx context.Context, attachment *entity.Attachment) error {
	if err := r.db.WithContext(ctx).Create(attachment).Error; err != nil {
		r.logger.WithError(err).Error("Failed to create attachment")
		return err
	}
	return nil
}

func (r *AttachmentRepository) GetAttachmentById(ctx context.Context, id uint) (*entity.Attachment, error) {
	var attachment entity.Attachment
	if err := r.db.WithContext(ctx).First(&attachment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.logger.WithError(err).Error("Failed to get attachment by id")
		return nil, err
	}
	return &attachment, nil
}

func (r *AttachmentRepository) GetAttachmentsByLessonId(ctx context.Context, lessonId uint) ([]entity.Attachment, error) {
	var attachments []entity.Attachment
	if err := r.db.WithContext(ctx).Where("lesson_id = ?", lessonId).Order("created_at DESC").Find(&attachments).Error; err != nil {
		r.logger.WithError(err).Error("Failed to get attachments by lesson id")
		return nil, err
	}
	return attachments, nil
}

func (r *AttachmentRepository) DeleteAttachment(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Attachment{}, id).Error; err != nil {
		r.logger.WithError(err).Error("Failed to delete attachment")
		return err
	}
	return nil
}