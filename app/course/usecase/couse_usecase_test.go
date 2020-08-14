package usecase_test

import (
	"testing"

	"github.com/meroedu/course-api/app/domain/mocks"
)

func TestGetAll(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockCourses := []domain.Course{
		domain.Course{
			ID: 1, Title: "title-1",
			Author: domain.User{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}
	t.Run("success", func(t *testing.T){
		mockCourseRepo.On("GetAll", mock.Anything, mock.AnythingOfType("string")
		)	
	})
}
func TestGetByID(t *testing.T) {

}
func TestGetByTitle(t *testing.T) {

}
func TestCreateCourse(t *testing.T) {

}
func TestUpdateCourse(t *testing.T) {

}
