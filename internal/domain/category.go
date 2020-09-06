package domain

import (
	"context"
	"time"
)

// Category represent Category Category
type Category struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name,omitempty"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

// CategoryUseCase represent the Category's repository contract
type CategoryUseCase interface {
	GetAll(ctx context.Context, start int, limit int) ([]Category, error)
	GetByID(ctx context.Context, id int64) (*Category, error)
	GetByName(ctx context.Context, name string) (*Category, error)
	UpdateCategory(ctx context.Context, Category *Category, id int64) error
	CreateCategory(ctx context.Context, Category *Category) error
	DeleteCategory(ctx context.Context, id int64) error
}

// CategoryRepository represent the Category's repository
type CategoryRepository interface {
	GetAll(ctx context.Context, start int, limit int) ([]Category, error)
	GetByID(ctx context.Context, id int64) (*Category, error)
	GetByName(ctx context.Context, name string) (*Category, error)
	UpdateCategory(ctx context.Context, Category *Category) error
	CreateCategory(ctx context.Context, Category *Category) error
	DeleteCategory(ctx context.Context, id int64) error
}
