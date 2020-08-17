package domain

import (
	"context"
	"time"
)

// User ...
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// UserRepository represent the User's repository contract
type UserRepository interface {
	GetByID(ctx context.Context, id int64) (User, error)
}
