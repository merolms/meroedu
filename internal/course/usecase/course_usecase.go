package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
)

// CourseUseCase ...
type CourseUseCase struct {
	courseRepo domain.CourseRepository
	userRepo   domain.UserRepository
	lessonRepo domain.LessonRepository
	// attachmentRepo domain.AttachmentRepository
	tagRepo        domain.TagRepository
	categoryRepo   domain.CategoryRepository
	contextTimeOut time.Duration
}

// NewCourseUseCase will create new an
func NewCourseUseCase(c domain.CourseRepository, timeout time.Duration) domain.CourseUseCase {
	return &CourseUseCase{
		courseRepo:     c,
		contextTimeOut: timeout,
	}
}

// GetAll ...
func (usecase *CourseUseCase) GetAll(c context.Context, start int, limit int) (res []domain.Course, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	// count, err := usecase.courseRepo.GetCourseCount(ctx)
	// log.Info(count)
	res, err = usecase.courseRepo.GetAll(ctx, start, limit)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID ...
func (usecase *CourseUseCase) GetByID(c context.Context, id int64) (*domain.Course, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	course, err := usecase.courseRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return course, nil
}

// GetByTitle ...
func (usecase *CourseUseCase) GetByTitle(c context.Context, title string) (*domain.Course, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	res, err := usecase.courseRepo.GetByTitle(ctx, title)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateCourse ..
func (usecase *CourseUseCase) CreateCourse(c context.Context, course *domain.Course) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existedCourse, err := usecase.GetByTitle(ctx, course.Title)
	fmt.Println(existedCourse)
	if existedCourse != nil {
		return domain.ErrConflict
	}
	course.UpdatedAt = time.Now()
	course.CreatedAt = time.Now()
	err = usecase.courseRepo.CreateCourse(ctx, course)
	if err != nil {
		return
	}
	return

}

// UpdateCourse ..
func (usecase *CourseUseCase) UpdateCourse(c context.Context, course *domain.Course, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existedCourse, err := usecase.GetByID(ctx, id)
	if &existedCourse == nil {
		return domain.ErrNotFound
	}
	course.ID = id
	course.UpdatedAt = time.Now()
	err = usecase.courseRepo.UpdateCourse(ctx, course)
	if err != nil {
		return
	}
	return

}

// DeleteCourse ...
func (usecase *CourseUseCase) DeleteCourse(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existedCourse, err := usecase.courseRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existedCourse == nil {
		return domain.ErrNotFound
	}
	return usecase.courseRepo.DeleteCourse(ctx, id)
}
