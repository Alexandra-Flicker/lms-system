package lms

import (
	"context"
	"lms_system/internal/domain/entity"
	"time"
)

func (s *Service) CreateCourse(ctx context.Context, course entity.Course) (uint, error) {
	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()

	id, err := s.mainRepo.Course().CreateCourse(ctx, course)
	if err != nil {
		return 0, err
	}
	return id, nil
}
