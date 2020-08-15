package domain

import (
	"context"
	"time"
)

// Status ...
type Status uint8

const (
	Publish Status = iota + 1
	Draft
	Archive
)

// Course ...
type Course struct {
	ID          int64        `json:"id" `
	Title       string       `json:"title" validate:"required"`
	Description string       `json:"descrition,omitempty"`
	ImageURL    string       `json:"image_url,omitempty"`
	Duration    uint16       `json:"duration,omitempty"`
	Category    Category     `json:"categories,omitempty"`
	LessonCount int16        `json:"lesson_count,omitempty"`
	Tags        []Tag        `json:"tags,omitempty"`
	Author      User         `json:"author,omitempty"`
	Users       []User       `json:"users,omitempty"`
	Lessons     []Lesson     `json:"lessons,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	Status      Status       `json:"status,omitempty"`
	UpdatedAt   time.Time    `json:"updated_at,omitempty"`
	CreatedAt   time.Time    `json:"created_at,omitempty"`
}

// Tags ...
type Tags struct {
	Tag Tag `json:"tag"`
}

// CourseUseCase represent the course's usecases
type CourseUseCase interface {
	GetAll(ctx context.Context, start int, limit int) ([]Course, error)
	GetByID(ctx context.Context, id int64) (Course, error)
	GetByTitle(ctx context.Context, title string) (Course, error)
	// Update(ctx context.Context, course *Course) error
	CreateCourse(ctx context.Context, course *Course) error
	// Archive(ctx context.Context, course *Course) error
	// AssignToUser(ctx context.Context, course *Course, user *User)
}

// CourseRepository represent the course's repository
type CourseRepository interface {
	GetAll(ctx context.Context, start int, limit int) ([]Course, error)
	GetByID(ctx context.Context, id int64) (Course, error)
	GetByTitle(ctx context.Context, title string) (Course, error)
	UpdateCourse(ctx context.Context, course *Course) error
	CreateCourse(ctx context.Context, course *Course) error
	DeleteCourse(ctx context.Context, id int64) error
}
