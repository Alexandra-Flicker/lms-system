package lms

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lms_system/internal/domain/entity"
	"time"
)

func (s *Service) UpdateChapterById(ctx context.Context, chapter entity.Chapter) error {
	existingChapter, err := s.mainRepo.Chapter().GetChapterById(ctx, chapter.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("chapter with id %d not found", chapter.ID)
		}
		return err
	}

	chapter.UpdatedAt = time.Now()
	chapter.CreatedAt = existingChapter.CreatedAt
	chapter.CourseID = existingChapter.CourseID

	err = s.mainRepo.Chapter().UpdateChapterById(ctx, &chapter)
	if err != nil {
		return err
	}
	return nil
}
