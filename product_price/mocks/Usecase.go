// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import models "github.com/soerjadi/exam/models"

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, price
func (_m *Usecase) Create(ctx context.Context, price *models.ProductPrice) error {
	ret := _m.Called(ctx, price)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.ProductPrice) error); ok {
		r0 = rf(ctx, price)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByProductID provides a mock function with given fields: ctx, id
func (_m *Usecase) DeleteByProductID(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByProductID provides a mock function with given fields: ctx, id
func (_m *Usecase) GetByProductID(ctx context.Context, id int64) ([]*models.ProductPrice, error) {
	ret := _m.Called(ctx, id)

	var r0 []*models.ProductPrice
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*models.ProductPrice); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ProductPrice)
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

// GetPriceByAmount provides a mock function with given fields: ctx, amount
func (_m *Usecase) GetPriceByAmount(ctx context.Context, amount int64) (*models.ProductPrice, error) {
	ret := _m.Called(ctx, amount)

	var r0 *models.ProductPrice
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.ProductPrice); ok {
		r0 = rf(ctx, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProductPrice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
