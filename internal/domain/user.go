package domain

import (
	"context"
)

// User ...
type User struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

// UserRepository represent the User's repository contract
type UserRepository interface {
	GetByID(ctx context.Context, id int64) (User, error)
}
