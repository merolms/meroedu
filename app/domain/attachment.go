package domain

import (
	"context"
	"time"
)

// Attachment ...
type Attachment struct {
	ID        int64     `json:"id"`
	File      string    `json:"file,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// AttachmentRepository represent the attachment's repository contract
type AttachmentRepository interface {
	GetByID(ctx context.Context, id int64) (Attachment, error)
}
