package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lms_system/internal/domain/entity"
	"time"
)

func (s *Service) CreateChapter(ctx context.Context, courseId uint, chapter entity.Chapter) (uint, error) {
	_, err := s.mainRepo.Course().GetCourseById(ctx, courseId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("course with id %d not found", courseId)
		}
		return 0, err
	}

	chapter.CourseID = courseId
	chapter.CreatedAt = time.Now()
	chapter.UpdatedAt = time.Now()

	id, err := s.mainRepo.Chapter().CreateChapter(ctx, courseId, &chapter)
	if err != nil {
		return 0, err
	}
	return id, nil
}
