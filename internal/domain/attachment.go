package domain

import (
	"context"
	"mime/multipart"
	"time"
)

// Attachment ...
type Attachment struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	File      multipart.File `json:"file,omitempty"`
	Type      string         `json:"file_type,omitempty"`
	Filename  string         `json:"filename,omitempty"`
	Size      int64          `json:"file_size,omitempty"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
}

// AttachmentUserCase represents attachments usecase contract
type AttachmentUserCase interface {
	Upload(ctx context.Context, attachment Attachment) error
}

// AttachmentRepository represent the attachment's repository contract
type AttachmentRepository interface {
	// Upload(ctx context.Context, a *Attachment)
}
