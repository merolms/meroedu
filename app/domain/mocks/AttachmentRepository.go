// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/meroedu/meroedu/app/domain"
	mock "github.com/stretchr/testify/mock"
)

// AttachmentRepository is an autogenerated mock type for the AttachmentRepository type
type AttachmentRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *AttachmentRepository) GetByID(ctx context.Context, id int64) (domain.Attachment, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Attachment
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Attachment); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Attachment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
