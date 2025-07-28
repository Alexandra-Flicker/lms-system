package entity

import "time"

type UserCourseAccess struct {
	UserID   uint
	CourseID uint
	Unlocked bool
	Created  time.Time
}
