package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/order/mocks"
	"github.com/soerjadi/exam/order/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetList(t *testing.T) {
	mockOrderRepo := new(mocks.Repository)
	mockOrder1 := models.Order{
		ID:        int64(8),
		ProductID: int64(2),
		Amount:    int64(20), // with amount 20 -> 8000
		Price:     160000.0,
		Status:    models.OrderProccessed,
	}
	mockOrder2 := models.Order{
		ID:        int64(9),
		ProductID: int64(3),
		Amount:    int64(1),
		Price:     10000.0,
		Status:    models.OrderShipped,
	}

	var mockOrders = make([]*models.Order, 0)
	mockOrders = append(mockOrders, &mockOrder1)
	mockOrders = append(mockOrders, &mockOrder2)

	t.Run("success", func(t *testing.T) {
		mockOrderRepo.On("GetList", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).
			Return(mockOrders, int64(2), nil).Once()

		p := usecase.NewOrderUsecase(mockOrderRepo, time.Second*2)

		orders, _, err := p.GetList(context.TODO(), int64(0), int64(10))

		assert.NoError(t, err)
		assert.Equal(t, mockOrders, orders)

		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockOrderRepo.On("GetList", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).
			Return(nil, int64(0), models.ErrInternalServerError).Once()

		o := usecase.NewOrderUsecase(mockOrderRepo, time.Second*2)
		orders, found, err := o.GetList(context.TODO(), int64(-1), int64(-2))

		assert.Error(t, err)
		assert.Equal(t, int64(0), found)
		assert.Nil(t, orders)

		mockOrderRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockOrderRepo := new(mocks.Repository)
	mockOrder1 := models.Order{
		ProductID: 1,
		Amount:    int64(2),
		Price:     18000.0,
		Status:    models.OrderPending,
	}
	mockOrder2 := models.Order{
		ProductID: int64(100),
		Amount:    int64(10),
		Price:     85000.0,
		Status:    models.OrderPending,
	}

	t.Run("success", func(t *testing.T) {
		tmpMockOrder := mockOrder1
		tmpMockOrder.ID = 0
		mockOrderRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Order")).Return(nil).Once()

		o := usecase.NewOrderUsecase(mockOrderRepo, time.Second*2)

		err := o.Create(context.TODO(), &mockOrder1)

		assert.NoError(t, err)
		assert.Equal(t, tmpMockOrder.ID, mockOrder1.ID)

		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockOrderRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Order")).Return(models.ErrNotFound).Once()

		o := usecase.NewOrderUsecase(mockOrderRepo, time.Second*2)

		err := o.Create(context.TODO(), &mockOrder2)

		assert.Error(t, err)
		assert.Equal(t, err, models.ErrNotFound)

		mockOrderRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockOrderRepo := new(mocks.Repository)
	mockOrder2 := models.Order{
		ID:        int64(9),
		ProductID: int64(3),
		Amount:    int64(1),
		Price:     10000.0,
		Status:    models.OrderShipped,
	}

	t.Run("success", func(t *testing.T) {
		mockOrderRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		p := usecase.NewOrderUsecase(mockOrderRepo, time.Second*2)

		err := p.Delete(context.TODO(), mockOrder2.ID)

		assert.NoError(t, err)
		mockOrderRepo.AssertExpectations(t)
	})
}
