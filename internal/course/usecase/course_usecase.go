package usecase

import (
	"context"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

// CourseUseCase ...
type CourseUseCase struct {
	courseRepo        domain.CourseRepository
	userRepo          domain.UserRepository
	lessonUseCase     domain.LessonUseCase
	attachmentUseCase domain.AttachmentUseCase
	tagRepo           domain.TagRepository
	categoryRepo      domain.CategoryRepository
	contextTimeOut    time.Duration
}

// NewCourseUseCase will create new an
func NewCourseUseCase(c domain.CourseRepository, l domain.LessonUseCase, a domain.AttachmentUseCase, timeout time.Duration) domain.CourseUseCase {
	return &CourseUseCase{
		courseRepo:        c,
		lessonUseCase:     l,
		attachmentUseCase: a,
		contextTimeOut:    timeout,
	}
}

// GetAll ...
func (usecase *CourseUseCase) GetAll(c context.Context, start int, limit int) (res []domain.Course, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

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
	lessonCount, err := usecase.lessonUseCase.GetLessonCountByCourse(ctx, id)
	if err != nil {
		log.Error("err")
	}
	course.LessonCount = lessonCount

	lessons, err := usecase.lessonUseCase.GetLessonByCourse(ctx, id)
	if err != nil {
		log.Error(err)
	}
	course.Lessons = lessons
	attachments, err := usecase.attachmentUseCase.GetAttachmentByCourse(ctx, id)
	if err != nil {
		log.Error(err)
	}
	course.Attachments = attachments
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
	if existedCourse != nil {
		return domain.ErrConflict
	}
	course.UpdatedAt = time.Now().Unix()
	course.CreatedAt = time.Now().Unix()
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
	course.UpdatedAt = time.Now().Unix()
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
