package usecase

import (
	"context"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

// CategoryUseCase ...
type CategoryUseCase struct {
	categoryRepo   domain.CategoryRepository
	contextTimeOut time.Duration
}

// NewCategoryUseCase will creae new an
func NewCategoryUseCase(c domain.CategoryRepository, timeout time.Duration) domain.CategoryUseCase {
	return &CategoryUseCase{
		categoryRepo:   c,
		contextTimeOut: timeout,
	}
}

// GetAll ...
func (usecase *CategoryUseCase) GetAll(c context.Context, start int, limit int) (res []domain.Category, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.categoryRepo.GetAll(ctx, start, limit)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID ...
func (usecase *CategoryUseCase) GetByID(c context.Context, id int64) (res domain.Category, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Category{}, err
	}

	return res, nil
}

// // GetByTitle ...
// func (usecase *CategoryUseCase) GetByTitle(c context.Context, title string) (res domain.Category, err error) {
// 	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
// 	defer cancel()
// 	res, err = usecase.categoryRepo.GetByTitle(ctx, title)
// 	if err != nil {
// 		return domain.Category{}, err
// 	}
// 	return res, nil
// }

// CreateCategory ..
func (usecase *CategoryUseCase) CreateCategory(c context.Context, category *domain.Category) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	category.UpdatedAt = time.Now()
	category.CreatedAt = time.Now()
	err = usecase.categoryRepo.CreateCategory(ctx, category)
	if err != nil {
		return
	}
	return

}

// UpdateCategory ..
func (usecase *CategoryUseCase) UpdateCategory(c context.Context, category *domain.Category, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existingCategory, err := usecase.GetByID(ctx, id)
	log.Info(existingCategory)
	log.Info(domain.Category{})
	// if existingCategory != (domain.Category{}) {
	// 	return domain.ErrConflict
	// }
	category.ID = id
	category.UpdatedAt = time.Now()
	err = usecase.categoryRepo.UpdateCategory(ctx, category)
	if err != nil {
		return
	}
	return

}
