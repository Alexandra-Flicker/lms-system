package lms

import (
	"context"
	"lms_system/internal/domain/entity"
)

func (s *Service) GetAllCourses(ctx context.Context) ([]entity.Course, error) {
	courses, err := s.repo.Course().GetAllCourses(ctx)
	if err != nil {
		return nil, err
	}
	return courses, nil
}
