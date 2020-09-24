package domain

import (
	"context"
)

// Content ...
type Content struct {
	ID          int64  `json:"id,omitempty"`
	LessonID    int64  `json:"lesson_id,omitempty"`
	Title       string `json:"title" validator:"required"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type" validator:"required"`
	Content     string `json:"content" validator:"required"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
}

// ContentUseCase represent the Content's repository contract
type ContentUseCase interface {
	GetAll(ctx context.Context, start int, limit int) ([]Content, error)
	GetByID(ctx context.Context, id int64) (*Content, error)
	UpdateContent(ctx context.Context, Content *Content, id int64) error
	CreateContent(ctx context.Context, Content *Content) error
	DeleteContent(ctx context.Context, id int64) error
	GetContentByLesson(ctx context.Context, lessonID int64) ([]Content, error)
}

// ContentRepository represent the Content's repository
type ContentRepository interface {
	GetAll(ctx context.Context, start int, limit int) ([]Content, error)
	GetByID(ctx context.Context, id int64) (*Content, error)
	UpdateContent(ctx context.Context, Content *Content) error
	CreateContent(ctx context.Context, Content *Content) error
	DeleteContent(ctx context.Context, id int64) error
	GetContentCountByLesson(ctx context.Context, lessonID int64) (int, error)
	GetContentByLesson(ctx context.Context, lessonID int64) ([]Content, error)
}
