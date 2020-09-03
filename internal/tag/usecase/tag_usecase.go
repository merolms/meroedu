package usecase

import (
	"context"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

// TagUseCase ...
type TagUseCase struct {
	tagRepo        domain.TagRepository
	contextTimeOut time.Duration
}

// NewTagUseCase will creae new an
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
func (usecase *TagUseCase) GetByID(c context.Context, id int64) (res domain.Tag, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.tagRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Tag{}, err
	}

	return res, nil
}

// // GetByTitle ...
// func (usecase *TagUseCase) GetByTitle(c context.Context, title string) (res domain.Tag, err error) {
// 	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
// 	defer cancel()
// 	res, err = usecase.tagRepo.GetByTitle(ctx, title)
// 	if err != nil {
// 		return domain.Tag{}, err
// 	}
// 	return res, nil
// }

// CreateTag ..
func (usecase *TagUseCase) CreateTag(c context.Context, tag *domain.Tag) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	tag.UpdatedAt = time.Now()
	tag.CreatedAt = time.Now()
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
	log.Info(existingTag)
	log.Info(domain.Tag{})
	// if existingTag != (domain.Tag{}) {
	// 	return domain.ErrConflict
	// }
	tag.ID = id
	tag.UpdatedAt = time.Now()
	err = usecase.tagRepo.UpdateTag(ctx, tag)
	if err != nil {
		return
	}
	return

}
