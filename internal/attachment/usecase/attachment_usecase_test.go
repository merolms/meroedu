package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/meroedu/meroedu/internal/attachment/usecase"
	"github.com/meroedu/meroedu/internal/domain/mocks"
)

func createFile(filename string) (*os.File, error) {
	rootDirectory, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path := rootDirectory + "/" + filename
	fmt.Println(path)
	dst, err := os.Create(path)
	if err != nil {
		log.Errorf("Error occur while creating file from path: %v, Error: %v", path, err)
		return nil, err
	}
	err = os.Remove(path)
	if err != nil {
		log.Errorf("Error occur while removing aile from path: %v, Error: %v", path, err)
	}
	defer dst.Close()
	file, err := os.Open(path)
	if err != nil {
		log.Errorf("Error while opeing file: %v", err)
	}
	defer file.Close()
	return file, nil
}
func TestCreateAttachment(t *testing.T) {
	mockAttachmentStore := new(mocks.AttachmentStorage)
	mockAttachmentRepo := new(mocks.AttachmentRepository)
	mockAttachment := domain.Attachment{
		ID:   1,
		Name: "123.md",
	}
	filetypes := []string{"image/png", "image/jpg", "text/markdown", "text/html"}
	t.Run("success", func(t *testing.T) {
		file, err := createFile("meroedu.png")
		for _, filetype := range filetypes {
			if err != nil {
				t.Errorf("Error creating temp file %v", err)
			}
			mockAttachment.File = file
			mockAttachment.Type = filetype
			mockAttachmentStore.On("CreateAttachment", mock.Anything, mock.AnythingOfType("domain.Attachment")).Return(nil).Once()
			u := usecase.NewAttachmentUseCase(mockAttachmentRepo, mockAttachmentStore, time.Second*2)
			a, err := u.CreateAttachment(context.TODO(), mockAttachment)
			assert.NoError(t, err)
			assert.Equal(t, mockAttachment.ID, a.ID)
			mockAttachmentStore.AssertExpectations(t)
		}
	})
	t.Run("error-failed", func(t *testing.T) {
		mockAttachment.Type = "text/xml"
		u := usecase.NewAttachmentUseCase(mockAttachmentRepo, mockAttachmentStore, time.Second*2)
		a, err := u.CreateAttachment(context.TODO(), mockAttachment)
		assert.Error(t, err)
		assert.Nil(t, a)

	})
	t.Run("error-saved", func(t *testing.T) {
		file, err := createFile("meroedu.html")
		if err != nil {
			t.Errorf("Error creating temp file %v", err)
		}
		mockAttachment.File = file
		mockAttachment.Type = "text/html"
		mockAttachmentStore.On("CreateAttachment", mock.Anything, mock.AnythingOfType("domain.Attachment")).Return(errors.New("Error occur while saving file")).Once()
		u := usecase.NewAttachmentUseCase(mockAttachmentRepo, mockAttachmentStore, time.Second*2)
		a, err := u.CreateAttachment(context.TODO(), mockAttachment)
		assert.Error(t, err)
		assert.Nil(t, a)
		mockAttachmentStore.AssertExpectations(t)
	})
}
