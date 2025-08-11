package lms

import (
	"context"
	"fmt"
	"lms_system/internal/utils"
)

func (s *Service) CheckUserAccessToLesson(ctx context.Context, userId uint, lessonId uint) (bool, error) {
	// TODO: For now, we'll allow access for testing purposes
	// In production, you'd want to implement proper user UUID to uint mapping
	// and check the user_access_course table properly
	
	// Temporary: allow access for admin user (ID 1) and test user
	if userId == 1 || userId == utils.ConvertKeycloakIDToUint("32bfb3d7-5b2c-4502-b08a-92ae81984f57") {
		return true, nil
	}
	// Get lesson to find chapter and course
	lesson, err := s.repo.Lesson().GetLessonById(ctx, lessonId)
	if err != nil {
		return false, fmt.Errorf("failed to get lesson: %w", err)
	}
	if lesson == nil {
		return false, fmt.Errorf("lesson not found")
	}

	// Get chapter to find course
	chapter, err := s.repo.Chapter().GetChapterById(ctx, lesson.ChapterID)
	if err != nil {
		return false, fmt.Errorf("failed to get chapter: %w", err)
	}
	if chapter == nil {
		return false, fmt.Errorf("chapter not found")
	}

	// Check if user has access to the course
	access, err := s.repo.UserCourseAccess().GetByUserIdAndCourseId(ctx, userId, chapter.CourseID)
	if err != nil {
		return false, fmt.Errorf("failed to check course access: %w", err)
	}

	// User has access if they have an unlocked access record
	return access != nil && access.Unlocked, nil
}