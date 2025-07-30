package lms

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (s *Service) DeleteLessonById(ctx context.Context, lessonId uint) error {
	_, err := s.mainRepo.Lesson().GetLessonById(ctx, lessonId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("lesson with id %d not found", lessonId)
		}
		return err
	}

	err = s.mainRepo.Lesson().DeleteLessonById(ctx, lessonId)
	if err != nil {
		return err
	}
	return nil
}
