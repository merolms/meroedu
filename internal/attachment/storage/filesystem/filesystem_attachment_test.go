package filesystem_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	filestore "github.com/meroedu/meroedu/internal/attachment/storage/filesystem"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

func createTempFile(filename string) (*os.File, error) {
	rootDirectory, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path := rootDirectory + "/" + filename
	dst, err := os.Create(path)
	if err != nil {
		log.Errorf("error occur while creating file from path: %v, error: %v", path, err)
		return nil, err
	}
	defer dst.Close()
	file, err := os.Open(path)
	if err != nil {
		log.Errorf("error while opening file: %v", err)
	}
	return file, nil
}
func removeFile(filename string) error {
	rootDirectory, err := os.Getwd()
	if err != nil {
		return err
	}
	path := rootDirectory + "/" + filename
	err = os.Remove(path)
	if err != nil {
		log.Errorf("error occur while removing aile from path: %v, error: %v", path, err)
		return err
	}
	return nil
}
func TestCreateAttachment(t *testing.T) {
	filename := "attachment.txt"
	file, err := createTempFile(filename)
	assert.NoError(t, err)
	mockAttachment := domain.Attachment{
		ID:   1,
		Name: filename,
		File: file,
	}
	defer file.Close()
	s, err := filestore.Init()
	if err != nil {
		t.Errorf("error init filestore")
	}
	log.Infof("FileSystem storage: %v %v", s, mockAttachment)
	t.Run("success", func(t *testing.T) {
		err = s.CreateAttachment(context.TODO(), mockAttachment)
		if err != nil {
			t.Errorf("error while creating attachment %v", err)
		}
		assert.NoError(t, err)
	})
	t.Run("error-nil-file", func(t *testing.T) {
		mockAttachment.File = nil
		err = s.CreateAttachment(context.TODO(), mockAttachment)
		assert.Error(t, err)
	})

	err = removeFile(filename)
	if err != nil {
		t.Errorf("error removing %v", filename)
	}
}

func TestDownloadAttachment(t *testing.T) {
	filename := "attachment.txt"
	file, err := createTempFile(filename)
	assert.NoError(t, err)
	defer file.Close()
	s, err := filestore.Init()
	if err != nil {
		t.Errorf("error init filestore")
	}
	t.Run("success", func(t *testing.T) {
		path, err := s.DownloadAttachment(context.TODO(), filename)
		if err != nil {
			t.Errorf("error while creating attachment %v", err)
		}
		assert.NoError(t, err)
		assert.Contains(t, path, filename)
	})
	t.Run("error-nil-file", func(t *testing.T) {
		path, err := s.DownloadAttachment(context.TODO(), "abc.txt")
		assert.Error(t, err)
		assert.Empty(t, path)
	})

	err = removeFile(filename)
	if err != nil {
		t.Errorf("error removing %v", filename)
	}
}
