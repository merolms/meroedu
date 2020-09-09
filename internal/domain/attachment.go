package domain

import (
	"context"
	"mime/multipart"
	"time"
)

// Attachment ...
type Attachment struct {
	ID          int64          `json:"id"`
	Title       string         `json:"title,omitempty"`
	Description string         `json:"description,omitempty"`
	CourseID    int64          `json:"course_id,omitempty"`
	Name        string         `json:"name"`
	File        multipart.File `json:"-" faker:"-"`
	Type        string         `json:"file_type,omitempty"`
	Filename    string         `json:"-"`
	Size        int64          `json:"file_size,omitempty"`
	UpdatedAt   time.Time      `json:"updated_at"`
	CreatedAt   time.Time      `json:"created_at"`
}

// AttachmentUserCase represents attachments usecase contract
type AttachmentUserCase interface {
	CreateAttachment(ctx context.Context, attachment Attachment) (*Attachment, error)
}

// AttachmentRepository represent the attachment's repository contract
type AttachmentRepository interface {
	CreateAttachment(ctx context.Context, attachment Attachment) error
}

// AttachmentStorage represent the attachment's storage contract
type AttachmentStorage interface {
	CreateAttachment(ctx context.Context, attachment Attachment) error
}
