package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	ucase "github.com/meroedu/meroedu/internal/category/usecase"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	mockCategoryRepo := new(mocks.CategoryRepository)
	mockListCategory := []domain.Category{
		domain.Category{
			ID: 1, Name: "title-1",
		},
	}
	t.Run("success", func(t *testing.T) {
		mockCategoryRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(mockListCategory, nil).Once()

		start := int(0)
		limit := int(1)
		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)
		list, err := u.GetAll(context.TODO(), start, limit)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListCategory))
		mockCategoryRepo.AssertExpectations(t)

	})
	t.Run("error-failed", func(t *testing.T) {
		mockCategoryRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, errors.New("Unexpexted Error")).Once()

		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)
		start := int(0)
		limit := int(1)
		list, err := u.GetAll(context.TODO(), start, limit)

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockCategoryRepo.AssertExpectations(t)
	})
}
func TestGetByID(t *testing.T) {
	mockCategoryRepo := new(mocks.CategoryRepository)
	mockCategory := domain.Category{
		Name: "title-1",
	}
	t.Run("success", func(t *testing.T) {
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockCategory, nil).Once()
		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockCategory.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockCategoryRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected")).Once()

		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockCategory.ID)

		assert.Error(t, err)
		assert.Nil(t, a)

		mockCategoryRepo.AssertExpectations(t)
	})
}

func TestGetByName(t *testing.T) {
	mockCategoryRepo := new(mocks.CategoryRepository)
	mockCategory := domain.Category{
		Name:      "category-1",
		UpdatedAt: time.Now().Unix(),
		CreatedAt: time.Now().Unix(),
	}
	t.Run("success", func(t *testing.T) {
		mockCategoryRepo.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(&mockCategory, nil).Once()
		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)

		a, err := u.GetByName(context.TODO(), mockCategory.Name)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockCategoryRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCategoryRepo.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("Unexpected")).Once()

		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)

		a, err := u.GetByName(context.TODO(), "random")

		assert.Error(t, err)
		assert.Nil(t, a)

		mockCategoryRepo.AssertExpectations(t)
	})
}
func TestCreateCategory(t *testing.T) {
	mockCategoryRepo := new(mocks.CategoryRepository)
	mockCategory := domain.Category{
		Name: "Hello",
	}

	t.Run("success", func(t *testing.T) {
		tempmockCategory := mockCategory
		tempmockCategory.ID = 0
		mockCategoryRepo.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(nil, domain.ErrNotFound).Once()
		mockCategoryRepo.On("CreateCategory", mock.Anything, mock.AnythingOfType("*domain.Category")).Return(nil).Once()
		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)

		err := u.CreateCategory(context.TODO(), &tempmockCategory)

		assert.NoError(t, err)
		assert.Equal(t, mockCategory.Name, tempmockCategory.Name)
		mockCategoryRepo.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		mockCategoryRepo.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(&mockCategory, nil).Once()
		mockCategoryRepo.On("CreateCategory", mock.Anything, mock.AnythingOfType("*domain.Category")).Return(domain.ErrConflict).Once()
		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)

		err := u.CreateCategory(context.TODO(), &mockCategory)

		assert.Error(t, err)
		// mockCategoryRepo.AssertExpectations(t)
	})
}
func TestUpdateCategory(t *testing.T) {
	mockCategoryRepo := new(mocks.CategoryRepository)
	mockCategory := domain.Category{
		ID:   1,
		Name: "Hello",
	}

	t.Run("success", func(t *testing.T) {
		tempmockCategory := mockCategory
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockCategory, nil).Once()
		mockCategoryRepo.On("UpdateCategory", mock.Anything, mock.AnythingOfType("*domain.Category")).Return(nil).Once()
		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)

		err := u.UpdateCategory(context.TODO(), &tempmockCategory, tempmockCategory.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockCategory.ID, tempmockCategory.ID)
		mockCategoryRepo.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		category := mockCategory
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil).Once()
		mockCategoryRepo.On("UpdateCategory", mock.Anything, mock.AnythingOfType("*domain.Category")).Return(domain.ErrNotFound).Once()
		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)

		err := u.UpdateCategory(context.TODO(), &mockCategory, category.ID)

		assert.Error(t, err)
		// mockCategoryRepo.AssertExpectations(t)
	})
}

func TestDeleteCategory(t *testing.T) {
	mockCategoryRepo := new(mocks.CategoryRepository)
	mockCategory := domain.Category{
		Name: "category",
	}

	t.Run("success", func(t *testing.T) {
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockCategory, nil).Once()

		mockCategoryRepo.On("DeleteCategory", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)

		err := u.DeleteCategory(context.TODO(), mockCategory.ID)

		assert.NoError(t, err)
		mockCategoryRepo.AssertExpectations(t)
	})
	t.Run("tag-is-not-exist", func(t *testing.T) {
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil).Once()

		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)

		err := u.DeleteCategory(context.TODO(), mockCategory.ID)

		assert.Error(t, err)
		mockCategoryRepo.AssertExpectations(t)
	})
	t.Run("error-happens-in-db", func(t *testing.T) {
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected Error")).Once()

		u := ucase.NewCategoryUseCase(mockCategoryRepo, time.Second*2)

		err := u.DeleteCategory(context.TODO(), mockCategory.ID)

		assert.Error(t, err)
		mockCategoryRepo.AssertExpectations(t)
	})

}
