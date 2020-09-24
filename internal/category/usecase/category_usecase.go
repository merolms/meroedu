package usecase

import (
	"context"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
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
func (usecase *CategoryUseCase) GetByID(c context.Context, id int64) (res *domain.Category, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByName ...
func (usecase *CategoryUseCase) GetByName(c context.Context, title string) (res *domain.Category, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	res, err = usecase.categoryRepo.GetByName(ctx, title)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateCategory ..
func (usecase *CategoryUseCase) CreateCategory(c context.Context, category *domain.Category) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existingCategory, err := usecase.GetByName(ctx, category.Name)
	if existingCategory != nil {
		return domain.ErrConflict
	}
	category.UpdatedAt = time.Now().Unix()
	category.CreatedAt = time.Now().Unix()
	err = usecase.categoryRepo.CreateCategory(ctx, category)
	if err != nil {
		return err
	}
	return err

}

// UpdateCategory ..
func (usecase *CategoryUseCase) UpdateCategory(c context.Context, category *domain.Category, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existingCategory, err := usecase.GetByID(ctx, id)
	if existingCategory == nil {
		return domain.ErrNotFound
	}
	category.ID = id
	category.UpdatedAt = time.Now().Unix()
	err = usecase.categoryRepo.UpdateCategory(ctx, category)
	if err != nil {
		return
	}
	return

}

// DeleteCategory ...
func (usecase *CategoryUseCase) DeleteCategory(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existedCategory, err := usecase.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existedCategory == nil {
		return domain.ErrNotFound
	}
	return usecase.categoryRepo.DeleteCategory(ctx, id)
}
