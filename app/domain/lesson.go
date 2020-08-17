package domain

import (
	"context"
	"time"
)

// Lesson ...
type Lesson struct {
	ID          int64     `json:"id"`
	CourseID    int64     `json:"course_id,omitempty"`
	Title       string    `json:"title,omitempty" validate:"required"`
	Description string    `json:"description,omitempty"`
	Tags        []Tag     `json:"tags,omitempty"`
	Contents    []Content `json:"contents,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

// LessonUseCase represent the Lesson's repository contract
type LessonUseCase interface {
	GetAll(ctx context.Context, start int, limit int) ([]Lesson, error)
	GetByID(ctx context.Context, id int64) (Lesson, error)
	UpdateLesson(ctx context.Context, Lesson *Lesson, id int64) error
	CreateLesson(ctx context.Context, Lesson *Lesson) error
}

// LessonRepository represent the Lesson's repository
type LessonRepository interface {
	GetAll(ctx context.Context, start int, limit int) ([]Lesson, error)
	GetByID(ctx context.Context, id int64) (Lesson, error)
	UpdateLesson(ctx context.Context, Lesson *Lesson) error
	CreateLesson(ctx context.Context, Lesson *Lesson) error
	DeleteLesson(ctx context.Context, id int64) error
}
