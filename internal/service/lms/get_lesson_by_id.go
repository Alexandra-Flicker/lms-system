package lms

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lms_system/internal/domain/entity"
)

func (s *Service) GetLessonById(ctx context.Context, id uint) (entity.Lesson, error) {
	lesson, err := s.mainRepo.Lesson().GetLessonById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Lesson{}, fmt.Errorf("lesson with id %d not found", id)
		}
		return entity.Lesson{}, err
	}
	return *lesson, nil
}
