package entity

import "time"

type Attachment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	URL       string    `json:"url" gorm:"type:varchar(255);not null"`
	LessonID  uint      `json:"lesson_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Relations
	Lesson Lesson `json:"-" gorm:"foreignKey:LessonID"`
}