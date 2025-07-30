package lms

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (s *Service) DeleteChapterById(ctx context.Context, chapterId uint) error {
	_, err := s.mainRepo.Chapter().GetChapterById(ctx, chapterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("chapter with id %d not found", chapterId)
		}
		return err
	}

	err = s.mainRepo.Chapter().DeleteChapterById(ctx, chapterId)
	if err != nil {
		return err
	}
	return nil
}
