package usecase

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
)

// AttachmentUseCase ...
type AttachmentUseCase struct {
	attachmentRepo domain.AttachmentRepository
	contextTimeOut time.Duration
}

// NewAttachmentUseCase ...
func NewAttachmentUseCase(a domain.AttachmentRepository, timeout time.Duration) domain.AttachmentUserCase {
	return &AttachmentUseCase{
		attachmentRepo: a,
		contextTimeOut: timeout,
	}
}

// Upload ...
func (usecase *AttachmentUseCase) Upload(ctx context.Context, attachment domain.Attachment) error {
	ctx, cancel := context.WithTimeout(ctx, usecase.contextTimeOut)
	defer cancel()

	src := attachment.File

	// Get project Root Directory
	rootDirectory, er := os.Getwd()
	if er != nil {
		return er
	}

	// generate random file
	// id := uuid.New()

	filePath := rootDirectory + "/uploads/" + attachment.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
