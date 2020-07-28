package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/meroedu/course-api/app/domain"
)

// CourseUseCase ...
type CourseUseCase struct {
	courseRepo     domain.CourseRepository
	userRepo       domain.UserRepository
	lessonRepo     domain.LessonRepository
	attachmentRepo domain.AttachmentRepository
	tagRepo        domain.TagRepository
	categoryRepo   domain.CategoryRepository
	contextTimeOut time.Duration
}

// NewCourseUseCase will creae new an
func NewCourseUseCase(c domain.CourseRepository, timeout time.Duration) domain.CourseUseCase {
	return &CourseUseCase{
		courseRepo:     c,
		contextTimeOut: timeout,
	}
}

// GetAll ...
func (usecase *CourseUseCase) GetAll(c context.Context, skip int, limit int) (res []domain.Course, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.courseRepo.GetAll(ctx, skip, limit)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID ...
func (usecase *CourseUseCase) GetByID(c context.Context, id int64) (res domain.Course, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.courseRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Course{}, err
	}

	return res, nil
}

// GetByTitle ...
func (usecase *CourseUseCase) GetByTitle(c context.Context, title string) (res domain.Course, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	res, err = usecase.courseRepo.GetByTitle(ctx, title)
	if err != nil {
		return domain.Course{}, nil
	}
	return res, nil
}

// CreateCourse ..
func (usecase *CourseUseCase) CreateCourse(c context.Context, course *domain.Course) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existedCourse, err := usecase.GetByTitle(ctx, course.Title)
	fmt.Printf("%v, %T\n", existedCourse, existedCourse)
	// if existedCourse != (domain.Course{}) {
	// 	return domain.ErrConflict
	// }
	usecase.courseRepo.CreateCourse(ctx, course)
	return

}
