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

// ContentRepository represent the Content's repository contract
type ContentRepository interface {
	GetByID(ctx context.Context, id int64) (Content, error)
}
