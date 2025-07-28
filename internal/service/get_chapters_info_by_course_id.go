package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lms_system/internal/domain/entity"
)

func (s *Service) GetChaptersInfoByCourseId(ctx context.Context, id uint) ([]entity.ChapterInfoAggregate, error) {
	_, err := s.mainRepo.Course().GetCourseById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("course with id %d not found", id)
		}
		return nil, err
	}

	chapters, err := s.mainRepo.Chapter().GetChaptersByCourseId(ctx, id)
	if err != nil {
		return nil, err
	}

	chaptersInfo := make([]entity.ChapterInfoAggregate, 0, len(chapters))
	for _, chapter := range chapters {
		lessons, err := s.mainRepo.Lesson().GetAllLessonsByChapterId(ctx, chapter.ID)
		if err != nil {
			return nil, err
		}

		lessonNames := make([]string, 0, len(lessons))
		for _, lesson := range lessons {
			lessonNames = append(lessonNames, lesson.Name)
		}

		chapterInfo := entity.ChapterInfoAggregate{
			Chapter:     chapter,
			LessonsName: lessonNames,
		}
		chaptersInfo = append(chaptersInfo, chapterInfo)
	}

	return chaptersInfo, nil
}
