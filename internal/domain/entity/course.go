package entity

import "time"

type CourseAggregate struct {
	Course
	Chapters []Chapter
}

type Course struct {
	ID          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
