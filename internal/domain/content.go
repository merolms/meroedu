package domain

import (
	"context"
	"time"
)

// Content ...
type Content struct {
	ID          int64     `json:"id"`
	LessonID    int64     `json:"lesson_id"`
	Title       string    `json:"title" validator:"required"`
	Description string    `json:"description"`
	Type        string    `json:"type" validator:"required"`
	Content     string    `json:"content" validator:"required"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// ContentUseCase represent the Content's repository contract
type ContentUseCase interface {
	GetAll(ctx context.Context, start int, limit int) ([]Content, error)
	GetByID(ctx context.Context, id int64) (*Content, error)
	UpdateContent(ctx context.Context, Content *Content, id int64) error
	CreateContent(ctx context.Context, Content *Content) error
	DeleteContent(ctx context.Context, id int64) error
}

// ContentRepository represent the Content's repository
type ContentRepository interface {
	GetAll(ctx context.Context, start int, limit int) ([]Content, error)
	GetByID(ctx context.Context, id int64) (*Content, error)
	UpdateContent(ctx context.Context, Content *Content) error
	CreateContent(ctx context.Context, Content *Content) error
	DeleteContent(ctx context.Context, id int64) error
}
