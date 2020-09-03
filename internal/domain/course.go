package domain

import (
	"context"
	"time"
)

// Status ...
type Status string

const (
	CourseInDraft   Status = "Draft"
	CourseArchived  Status = "Archived"
	CourseAssigned  Status = "Assigned"
	CoursePublished Status = "Published"
	CoursePublic    Status = "Public"
	CourseCreated   Status = "Created"
	CourseComplete  Status = "Completed"
	StatusSuccess   Status = "Success"
	StatusQueued    Status = "Queued"
	StatusSending   Status = "Sending"
	StatusUnknown   Status = "Unknown"
	StatusScheduled Status = "Scheduled"
	StatusRetry     Status = "Retrying"
	Error           Status = "Error"
)

// Course is a struct represent a created Course
type Course struct {
	ID          int64        `json:"id" `
	Title       string       `json:"title" validate:"required"`
	Description string       `json:"description,omitempty"`
	ImageURL    string       `json:"image_url,omitempty"`
	Duration    uint16       `json:"duration,omitempty"`
	CategoryID  NullInt64    `json:"-,omitempty"`
	Category    Category     `json:"categories,omitempty"`
	Tags        []Tag        `json:"tags,omitempty"`
	AuthorID    NullInt64    `json:"-,omitempty"`
	Author      User         `json:"author,omitempty"`
	Users       []User       `json:"users,omitempty"`
	Lessons     []Lesson     `json:"lessons,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	Status      Status       `json:"status,omitempty"`
	UpdatedAt   time.Time    `json:"updated_at,omitempty"`
	CreatedAt   time.Time    `json:"created_at,omitempty"`
}

// CourseStats is a struct representing the statistics for a single Course
type CourseStats struct {
	TotalEnroll    int64 `json:"total_enroll"`
	LessonCount    int64 `json:"lesson_count"`
	TotalCompleted int64 `json:"total_complete"`
	TotalAssigned  int64 `json:"total_assign"`
}

// CourseReponse is a struct representing the overview of Courses
type CourseSummaries struct {
	Total   int64    `json:"total"`
	Courses []Course `json:"Courses"`
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
	UpdateCourse(ctx context.Context, course *Course, id int64) error
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
	GetCourseCount(ctx context.Context) (int64, error)
}
