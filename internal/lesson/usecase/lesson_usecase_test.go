package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/domain/mocks"
	ucase "github.com/meroedu/meroedu/internal/lesson/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	mockLessonRepo := new(mocks.LessonRepository)
	mockListLesson := []domain.Lesson{
		domain.Lesson{
			ID: 1, Title: "title-1",
		},
	}
	t.Run("success", func(t *testing.T) {
		mockLessonRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(mockListLesson, nil).Once()

		start := int(0)
		limit := int(1)
		u := ucase.NewLessonUseCase(mockLessonRepo, time.Second*2)
		list, err := u.GetAll(context.TODO(), start, limit)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListLesson))
		mockLessonRepo.AssertExpectations(t)

	})
	t.Run("error-failed", func(t *testing.T) {
		mockLessonRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, errors.New("Unexpexted Error")).Once()

		u := ucase.NewLessonUseCase(mockLessonRepo, time.Second*2)
		start := int(0)
		limit := int(1)
		list, err := u.GetAll(context.TODO(), start, limit)

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockLessonRepo.AssertExpectations(t)
	})
}
func TestGetByID(t *testing.T) {
	mockLessonRepo := new(mocks.LessonRepository)
	mockLesson := domain.Lesson{
		Title: "title-1",
	}
	t.Run("success", func(t *testing.T) {
		mockLessonRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockLesson, nil).Once()
		u := ucase.NewLessonUseCase(mockLessonRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockLesson.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockLessonRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockLessonRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected")).Once()

		u := ucase.NewLessonUseCase(mockLessonRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockLesson.ID)

		assert.Error(t, err)
		assert.Nil(t, a)
		mockLessonRepo.AssertExpectations(t)
	})
}

func TestCreateLesson(t *testing.T) {
	mockLessonRepo := new(mocks.LessonRepository)
	mockLesson := domain.Lesson{
		Title: "Hello",
	}

	t.Run("success", func(t *testing.T) {
		tempmockLesson := mockLesson
		tempmockLesson.ID = 0
		mockLessonRepo.On("CreateLesson", mock.Anything, mock.AnythingOfType("*domain.Lesson")).Return(nil).Once()
		u := ucase.NewLessonUseCase(mockLessonRepo, time.Second*2)

		err := u.CreateLesson(context.TODO(), &tempmockLesson)

		assert.NoError(t, err)
		assert.Equal(t, mockLesson.Title, tempmockLesson.Title)
		mockLessonRepo.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		mockLessonRepo.On("CreateLesson", mock.Anything, mock.AnythingOfType("*domain.Lesson")).Return(errors.New("unexpected error occur")).Once()
		u := ucase.NewLessonUseCase(mockLessonRepo, time.Second*2)

		err := u.CreateLesson(context.TODO(), &mockLesson)

		assert.Error(t, err)
		mockLessonRepo.AssertExpectations(t)
	})
}
func TestUpdateLesson(t *testing.T) {
	mockLessonRepo := new(mocks.LessonRepository)
	mockLesson := domain.Lesson{
		ID:    1,
		Title: "Hello",
	}

	t.Run("success", func(t *testing.T) {
		tempmockLesson := mockLesson
		mockLessonRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&tempmockLesson, nil).Once()
		mockLessonRepo.On("UpdateLesson", mock.Anything, mock.AnythingOfType("*domain.Lesson")).Return(nil).Once()
		u := ucase.NewLessonUseCase(mockLessonRepo, time.Second*2)

		err := u.UpdateLesson(context.TODO(), &tempmockLesson, tempmockLesson.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockLesson.ID, tempmockLesson.ID)
		mockLessonRepo.AssertExpectations(t)
	})
	t.Run("error-lesson", func(t *testing.T) {
		existingLesson := mockLesson
		mockLessonRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil).Once()
		mockLessonRepo.On("UpdateLesson", mock.Anything, mock.AnythingOfType("*domain.Lesson")).Return(domain.ErrNotFound).Once()
		u := ucase.NewLessonUseCase(mockLessonRepo, time.Second*2)

		err := u.UpdateLesson(context.TODO(), &mockLesson, existingLesson.ID)

		assert.Error(t, err)
		// mockLessonRepo.AssertExpectations(t)
	})
}

func TestDeleteLesson(t *testing.T) {
	mockLessonRepo := new(mocks.LessonRepository)
	mockLesson := domain.Lesson{
		Title:       "Hello",
		Description: "Description here",
	}

	t.Run("success", func(t *testing.T) {
		mockLessonRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockLesson, nil).Once()

		mockLessonRepo.On("DeleteLesson", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		u := ucase.NewLessonUseCase(mockLessonRepo, time.Second*2)

		err := u.DeleteLesson(context.TODO(), mockLesson.ID)

		assert.NoError(t, err)
		mockLessonRepo.AssertExpectations(t)
	})
	t.Run("course-is-not-exist", func(t *testing.T) {
		mockLessonRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil).Once()

		u := ucase.NewLessonUseCase(mockLessonRepo, time.Second*2)

		err := u.DeleteLesson(context.TODO(), mockLesson.ID)

		assert.Error(t, err)
		mockLessonRepo.AssertExpectations(t)
	})
	t.Run("error-happens-in-db", func(t *testing.T) {
		mockLessonRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected Error")).Once()

		u := ucase.NewLessonUseCase(mockLessonRepo, time.Second*2)

		err := u.DeleteLesson(context.TODO(), mockLesson.ID)

		assert.Error(t, err)
		mockLessonRepo.AssertExpectations(t)
	})

}
