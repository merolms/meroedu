package filesystem

import (
	"context"
	"io"
	"os"

	"github.com/meroedu/meroedu/internal/config"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

type fileStorage struct {
	path string
}

// Init will create an object that represent the content's Repository interface
func Init() (domain.ContentStorage, error) {
	rootDirectory, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path := rootDirectory + "/" + config.C.Filesystem.RelativePath + "/contents"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700)
	}
	return &fileStorage{
		path: path,
	}, nil
}

func (repo *fileStorage) CreateContent(ctx context.Context, content domain.Content) error {
	src := content.File
	if src == nil {
		return domain.ErrFileEmpty
	}
	filePath := repo.path + "/" + content.Name
	dst, err := os.Create(filePath)
	if err != nil {
		log.Errorf("error occur while creating filepath: %v, error: %v", filePath, err)
		return err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		log.Errorf("error occur while copying: %v", err)
		return err
	}
	return nil
}

func (repo *fileStorage) DownloadContent(ctx context.Context, fileName string) (string, error) {
	filePath := repo.path + "/" + fileName
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", err
	}
	return filePath, nil
}
