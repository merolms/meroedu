package usecase

import (
	"context"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
)

// TagUseCase ...
type TagUseCase struct {
	tagRepo        domain.TagRepository
	contextTimeOut time.Duration
}

// NewTagUseCase will create new an
func NewTagUseCase(c domain.TagRepository, timeout time.Duration) domain.TagUseCase {
	return &TagUseCase{
		tagRepo:        c,
		contextTimeOut: timeout,
	}
}

// GetAll ...
func (usecase *TagUseCase) GetAll(c context.Context, start int, limit int) (res []domain.Tag, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.tagRepo.GetAll(ctx, start, limit)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID ...
func (usecase *TagUseCase) GetByID(c context.Context, id int64) (res *domain.Tag, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.tagRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByName ...
func (usecase *TagUseCase) GetByName(c context.Context, title string) (res *domain.Tag, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	res, err = usecase.tagRepo.GetByName(ctx, title)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateTag ..
func (usecase *TagUseCase) CreateTag(c context.Context, tag *domain.Tag) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existingTag, err := usecase.tagRepo.GetByName(ctx, tag.Name)
	if existingTag != nil {
		return domain.ErrConflict
	}
	tag.UpdatedAt = time.Now().Unix()
	tag.CreatedAt = time.Now().Unix()
	err = usecase.tagRepo.CreateTag(ctx, tag)
	if err != nil {
		return
	}
	return

}

// UpdateTag ..
func (usecase *TagUseCase) UpdateTag(c context.Context, tag *domain.Tag, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existingTag, err := usecase.GetByID(ctx, id)
	if existingTag == nil {
		return domain.ErrNotFound
	}
	tag.ID = id
	tag.UpdatedAt = time.Now().Unix()
	err = usecase.tagRepo.UpdateTag(ctx, tag)
	if err != nil {
		return err
	}
	return nil

}

// DeleteTag ...
func (usecase *TagUseCase) DeleteTag(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existedTag, err := usecase.tagRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existedTag == nil {
		return domain.ErrNotFound
	}
	return usecase.tagRepo.DeleteTag(ctx, id)
}
