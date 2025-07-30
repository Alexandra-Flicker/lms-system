package lms

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lms_system/internal/domain/entity"
)

func (s *Service) GetCourseById(ctx context.Context, id uint) (entity.CourseAggregate, error) {
	courseAggregate, err := s.mainRepo.Course().GetCourseById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.CourseAggregate{}, fmt.Errorf("course with id %d not found", id)
		}
		return entity.CourseAggregate{}, err
	}
	return *courseAggregate, nil
}
