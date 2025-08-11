package lms

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lms_system/internal/domain/entity"
	"time"
)

func (s *Service) UpdateCourseById(ctx context.Context, course entity.Course) error {
	existingCourse, err := s.repo.Course().GetCourseById(ctx, course.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("course with id %d not found", course.ID)
		}
		return err
	}

	course.UpdatedAt = time.Now()
	course.CreatedAt = existingCourse.Course.CreatedAt

	err = s.repo.Course().UpdateCourseById(ctx, course)
	if err != nil {
		return err
	}
	return nil
}
