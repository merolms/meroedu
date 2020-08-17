package domain

import (
	"context"
	"time"
)

// Lesson ...
type Lesson struct {
	ID          int64     `json:"id"`
	CourseID    int64     `json:"course_id,omitempty"`
	Title       int64     `json:"title,omitempty" validate:"required"`
	Description string    `json:"description,omitempty"`
	Tags        []Tag     `json:"tags,omitempty"`
	Contents    []Content `json:"contents,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

// LessonRepository represent the Lesson's repository contract
type LessonRepository interface {
	GetByID(ctx context.Context, id int64) (Lesson, error)
}
