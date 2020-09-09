package filesystem_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	filestore "github.com/meroedu/meroedu/internal/attachment/storage/filesystem"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
	"github.com/stretchr/testify/assert"
)

func createTempFile(filename string) (*os.File, error) {
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
	defer dst.Close()
	file, err := os.Open(path)
	if err != nil {
		log.Errorf("Error while opeing file: %v", err)
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
		log.Errorf("Error occur while removing aile from path: %v, Error: %v", path, err)
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
		t.Errorf("Error init filestore")
	}
	log.Infof("FileSystem storage: %v %v", s, mockAttachment)
	t.Run("success", func(t *testing.T) {
		err = s.CreateAttachment(context.TODO(), mockAttachment)
		if err != nil {
			t.Errorf("Error While creating attachment %v", err)
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
		t.Errorf("Error removing %v", filename)
	}
}
