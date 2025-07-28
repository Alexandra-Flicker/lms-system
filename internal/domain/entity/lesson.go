package entity

import "time"

type Lesson struct {
	ID            uint
	Name          string
	Description   string
	Content       string
	OrderPosition int
	ChapterID     uint
	Title         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
