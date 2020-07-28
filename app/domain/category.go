package domain

import (
	"context"
	"time"
)

// Category represent Course Category
type Category struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// CategoryRepository represent the Category's repository contract
type CategoryRepository interface {
	GetByID(ctx context.Context, id int64) (Category, error)
}
