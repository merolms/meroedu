package usecase

import (
	// "context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/meroedu/meroedu/app/domain"
)

// AttachmentUseCase ...
type AttachmentUseCase struct {
	attachmentRepo domain.AttachmentRepository
	contextTimeOut time.Duration
}

// NewCourseUseCase will creae new an
func NewCourseUseCase(a domain.AttachmentRepository, timeout time.Duration) domain.AttachmentUserCase {
	return &AttachmentUseCase{
		attachmentRepo: a,
		contextTimeOut: timeout,
	}
}

// Upload ...
func (usecase *AttachmentUseCase) Upload(echoContext echo.Context) error {
	// ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	// defer cancel()
	file, err := echoContext.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	path, er := os.Getwd()
	if er != nil {
		return er
	}
	fmt.Println(path)

	filePath := "uploads/" + file.Filename

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
