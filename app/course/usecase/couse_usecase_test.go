package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	ucase "github.com/meroedu/course-api/app/course/usecase"
	"github.com/meroedu/course-api/app/domain"
	"github.com/meroedu/course-api/app/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	// mockUserRepo := new(mocks.UserRepository)
	// mockLessonRepo := new(mocks.LessonRepository)
	// mockAttachmentRepo:=new(mocks.AttachmentRepository)
	// mockCategoryRepo:=new(mocks.CategoryRepository)
	mockListCourse := []domain.Course{
		domain.Course{
			ID: 1, Title: "title-1",
			Author: domain.User{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}
	t.Run("success", func(t *testing.T) {
		mockCourseRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(mockListCourse, nil).Once()

		start := int(0)
		limit := int(1)
		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)
		list, err := u.GetAll(context.TODO(), start, limit)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListCourse))
		mockCourseRepo.AssertExpectations(t)

	})
	t.Run("error-failed", func(t *testing.T) {
		mockCourseRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, errors.New("Unexpexted Error")).Once()

		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)
		start := int(0)
		limit := int(1)
		list, err := u.GetAll(context.TODO(), start, limit)

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockCourseRepo.AssertExpectations(t)
	})
}
func TestGetByID(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockCourse := domain.Course{
		Title:  "title-1",
		Author: domain.User{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
	}
	t.Run("success", func(t *testing.T) {
		mockCourseRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockCourse, nil).Once()
		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockCourse.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockCourseRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCourseRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Course{}, errors.New("Unexpected")).Once()

		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockCourse.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.Course{}, a)

		mockCourseRepo.AssertExpectations(t)
	})
}
func TestGetByTitle(t *testing.T) {

}
func TestCreateCourse(t *testing.T) {

}
func TestUpdateCourse(t *testing.T) {

}
