package domain

import (
	"context"
	"mime/multipart"
)

// Attachment ...
type Attachment struct {
	ID          int64          `json:"id,omitempty"`
	Title       string         `json:"title,omitempty"`
	Description string         `json:"description,omitempty"`
	CourseID    int64          `json:"course_id,omitempty"`
	Name        string         `json:"name,omitempty"`
	File        multipart.File `json:"-" faker:"-"`
	Type        string         `json:"file_type,omitempty"`
	Filename    string         `json:"-"`
	Size        int64          `json:"file_size,omitempty"`
	UpdatedAt   int64          `json:"updated_at,omitempty"`
	CreatedAt   int64          `json:"created_at,omitempty"`
}

// AttachmentUseCase represents attachments usecase contract
type AttachmentUseCase interface {
	CreateAttachment(ctx context.Context, attachment Attachment) (*Attachment, error)
	DownloadAttachment(ctx context.Context, fileName string) (string, error)
}

// AttachmentRepository represent the attachment's repository contract
type AttachmentRepository interface {
	CreateAttachment(ctx context.Context, attachment Attachment) error
	GetAttachmentByCourse(ctx context.Context, courseID int64) ([]Attachment, error)
}

// AttachmentStorage represent the attachment's storage contract
type AttachmentStorage interface {
	CreateAttachment(ctx context.Context, attachment Attachment) error
	DownloadAttachment(ctx context.Context, fileName string) (string, error)
}
