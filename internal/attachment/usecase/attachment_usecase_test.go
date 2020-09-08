package usecase_test

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
	"github.com/stretchr/testify/assert"

	"github.com/meroedu/meroedu/internal/attachment/usecase"
	"github.com/meroedu/meroedu/internal/domain/mocks"
)

func createFile() {
	path := "README.md" //The path to upload the file
	file, err := os.Open(path)
	if err != nil {
		log.Errorf("Error while opeing file: %v", err)
	}

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("my_file", filepath.Base(path))
	if err != nil {
		writer.Close()
	}
	io.Copy(part, file)

}
func TestCreateAttachment(t *testing.T) {
	createFile()
	mockAttachmentStore := new(mocks.AttachmentStorage)
	mockAttachmentRepo := new(mocks.AttachmentRepository)
	mockAttachment := domain.Attachment{
		ID:   1,
		Name: "123.png",
	}
	t.Run("success", func(t *testing.T) {
		mockAttachmentStore.On("CreateAttachment")
		u := usecase.NewAttachmentUseCase(mockAttachmentRepo, mockAttachmentStore, time.Second*2)
		a, err := u.CreateAttachment(context.TODO(), mockAttachment)
		assert.NoError(t, err)
		assert.Equal(t, mockAttachment.ID, a.ID)

	})
}
