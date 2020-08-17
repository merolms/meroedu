package domain

import (
	"context"
	"time"
)

// Tag ...
type Tag struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name,omitempty"`
	UpdatedAt time.Time `json:"-,omitempty"`
	CreatedAt time.Time `json:"-,omitempty"`
}

// TagRepository represent the Tag's repository contract
type TagRepository interface {
	GetByID(ctx context.Context, id int64) (Tag, error)
}
