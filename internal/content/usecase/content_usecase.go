package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

// ContentUseCase ...
type ContentUseCase struct {
	contentStore   domain.ContentStorage
	contentRepo    domain.ContentRepository
	contextTimeOut time.Duration
}

// NewContentUseCase will create new an
func NewContentUseCase(c domain.ContentRepository, s domain.ContentStorage, timeout time.Duration) domain.ContentUseCase {
	return &ContentUseCase{
		contentRepo:    c,
		contentStore:   s,
		contextTimeOut: timeout,
	}
}

// GetAll ...
func (usecase *ContentUseCase) GetAll(c context.Context, start int, limit int) (res []domain.Content, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.contentRepo.GetAll(ctx, start, limit)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID ...
func (usecase *ContentUseCase) GetByID(c context.Context, id int64) (res *domain.Content, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.contentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// CreateContent ..
func (usecase *ContentUseCase) CreateContent(c context.Context, content *domain.Content) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	content.UpdatedAt = time.Now().Unix()
	content.CreatedAt = time.Now().Unix()
	err = usecase.contentRepo.CreateContent(ctx, content)
	if err != nil {
		return
	}
	return
}

// UpdateContent ..
func (usecase *ContentUseCase) UpdateContent(c context.Context, content *domain.Content, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existingContent, err := usecase.GetByID(ctx, id)
	if existingContent == nil {
		return domain.ErrNotFound
	}
	content.ID = id
	content.UpdatedAt = time.Now().Unix()
	err = usecase.contentRepo.UpdateContent(ctx, content)
	if err != nil {
		return
	}
	return
}

// DeleteContent ...
func (usecase *ContentUseCase) DeleteContent(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existedTag, err := usecase.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existedTag == nil {
		return domain.ErrNotFound
	}
	return usecase.contentRepo.DeleteContent(ctx, id)
}

// GetContentByLesson ...
func (usecase *ContentUseCase) GetContentByLesson(c context.Context, LessonID int64) ([]domain.Content, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err := usecase.contentRepo.GetContentByLesson(ctx, LessonID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetFileName will return file name with concating with unique id(uuid)
func getFileName(fileType string) string {
	log.Infof("Requested file type:%v", fileType)
	var filename string = uuid.New().String()
	switch fileType {
	case "image/png":
		return filename + ".png"
	case "image/jpg", "image/jpeg":
		return filename + ".jpg"
	case "text/markdown":
		return filename + ".md"
	case "application/pdf":
		return filename + ".pdf"
	case "text/html":
		return filename + ".html"
	case "video/mp4":
		return filename + ".mp4"
	}
	return ""
}

// DownloadContent will return filepath as string
func (usecase *ContentUseCase) DownloadContent(ctx context.Context, fileName string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.contextTimeOut)
	defer cancel()
	filePath, err := usecase.contentStore.DownloadContent(ctx, fileName)
	if err != nil {
		log.Errorf("error occur %v", err)
		return "", err
	}
	return filePath, nil
}
