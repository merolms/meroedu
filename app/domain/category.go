package domain

import (
	"context"
	"time"
)

// Category represent Course Category
type Category struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name,omitempty"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

// CategoryRepository represent the Category's repository contract
type CategoryRepository interface {
	GetByID(ctx context.Context, id int64) (Category, error)
}
