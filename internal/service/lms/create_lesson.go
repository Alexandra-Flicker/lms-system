package lms

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lms_system/internal/domain/entity"
	"time"
)

func (s *Service) CreateLesson(ctx context.Context, chapterId uint, lesson entity.Lesson) (uint, error) {
	_, err := s.repo.Chapter().GetChapterById(ctx, chapterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("chapter with id %d not found", chapterId)
		}
		return 0, err
	}

	lesson.ChapterID = chapterId
	lesson.CreatedAt = time.Now()
	lesson.UpdatedAt = time.Now()

	id, err := s.repo.Lesson().CreateLesson(ctx, chapterId, lesson)
	if err != nil {
		return 0, err
	}
	return id, nil
}
