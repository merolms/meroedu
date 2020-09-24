package usecase

import (
	"context"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
)

// LessonUseCase ...
type LessonUseCase struct {
	lessonRepo     domain.LessonRepository
	contentUseCase domain.ContentUseCase
	contextTimeOut time.Duration
}

// NewLessonUseCase will create new an
func NewLessonUseCase(l domain.LessonRepository, c domain.ContentUseCase, timeout time.Duration) domain.LessonUseCase {
	return &LessonUseCase{
		lessonRepo:     l,
		contentUseCase: c,
		contextTimeOut: timeout,
	}
}

// GetAll ...
func (usecase *LessonUseCase) GetAll(c context.Context, start int, limit int) (res []domain.Lesson, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.lessonRepo.GetAll(ctx, start, limit)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID ...
func (usecase *LessonUseCase) GetByID(c context.Context, id int64) (res *domain.Lesson, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.lessonRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// CreateLesson ..
func (usecase *LessonUseCase) CreateLesson(c context.Context, lesson *domain.Lesson) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	lesson.UpdatedAt = time.Now().Unix()
	lesson.CreatedAt = time.Now().Unix()
	err = usecase.lessonRepo.CreateLesson(ctx, lesson)
	if err != nil {
		return
	}
	return

}

// UpdateLesson ..
func (usecase *LessonUseCase) UpdateLesson(c context.Context, lesson *domain.Lesson, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existingLesson, err := usecase.GetByID(ctx, id)
	if existingLesson == nil {
		return domain.ErrNotFound
	}
	lesson.ID = id
	lesson.UpdatedAt = time.Now().Unix()
	err = usecase.lessonRepo.UpdateLesson(ctx, lesson)
	if err != nil {
		return
	}
	return

}

// DeleteLesson ...
func (usecase *LessonUseCase) DeleteLesson(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	existedCourse, err := usecase.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existedCourse == nil {
		return domain.ErrNotFound
	}
	return usecase.lessonRepo.DeleteLesson(ctx, id)
}

// GetLessonByCourse ...
func (usecase *LessonUseCase) GetLessonByCourse(c context.Context, courseID int64) ([]domain.Lesson, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err := usecase.lessonRepo.GetLessonByCourse(ctx, courseID)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(res); i++ {
		contents, err := usecase.contentUseCase.GetContentByLesson(ctx, res[i].ID)
		if err != nil {
			return nil, err
		}
		res[i].Contents = contents
	}

	return res, nil
}

// GetLessonCountByCourse ...
func (usecase *LessonUseCase) GetLessonCountByCourse(c context.Context, courseID int64) (int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err := usecase.lessonRepo.GetLessonCountByCourse(ctx, courseID)
	if err != nil {
		return 0, err
	}

	return res, nil
}
