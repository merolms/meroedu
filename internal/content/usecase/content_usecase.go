package usecase

import (
	"context"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
)

// ContentUseCase ...
type ContentUseCase struct {
	contentRepo    domain.ContentRepository
	contextTimeOut time.Duration
}

// NewContentUseCase will creae new an
func NewContentUseCase(c domain.ContentRepository, timeout time.Duration) domain.ContentUseCase {
	return &ContentUseCase{
		contentRepo:    c,
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
	content.UpdatedAt = time.Now()
	content.CreatedAt = time.Now()
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
	content.UpdatedAt = time.Now()
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
