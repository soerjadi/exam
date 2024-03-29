// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import models "github.com/soerjadi/exam/models"

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Compare provides a mock function with given fields: ctx, id1, id2
func (_m *Usecase) Compare(ctx context.Context, id1 int64, id2 int64) ([]*models.Product, error) {
	ret := _m.Called(ctx, id1, id2)

	var r0 []*models.Product
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) []*models.Product); ok {
		r0 = rf(ctx, id1, id2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, id1, id2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *Usecase) Create(ctx context.Context, _a1 *models.Product) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Product) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Usecase) Delete(ctx context.Context, id int64) error {
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
func (_m *Usecase) GetByID(ctx context.Context, id int64) (*models.Product, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Product
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Product); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
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
func (_m *Usecase) Search(ctx context.Context, query *string, offset int64, limit int64) ([]*models.Product, int64, error) {
	ret := _m.Called(ctx, query, offset, limit)

	var r0 []*models.Product
	if rf, ok := ret.Get(0).(func(context.Context, *string, int64, int64) []*models.Product); ok {
		r0 = rf(ctx, query, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Product)
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
func (_m *Usecase) Update(ctx context.Context, _a1 *models.Product) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Product) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
