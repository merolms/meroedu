package domain

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

// Attachment ...
type Attachment struct {
	ID        int64     `json:"id"`
	File      string    `json:"file,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// AttachmentUserCase represents attachments usecase contract
type AttachmentUserCase interface {
	Upload(echoContext echo.Context) error
}

// AttachmentRepository represent the attachment's repository contract
type AttachmentRepository interface {
	Upload(ctx context.Context, a *Attachment)
}
