package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (s *Service) DeleteCourseById(ctx context.Context, courseId uint) error {
	_, err := s.mainRepo.Course().GetCourseById(ctx, courseId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("course with id %d not found", courseId)
		}
		return err
	}

	err = s.mainRepo.Course().DeleteCourseById(ctx, courseId)
	if err != nil {
		return err
	}
	return nil
}
