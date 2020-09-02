package usecase

import (
	"context"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

// LessonUseCase ...
type LessonUseCase struct {
	lessonRepo     domain.LessonRepository
	contextTimeOut time.Duration
}

// NewLessonUseCase will creae new an
func NewLessonUseCase(c domain.LessonRepository, timeout time.Duration) domain.LessonUseCase {
	return &LessonUseCase{
		lessonRepo:     c,
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
func (usecase *LessonUseCase) GetByID(c context.Context, id int64) (res domain.Lesson, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()

	res, err = usecase.lessonRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Lesson{}, err
	}

	return res, nil
}

// // GetByTitle ...
// func (usecase *LessonUseCase) GetByTitle(c context.Context, title string) (res domain.Lesson, err error) {
// 	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
// 	defer cancel()
// 	res, err = usecase.lessonRepo.GetByTitle(ctx, title)
// 	if err != nil {
// 		return domain.Lesson{}, err
// 	}
// 	return res, nil
// }

// CreateLesson ..
func (usecase *LessonUseCase) CreateLesson(c context.Context, lesson *domain.Lesson) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeOut)
	defer cancel()
	// existingLesson, err := usecase.GetByTitle(ctx, lesson.Title)
	// log.Info(existingLesson)
	// log.Info(domain.Lesson{})
	// if existingLesson != (domain.Lesson{}) {
	// 	return domain.ErrConflict
	// }
	lesson.UpdatedAt = time.Now()
	lesson.CreatedAt = time.Now()
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
	log.Info(existingLesson)
	log.Info(domain.Lesson{})
	// if existingLesson != (domain.Lesson{}) {
	// 	return domain.ErrConflict
	// }
	lesson.ID = id
	lesson.UpdatedAt = time.Now()
	err = usecase.lessonRepo.UpdateLesson(ctx, lesson)
	if err != nil {
		return
	}
	return

}
