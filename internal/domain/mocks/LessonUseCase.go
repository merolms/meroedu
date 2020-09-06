// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/meroedu/meroedu/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// LessonUseCase is an autogenerated mock type for the LessonUseCase type
type LessonUseCase struct {
	mock.Mock
}

// CreateLesson provides a mock function with given fields: ctx, Lesson
func (_m *LessonUseCase) CreateLesson(ctx context.Context, Lesson *domain.Lesson) error {
	ret := _m.Called(ctx, Lesson)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Lesson) error); ok {
		r0 = rf(ctx, Lesson)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteLesson provides a mock function with given fields: ctx, id
func (_m *LessonUseCase) DeleteLesson(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, start, limit
func (_m *LessonUseCase) GetAll(ctx context.Context, start int, limit int) ([]domain.Lesson, error) {
	ret := _m.Called(ctx, start, limit)

	var r0 []domain.Lesson
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []domain.Lesson); ok {
		r0 = rf(ctx, start, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Lesson)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, start, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *LessonUseCase) GetByID(ctx context.Context, id int64) (*domain.Lesson, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Lesson
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.Lesson); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Lesson)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLesson provides a mock function with given fields: ctx, Lesson, id
func (_m *LessonUseCase) UpdateLesson(ctx context.Context, Lesson *domain.Lesson, id int64) error {
	ret := _m.Called(ctx, Lesson, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Lesson, int64) error); ok {
		r0 = rf(ctx, Lesson, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
