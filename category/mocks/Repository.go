// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import models "github.com/soerjadi/exam/models"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *Repository) Create(ctx context.Context, _a1 *models.Category) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Category) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetByID(ctx context.Context, id int64) (*models.Category, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Category
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Category); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Category)
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

// Search provides a mock function with given fields: ctx, query, offset, limit
func (_m *Repository) Search(ctx context.Context, query *string, offset int64, limit int64) ([]*models.Category, int64, error) {
	ret := _m.Called(ctx, query, offset, limit)

	var r0 []*models.Category
	if rf, ok := ret.Get(0).(func(context.Context, *string, int64, int64) []*models.Category); ok {
		r0 = rf(ctx, query, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Category)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context, *string, int64, int64) int64); ok {
		r1 = rf(ctx, query, offset, limit)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *string, int64, int64) error); ok {
		r2 = rf(ctx, query, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *Repository) Update(ctx context.Context, _a1 *models.Category) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Category) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
