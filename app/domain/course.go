package domain

import (
	"context"
	"time"
)

// const (
// 	_            = iota
// 	Publish uint8
// 	Draft uint8
// 	Archive uint8
// )

// Course ...
type Course struct {
	ID          int64        `json:"id" `
	Title       string       `json:"title" validate:"required"`
	Description string       `json:"descrition,omitempty"`
	ImageURL    string       `json:"image_url,omitempty"`
	Duration    uint16       `json:"duration,omitempty"`
	CategoryID  int64        `json:"category_id,omitempty"`
	LessonCount int16        `json:"lesson_count,omitempty"`
	Tags        []Tag        `json:"tags,omitempty" gorm:"many2many:course_tag;"`
	Author      User         `json:"authors,omitempty" gorm:"many2many:course_author;"`
	Users       []User       `json:"users,omitempty" gorm:"many2many:assign_to_users;"`
	Lessons     []Lesson     `json:"lessons,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	Status      uint8        `json:"status,omitempty"`
	UpdatedAt   time.Time    `json:"updated_at,omitempty"`
	CreatedAt   time.Time    `json:"created_at,omitempty"`
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
