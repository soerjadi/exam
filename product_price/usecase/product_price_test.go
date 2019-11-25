package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/product_price/mocks"
	"github.com/soerjadi/exam/product_price/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockProductPriceRepo := new(mocks.Repository)
	mockProductPrice := models.ProductPrice{
		Amount:    10,
		Price:     9000.0,
		ProductID: 2,
	}

	t.Run("success", func(t *testing.T) {
		tmpMockProductPrice := mockProductPrice
		tmpMockProductPrice.ID = 0
		mockProductPriceRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.ProductPrice")).Return(nil).Once()

		p := usecase.NewProductPriceUsecase(mockProductPriceRepo, time.Second*2)

		err := p.Create(context.TODO(), &tmpMockProductPrice)

		assert.NoError(t, err)
		assert.Equal(t, mockProductPrice, tmpMockProductPrice)

		mockProductPriceRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockProductPriceRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.ProductPrice")).Return(errors.New("Unexpected")).Once()

		p := usecase.NewProductPriceUsecase(mockProductPriceRepo, time.Second*2)

		err := p.Create(context.TODO(), &mockProductPrice)

		assert.Error(t, err)

		mockProductPriceRepo.AssertExpectations(t)
	})
}

func TestDeleteByProductID(t *testing.T) {
	mockProductPriceRepo := new(mocks.Repository)
	mockProductPrice := models.ProductPrice{
		Amount:    10,
		Price:     9000.0,
		ProductID: 2,
	}

	t.Run("success", func(t *testing.T) {
		mockProductPriceRepo.On("DeleteByProductID", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		p := usecase.NewProductPriceUsecase(mockProductPriceRepo, time.Second*2)
		err := p.DeleteByProductID(context.TODO(), mockProductPrice.ProductID)

		assert.NoError(t, err)
		mockProductPriceRepo.AssertExpectations(t)
	})

	t.Run("item is not exist", func(t *testing.T) {
		mockProductPriceRepo.On("GetByProductID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, models.ErrNotFound).Once()

		p := usecase.NewProductPriceUsecase(mockProductPriceRepo, time.Second*2)

		result, err := p.GetByProductID(context.TODO(), mockProductPrice.ProductID)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockProductPriceRepo.AssertExpectations(t)
	})
}

func TestGetByProductID(t *testing.T) {
	mockProductPriceRepo := new(mocks.Repository)
	mockProductPrice := models.ProductPrice{
		Amount:    10,
		Price:     9000.0,
		ProductID: 2,
	}

	t.Run("success", func(t *testing.T) {
		mockProductPriceRepo.On("GetByProductID", mock.Anything, mock.AnythingOfType("int64")).Return([]*models.ProductPrice{&mockProductPrice}, nil).Once()

		p := usecase.NewProductPriceUsecase(mockProductPriceRepo, time.Second*2)

		result, err := p.GetByProductID(context.TODO(), mockProductPrice.ProductID)

		assert.NoError(t, err)
		assert.Len(t, result, int(1))

		mockProductPriceRepo.AssertExpectations(t)
	})
}

func TestGetPriceByAmount(t *testing.T) {
	mockProductPriceRepo := new(mocks.Repository)
	mockProductPrice := models.ProductPrice{
		Amount:    10,
		Price:     9000.0,
		ProductID: 2,
	}
	mockProductPrice2 := models.ProductPrice{
		Amount:    20,
		Price:     8000.0,
		ProductID: 2,
	}

	t.Run("success", func(t *testing.T) {
		mockProductPriceRepo.On("GetPriceByAmount", mock.Anything, mock.AnythingOfType("int64")).Return(&mockProductPrice2, nil).Once()

		p := usecase.NewProductPriceUsecase(mockProductPriceRepo, time.Second*2)

		result, err := p.GetPriceByAmount(context.TODO(), int64(24))

		assert.NoError(t, err)
		assert.Equal(t, &mockProductPrice2, result)

		mockProductPriceRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockProductPriceRepo.On("GetPriceByAmount", mock.Anything, mock.AnythingOfType("int64")).Return(&mockProductPrice2, nil).Once()

		p := usecase.NewProductPriceUsecase(mockProductPriceRepo, time.Second*2)

		result, err := p.GetPriceByAmount(context.TODO(), int64(24))

		assert.NoError(t, err)
		assert.NotEqual(t, &mockProductPrice, result)

		mockProductPriceRepo.AssertExpectations(t)
	})
}
