package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	ucase "github.com/meroedu/meroedu/internal/content/usecase"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	mockContentRepo := new(mocks.ContentRepository)
	mockListContent := []domain.Content{
		domain.Content{
			ID: 1, Title: "title-1",
		},
	}
	t.Run("success", func(t *testing.T) {
		mockContentRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(mockListContent, nil).Once()

		start := int(0)
		limit := int(1)
		u := ucase.NewContentUseCase(mockContentRepo, time.Second*2)
		list, err := u.GetAll(context.TODO(), start, limit)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListContent))
		mockContentRepo.AssertExpectations(t)

	})
	t.Run("error-failed", func(t *testing.T) {
		mockContentRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, errors.New("Unexpexted Error")).Once()

		u := ucase.NewContentUseCase(mockContentRepo, time.Second*2)
		start := int(0)
		limit := int(1)
		list, err := u.GetAll(context.TODO(), start, limit)

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockContentRepo.AssertExpectations(t)
	})
}
func TestGetByID(t *testing.T) {
	mockContentRepo := new(mocks.ContentRepository)
	mockContent := domain.Content{
		Title: "title-1",
	}
	t.Run("success", func(t *testing.T) {
		mockContentRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockContent, nil).Once()
		u := ucase.NewContentUseCase(mockContentRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockContent.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockContentRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockContentRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Content{}, errors.New("Unexpected")).Once()

		u := ucase.NewContentUseCase(mockContentRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockContent.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.Content{}, a)

		mockContentRepo.AssertExpectations(t)
	})
}

// func TestGetByTitle(t *testing.T) {
// 	mockContentRepo := new(mocks.ContentRepository)
// 	mockContent := domain.Content{
// 		Title: "title-1",
// 	}
// 	t.Run("success", func(t *testing.T) {
// 		mockContentRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(mockContent, nil).Once()
// 		u := ucase.NewContentUseCase(mockContentRepo, time.Second*2)

// 		a, err := u.GetByTitle(context.TODO(), mockContent.Title)

// 		assert.NoError(t, err)
// 		assert.NotNil(t, a)

// 		mockContentRepo.AssertExpectations(t)
// 	})
// 	t.Run("error-failed", func(t *testing.T) {
// 		mockContentRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(domain.Content{}, errors.New("Unexpected")).Once()

// 		u := ucase.NewContentUseCase(mockContentRepo, time.Second*2)

// 		a, err := u.GetByTitle(context.TODO(), "random")

// 		assert.Error(t, err)
// 		assert.Equal(t, domain.Content{}, a)

// 		mockContentRepo.AssertExpectations(t)
// 	})
// }
func TestCreateContent(t *testing.T) {
	mockContentRepo := new(mocks.ContentRepository)
	mockContent := domain.Content{
		Title: "Hello",
	}

	t.Run("success", func(t *testing.T) {
		tempmockContent := mockContent
		tempmockContent.ID = 0
		// mockContentRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(domain.Content{}, domain.ErrNotFound).Once()
		mockContentRepo.On("CreateContent", mock.Anything, mock.AnythingOfType("*domain.Content")).Return(nil).Once()
		u := ucase.NewContentUseCase(mockContentRepo, time.Second*2)

		err := u.CreateContent(context.TODO(), &tempmockContent)

		assert.NoError(t, err)
		assert.Equal(t, mockContent.Title, tempmockContent.Title)
		mockContentRepo.AssertExpectations(t)
	})
	t.Run("existing-title", func(t *testing.T) {
		// mockContentRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(existingCourse, nil).Once()
		mockContentRepo.On("CreateContent", mock.Anything, mock.AnythingOfType("*domain.Content")).Return(nil).Once()
		u := ucase.NewContentUseCase(mockContentRepo, time.Second*2)

		err := u.CreateContent(context.TODO(), &mockContent)

		assert.NoError(t, err)
		mockContentRepo.AssertExpectations(t)
	})
}
func TestUpdateContent(t *testing.T) {
	mockContentRepo := new(mocks.ContentRepository)
	mockContent := domain.Content{
		ID:    1,
		Title: "Hello",
	}

	t.Run("success", func(t *testing.T) {
		tempmockContent := mockContent
		mockContentRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Content{}, domain.ErrNotFound).Once()
		mockContentRepo.On("UpdateContent", mock.Anything, mock.AnythingOfType("*domain.Content")).Return(nil).Once()
		u := ucase.NewContentUseCase(mockContentRepo, time.Second*2)

		err := u.UpdateContent(context.TODO(), &tempmockContent, tempmockContent.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockContent.ID, tempmockContent.ID)
		mockContentRepo.AssertExpectations(t)
	})
	t.Run("existing-title", func(t *testing.T) {
		existingCourse := mockContent
		mockContentRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(existingCourse, nil).Once()
		mockContentRepo.On("UpdateContent", mock.Anything, mock.AnythingOfType("*domain.Content")).Return(nil).Once()
		u := ucase.NewContentUseCase(mockContentRepo, time.Second*2)

		err := u.UpdateContent(context.TODO(), &mockContent, existingCourse.ID)

		assert.NoError(t, err)
		mockContentRepo.AssertExpectations(t)
	})
}
