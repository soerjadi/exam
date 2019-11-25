package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/product_category/mocks"
	"github.com/soerjadi/exam/product_category/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByProductID(t *testing.T) {
	mockCatsRepo := new(mocks.Repository)
	mockCats := make([]*models.ProductCategory, 0)

	cat1 := models.ProductCategory{
		ID:         8,
		ProductID:  8,
		CategoryID: 2,
	}
	cat2 := models.ProductCategory{
		ID:         9,
		ProductID:  8,
		CategoryID: 5,
	}
	cat3 := models.ProductCategory{
		ID:         10,
		ProductID:  8,
		CategoryID: 8,
	}

	mockCats = append(mockCats, &cat1)
	mockCats = append(mockCats, &cat2)
	mockCats = append(mockCats, &cat3)

	t.Run("success", func(t *testing.T) {
		mockCatsRepo.On("GetByProductID", mock.Anything, mock.AnythingOfType("int64")).Return(mockCats, nil).Once()

		p := usecase.NewPCUsecase(mockCatsRepo, time.Second*2)

		cats, err := p.GetByProductID(context.TODO(), int64(8))

		assert.NoError(t, err)
		assert.Equal(t, mockCats, cats)

		mockCatsRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockCatsRepo.On("GetByProductID", mock.Anything, mock.AnythingOfType("int64")).Return(mockCats[:1], nil).Once()

		p := usecase.NewPCUsecase(mockCatsRepo, time.Second*2)

		cats, err := p.GetByProductID(context.TODO(), int64(8))

		assert.NoError(t, err)
		assert.NotEqual(t, mockCats, cats)

		mockCatsRepo.AssertExpectations(t)
	})

}

func TestGetByCategoryID(t *testing.T) {
	mockCatsRepo := new(mocks.Repository)
	mockCats := make([]*models.ProductCategory, 0)

	cat1 := models.ProductCategory{
		ID:         8,
		ProductID:  18,
		CategoryID: 30,
	}
	cat2 := models.ProductCategory{
		ID:         9,
		ProductID:  8,
		CategoryID: 30,
	}
	cat3 := models.ProductCategory{
		ID:         10,
		ProductID:  81,
		CategoryID: 30,
	}

	mockCats = append(mockCats, &cat1)
	mockCats = append(mockCats, &cat2)
	mockCats = append(mockCats, &cat3)

	t.Run("success", func(t *testing.T) {
		mockCatsRepo.On("GetByCategoryID", mock.Anything, mock.AnythingOfType("int64")).Return(mockCats, nil).Once()

		p := usecase.NewPCUsecase(mockCatsRepo, time.Second*2)

		cats, err := p.GetByCategoryID(context.TODO(), int64(8))

		assert.NoError(t, err)
		assert.Equal(t, mockCats, cats)

		mockCatsRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockCatsRepo.On("GetByCategoryID", mock.Anything, mock.AnythingOfType("int64")).Return(mockCats[:1], nil).Once()

		p := usecase.NewPCUsecase(mockCatsRepo, time.Second*2)

		cats, err := p.GetByCategoryID(context.TODO(), int64(8))

		assert.NoError(t, err)
		assert.NotEqual(t, mockCats, cats)

		mockCatsRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockCatsRepo := new(mocks.Repository)
	mockCat := models.ProductCategory{
		ProductID:  8,
		CategoryID: 10,
	}

	t.Run("success", func(t *testing.T) {
		tmpMockCat := mockCat
		tmpMockCat.ID = 0

		mockCatsRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.ProductCategory")).Return(nil).Once()

		p := usecase.NewPCUsecase(mockCatsRepo, time.Second*2)

		err := p.Create(context.TODO(), &tmpMockCat)

		assert.NoError(t, err)
		assert.Equal(t, mockCat.ID, tmpMockCat.ID)

		mockCatsRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		tmpMockCat := mockCat
		tmpMockCat.ID = 0

		mockCatsRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.ProductCategory")).Return(errors.New("Unexpected error")).Once()

		p := usecase.NewPCUsecase(mockCatsRepo, time.Second*2)

		err := p.Create(context.TODO(), &tmpMockCat)

		assert.Error(t, err)

		mockCatsRepo.AssertExpectations(t)
	})
}

func TestDeleteByProductID(t *testing.T) {
	mockCatsRepo := new(mocks.Repository)
	mockCats := make([]*models.ProductCategory, 0)

	cat1 := models.ProductCategory{
		ID:         8,
		ProductID:  8,
		CategoryID: 2,
	}
	cat2 := models.ProductCategory{
		ID:         9,
		ProductID:  8,
		CategoryID: 5,
	}
	cat3 := models.ProductCategory{
		ID:         10,
		ProductID:  8,
		CategoryID: 8,
	}

	mockCats = append(mockCats, &cat1)
	mockCats = append(mockCats, &cat2)
	mockCats = append(mockCats, &cat3)

	t.Run("success", func(t *testing.T) {
		mockCatsRepo.On("DeleteByProductID", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		p := usecase.NewPCUsecase(mockCatsRepo, time.Second*2)

		err := p.DeleteByProductID(context.TODO(), int64(8))

		assert.NoError(t, err)
		mockCatsRepo.AssertExpectations(t)
	})

	t.Run("items not found", func(t *testing.T) {
		mockCatsRepo.On("GetByProductID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil).Once()

		p := usecase.NewPCUsecase(mockCatsRepo, time.Second*2)

		result, err := p.GetByProductID(context.TODO(), int64(8))

		assert.NoError(t, err)
		assert.Nil(t, result)
		mockCatsRepo.AssertExpectations(t)
	})
}

func TestDeleteByCategoryID(t *testing.T) {
	mockCatsRepo := new(mocks.Repository)
	mockCats := make([]*models.ProductCategory, 0)

	cat1 := models.ProductCategory{
		ID:         8,
		ProductID:  18,
		CategoryID: 30,
	}
	cat2 := models.ProductCategory{
		ID:         9,
		ProductID:  8,
		CategoryID: 30,
	}
	cat3 := models.ProductCategory{
		ID:         10,
		ProductID:  81,
		CategoryID: 30,
	}

	mockCats = append(mockCats, &cat1)
	mockCats = append(mockCats, &cat2)
	mockCats = append(mockCats, &cat3)

	t.Run("success", func(t *testing.T) {
		mockCatsRepo.On("DeleteByCategoryID", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		p := usecase.NewPCUsecase(mockCatsRepo, time.Second*2)

		err := p.DeleteByCategoryID(context.TODO(), int64(8))

		assert.NoError(t, err)
		mockCatsRepo.AssertExpectations(t)
	})

	t.Run("items not found", func(t *testing.T) {
		mockCatsRepo.On("GetByCategoryID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil).Once()

		p := usecase.NewPCUsecase(mockCatsRepo, time.Second*2)

		result, err := p.GetByCategoryID(context.TODO(), int64(8))

		assert.NoError(t, err)
		assert.Nil(t, result)
		mockCatsRepo.AssertExpectations(t)
	})
}
