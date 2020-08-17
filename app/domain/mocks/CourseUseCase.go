// Code generated by mockery v2.2.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/meroedu/meroedu/app/domain"
	mock "github.com/stretchr/testify/mock"
)

// CourseUseCase is an autogenerated mock type for the CourseUseCase type
type CourseUseCase struct {
	mock.Mock
}

// CreateCourse provides a mock function with given fields: ctx, course
func (_m *CourseUseCase) CreateCourse(ctx context.Context, course *domain.Course) error {
	ret := _m.Called(ctx, course)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Course) error); ok {
		r0 = rf(ctx, course)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, start, limit
func (_m *CourseUseCase) GetAll(ctx context.Context, start int, limit int) ([]domain.Course, error) {
	ret := _m.Called(ctx, start, limit)

	var r0 []domain.Course
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []domain.Course); ok {
		r0 = rf(ctx, start, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Course)
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
func (_m *CourseUseCase) GetByID(ctx context.Context, id int64) (domain.Course, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Course
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Course); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Course)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *CourseUseCase) GetByTitle(ctx context.Context, title string) (domain.Course, error) {
	ret := _m.Called(ctx, title)

	var r0 domain.Course
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Course); ok {
		r0 = rf(ctx, title)
	} else {
		r0 = ret.Get(0).(domain.Course)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
