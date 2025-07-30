package lms

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lms_system/internal/domain/dto"
	"lms_system/internal/domain/entity"
	"time"
)

func (s *Service) BuyCourse(ctx context.Context, request dto.BuyCourseRequest) error {
	_, err := s.mainRepo.Course().GetCourseById(ctx, request.CourseId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("course with id %d not found", request.CourseId)
		}

		return err
	}

	existingAccess, err := s.mainRepo.UserCourseAccess().GetByUserIdAndCourseId(ctx, request.UserId, request.CourseId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existingAccess != nil {
		return fmt.Errorf("user already has access to this course")
	}

	userAccess := entity.UserCourseAccess{
		UserID:   request.UserId,
		CourseID: request.CourseId,
		Unlocked: true,
		Created:  time.Now(),
	}

	err = s.mainRepo.UserCourseAccess().CreateUserCourseAccess(ctx, userAccess)
	if err != nil {
		return err
	}
	return nil
}
