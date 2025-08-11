package lms

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lms_system/internal/domain/entity"
	"time"
)

func (s *Service) UpdateLessonById(ctx context.Context, lesson entity.Lesson) error {
	existingLesson, err := s.repo.Lesson().GetLessonById(ctx, lesson.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("lesson with id %d not found", lesson.ID)
		}
		return err
	}

	lesson.UpdatedAt = time.Now()
	lesson.CreatedAt = existingLesson.CreatedAt
	lesson.ChapterID = existingLesson.ChapterID

	err = s.repo.Lesson().UpdateLessonById(ctx, lesson)
	if err != nil {
		return err
	}
	return nil
}
