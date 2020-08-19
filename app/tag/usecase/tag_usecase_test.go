package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/meroedu/meroedu/app/domain"
	"github.com/meroedu/meroedu/app/domain/mocks"
	ucase "github.com/meroedu/meroedu/app/tag/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	mockTagRepo := new(mocks.TagRepository)
	// mockUserRepo := new(mocks.UserRepository)
	// mockLessonRepo := new(mocks.LessonRepository)
	// mockAttachmentRepo:=new(mocks.AttachmentRepository)
	// mockCategoryRepo:=new(mocks.CategoryRepository)
	mockListTag := []domain.Tag{
		domain.Tag{
			ID: 1, Name: "title-1",
		},
	}
	t.Run("success", func(t *testing.T) {
		mockTagRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(mockListTag, nil).Once()

		start := int(0)
		limit := int(1)
		u := ucase.NewTagUseCase(mockTagRepo, time.Second*2)
		list, err := u.GetAll(context.TODO(), start, limit)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListTag))
		mockTagRepo.AssertExpectations(t)

	})
	t.Run("error-failed", func(t *testing.T) {
		mockTagRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, errors.New("Unexpexted Error")).Once()

		u := ucase.NewTagUseCase(mockTagRepo, time.Second*2)
		start := int(0)
		limit := int(1)
		list, err := u.GetAll(context.TODO(), start, limit)

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockTagRepo.AssertExpectations(t)
	})
}
func TestGetByID(t *testing.T) {
	mockTagRepo := new(mocks.TagRepository)
	mockTag := domain.Tag{
		Name: "title-1",
	}
	t.Run("success", func(t *testing.T) {
		mockTagRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockTag, nil).Once()
		u := ucase.NewTagUseCase(mockTagRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockTag.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockTagRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockTagRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Tag{}, errors.New("Unexpected")).Once()

		u := ucase.NewTagUseCase(mockTagRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockTag.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.Tag{}, a)

		mockTagRepo.AssertExpectations(t)
	})
}

// func TestGetByTitle(t *testing.T) {
// 	mockTagRepo := new(mocks.TagRepository)
// 	mockTag := domain.Tag{
// 		Name: "title-1",
// 	}
// 	t.Run("success", func(t *testing.T) {
// 		mockTagRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(mockTag, nil).Once()
// 		u := ucase.NewTagUseCase(mockTagRepo, time.Second*2)

// 		a, err := u.GetByTitle(context.TODO(), mockTag.Title)

// 		assert.NoError(t, err)
// 		assert.NotNil(t, a)

// 		mockTagRepo.AssertExpectations(t)
// 	})
// 	t.Run("error-failed", func(t *testing.T) {
// 		mockTagRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(domain.Tag{}, errors.New("Unexpected")).Once()

// 		u := ucase.NewTagUseCase(mockTagRepo, time.Second*2)

// 		a, err := u.GetByTitle(context.TODO(), "random")

// 		assert.Error(t, err)
// 		assert.Equal(t, domain.Tag{}, a)

// 		mockTagRepo.AssertExpectations(t)
// 	})
// }
func TestCreateTag(t *testing.T) {
	mockTagRepo := new(mocks.TagRepository)
	mockTag := domain.Tag{
		Name: "Hello",
	}

	t.Run("success", func(t *testing.T) {
		tempmockTag := mockTag
		tempmockTag.ID = 0
		// mockTagRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(domain.Tag{}, domain.ErrNotFound).Once()
		mockTagRepo.On("CreateTag", mock.Anything, mock.AnythingOfType("*domain.Tag")).Return(nil).Once()
		u := ucase.NewTagUseCase(mockTagRepo, time.Second*2)

		err := u.CreateTag(context.TODO(), &tempmockTag)

		assert.NoError(t, err)
		assert.Equal(t, mockTag.Name, tempmockTag.Name)
		mockTagRepo.AssertExpectations(t)
	})
	t.Run("existing-title", func(t *testing.T) {
		// mockTagRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(existingCourse, nil).Once()
		mockTagRepo.On("CreateTag", mock.Anything, mock.AnythingOfType("*domain.Tag")).Return(nil).Once()
		u := ucase.NewTagUseCase(mockTagRepo, time.Second*2)

		err := u.CreateTag(context.TODO(), &mockTag)

		assert.NoError(t, err)
		mockTagRepo.AssertExpectations(t)
	})
}
func TestUpdateTag(t *testing.T) {
	mockTagRepo := new(mocks.TagRepository)
	mockTag := domain.Tag{
		ID:   1,
		Name: "Hello",
	}

	t.Run("success", func(t *testing.T) {
		tempmockTag := mockTag
		mockTagRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Tag{}, domain.ErrNotFound).Once()
		mockTagRepo.On("UpdateTag", mock.Anything, mock.AnythingOfType("*domain.Tag")).Return(nil).Once()
		u := ucase.NewTagUseCase(mockTagRepo, time.Second*2)

		err := u.UpdateTag(context.TODO(), &tempmockTag, tempmockTag.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockTag.ID, tempmockTag.ID)
		mockTagRepo.AssertExpectations(t)
	})
	t.Run("existing-title", func(t *testing.T) {
		existingCourse := mockTag
		mockTagRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(existingCourse, nil).Once()
		mockTagRepo.On("UpdateTag", mock.Anything, mock.AnythingOfType("*domain.Tag")).Return(nil).Once()
		u := ucase.NewTagUseCase(mockTagRepo, time.Second*2)

		err := u.UpdateTag(context.TODO(), &mockTag, existingCourse.ID)

		assert.NoError(t, err)
		mockTagRepo.AssertExpectations(t)
	})
}
