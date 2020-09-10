package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/meroedu/meroedu/internal/attachment/usecase"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/domain/mocks"
	"github.com/meroedu/meroedu/pkg/log"
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
		log.Errorf("error occur while creating file from path: %v, error: %v", path, err)
		return nil, err
	}
	defer dst.Close()
	file, err := os.Open(path)
	if err != nil {
		log.Errorf("error while opeing file: %v", err)
		return nil, err
	}
	err = os.Remove(path)
	if err != nil {
		log.Errorf("error occur while removing aile from path: %v, error: %v", path, err)
	}
	defer file.Close()
	return file, nil
}
func TestCreateAttachment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockAttachmentStore := new(mocks.AttachmentStorage)
		mockAttachmentRepo := new(mocks.AttachmentRepository)
		mockAttachment := domain.Attachment{
			ID:   1,
			Name: "123.md",
		}
		filetypes := []string{"image/png", "image/jpg", "text/markdown", "text/html", "video/mp4", "image/jpeg"}
		file, err := createFile("meroedu.png")
		if err != nil {
			t.Errorf("error creating temp file %v", err)
		}
		for _, filetype := range filetypes {
			mockAttachment.File = file
			mockAttachment.Type = filetype
			mockAttachmentStore.On("CreateAttachment", mock.Anything, mock.AnythingOfType("domain.Attachment")).Return(nil).Once()
			mockAttachmentRepo.On("CreateAttachment", mock.Anything, mock.AnythingOfType("domain.Attachment")).Return(nil).Once()
			u := usecase.NewAttachmentUseCase(mockAttachmentRepo, mockAttachmentStore, time.Second*2)
			a, err := u.CreateAttachment(context.TODO(), mockAttachment)
			assert.NoError(t, err)
			assert.Equal(t, mockAttachment.ID, a.ID)
			mockAttachmentStore.AssertExpectations(t)
			mockAttachmentRepo.AssertExpectations(t)
		}
	})
	t.Run("error-store", func(t *testing.T) {
		mockAttachmentStore := new(mocks.AttachmentStorage)
		mockAttachmentRepo := new(mocks.AttachmentRepository)
		mockAttachment := domain.Attachment{
			ID:   1,
			Name: "123.md",
			Type: "text/xml",
		}
		u := usecase.NewAttachmentUseCase(mockAttachmentRepo, mockAttachmentStore, time.Second*2)
		a, err := u.CreateAttachment(context.TODO(), mockAttachment)
		assert.Error(t, err)
		assert.Nil(t, a)
	})
	t.Run("error-db-saved", func(t *testing.T) {
		mockAttachmentStore := new(mocks.AttachmentStorage)
		mockAttachmentRepo := new(mocks.AttachmentRepository)
		mockAttachment := domain.Attachment{
			ID:   1,
			Name: "123.png",
			Type: "image/png",
		}
		mockAttachmentStore.On("CreateAttachment", mock.Anything, mock.AnythingOfType("domain.Attachment")).Return(nil).Once()
		mockAttachmentRepo.On("CreateAttachment", mock.Anything, mock.AnythingOfType("domain.Attachment")).Return(errors.New("unexpected to save in database")).Once()
		u := usecase.NewAttachmentUseCase(mockAttachmentRepo, mockAttachmentStore, time.Second*2)
		a, err := u.CreateAttachment(context.TODO(), mockAttachment)
		assert.Error(t, err)
		assert.Nil(t, a)
		mockAttachmentStore.AssertExpectations(t)
		mockAttachmentRepo.AssertExpectations(t)

	})
	t.Run("error-store-saved", func(t *testing.T) {
		mockAttachmentStore := new(mocks.AttachmentStorage)
		mockAttachmentRepo := new(mocks.AttachmentRepository)

		file, err := createFile("meroedu.html")
		if err != nil {
			t.Errorf("error creating temp file %v", err)
		}
		mockAttachmentStore.On("CreateAttachment", mock.Anything, mock.AnythingOfType("domain.Attachment")).Return(errors.New("error occur while saving file")).Once()
		mockAttachment := domain.Attachment{
			ID:   1,
			Name: "123.md",
			File: file,
			Type: "text/html",
		}
		u := usecase.NewAttachmentUseCase(mockAttachmentRepo, mockAttachmentStore, time.Second*2)
		a, err := u.CreateAttachment(context.TODO(), mockAttachment)
		assert.Error(t, err)
		assert.Nil(t, a)
		mockAttachmentStore.AssertExpectations(t)
		mockAttachmentRepo.AssertExpectations(t)
	})
}
