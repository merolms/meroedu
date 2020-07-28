package domain

import (
	"context"
	"time"
)

// Tag ...
type Tag struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name,omitempty" validator:"required"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// TagRepository represent the Tag's repository contract
type TagRepository interface {
	GetByID(ctx context.Context, id int64) (Tag, error)
}
