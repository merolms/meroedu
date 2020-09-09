package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

// AttachmentUseCase ...
type AttachmentUseCase struct {
	attachmentStore domain.AttachmentStorage
	attachmentRepo  domain.AttachmentRepository
	contextTimeOut  time.Duration
}

// NewAttachmentUseCase ...
func NewAttachmentUseCase(a domain.AttachmentRepository, store domain.AttachmentStorage, timeout time.Duration) domain.AttachmentUserCase {
	return &AttachmentUseCase{
		attachmentRepo:  a,
		attachmentStore: store,
		contextTimeOut:  timeout,
	}
}

// CreateAttachment ...
func (usecase *AttachmentUseCase) CreateAttachment(ctx context.Context, attachment domain.Attachment) (*domain.Attachment, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.contextTimeOut)
	defer cancel()
	filename := getFileName(attachment.Type)
	if filename == nil {
		return nil, domain.ErrUnsupportedFileType
	}
	attachment.Name = *filename
	err := usecase.attachmentStore.CreateAttachment(ctx, attachment)
	if err != nil {
		log.Errorf("Error occur %v", err)
		return nil, err
	}
	err = usecase.attachmentRepo.CreateAttachment(ctx, attachment)
	if err != nil {
		log.Errorf("Error occur %v", err)
		return nil, err
	}
	return &attachment, nil
}
func getUUID() string {
	id := uuid.New()
	return id.String()
}

// GetFileName will return file name with concating with uniquie id(uuid)
func getFileName(fileType string) *string {
	log.Infof("Requested file type:%v", fileType)
	var filename string = ""
	switch fileType {
	case "image/png":
		filename = getUUID() + ".png"
		return &filename
	case "image/jpg", "image/jpeg":
		filename = getUUID() + ".jpg"
		return &filename
	case "text/markdown":
		filename = getUUID() + ".md"
		return &filename
	case "text/html":
		filename = getUUID() + ".html"
		return &filename
	case "video/mp4":
		filename = getUUID() + ".mp4"
		return &filename
	}
	return nil
}
