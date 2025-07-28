package postgres

import (
	"lms_system/internal/domain"
	"lms_system/internal/infrastructure/repository/postgres/chapter"
	"lms_system/internal/infrastructure/repository/postgres/course"
	"lms_system/internal/infrastructure/repository/postgres/lesson"
	"lms_system/internal/infrastructure/repository/postgres/user_course_access"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MainRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewMainRepository(db *gorm.DB, logger *logrus.Logger) *MainRepository {
	return &MainRepository{
		db:     db,
		logger: logger,
	}
}

func (r *MainRepository) Course() domain.CourseRepositoryInterface {
	return course.NewRepository(r.db, r.logger)
}

func (r *MainRepository) Chapter() domain.ChapterRepositoryInterface {
	return chapter.NewRepository(r.db, r.logger)
}

func (r *MainRepository) Lesson() domain.LessonRepositoryInterface {
	return lesson.NewRepository(r.db, r.logger)
}

func (r *MainRepository) UserCourseAccess() domain.UserCourseAccessInterface {
	return user_course_access.NewRepository(r.db, r.logger)
}
