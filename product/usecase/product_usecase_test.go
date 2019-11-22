package usecase_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/product/mocks"
	"github.com/soerjadi/exam/product/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSearch(t *testing.T) {
	mockProductRepo := new(mocks.Repository)
	mockProduct := &models.Product{
		Name:    "product 1",
		SKU:     "1234",
		Created: time.Now(),
	}

	mockListProducts := make([]*models.Product, 0)
	mockListProducts = append(mockListProducts, mockProduct)
	searchQuery := strings.ToLower("product")

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Search", mock.Anything, &searchQuery,
			mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).
			Return(mockListProducts, int64(1), nil).Once()

		p := usecase.NewProductUsecase(mockProductRepo, time.Second*2)
		products, found, err := p.Search(context.TODO(), &searchQuery, int64(0), int64(10))

		assert.NoError(t, err)
		assert.Equal(t, int64(1), found)
		assert.Len(t, products, len(mockListProducts))

		mockProductRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockProductRepo.On("Search", mock.Anything, &searchQuery,
			mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).
			Return(nil, int64(0), errors.New("Unexpected Error")).Once()

		p := usecase.NewProductUsecase(mockProductRepo, time.Second*2)
		products, found, err := p.Search(context.TODO(), &searchQuery, int64(0), int64(10))

		assert.Error(t, err)
		assert.Len(t, products, 0)
		assert.Equal(t, int64(0), found)

		mockProductRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	mockProductRepo := new(mocks.Repository)
	mockProduct := &models.Product{
		ID:      64,
		Name:    "product 64",
		SKU:     "sku64",
		Created: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockProduct, nil).Once()

		p := usecase.NewProductUsecase(mockProductRepo, time.Second*2)

		product, err := p.GetByID(context.TODO(), mockProduct.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockProduct, product)

		mockProductRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected errors")).Once()

		p := usecase.NewProductUsecase(mockProductRepo, time.Second*2)

		product, err := p.GetByID(context.TODO(), mockProduct.ID)

		assert.Error(t, err)
		assert.Nil(t, product)

		mockProductRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockProductRepo := new(mocks.Repository)
	mockProduct := models.Product{
		Name: "product 64",
		SKU:  "sku64",
	}

	t.Run("success", func(t *testing.T) {
		tmpMockProduct := mockProduct
		tmpMockProduct.ID = 0
		mockProductRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Product")).Return(nil).Once()

		p := usecase.NewProductUsecase(mockProductRepo, time.Second*2)

		err := p.Create(context.TODO(), &tmpMockProduct)

		assert.NoError(t, err)
		assert.Equal(t, mockProduct.Name, tmpMockProduct.Name)

		mockProductRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockProductRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Product")).Return(errors.New("Unexpected error")).Once()

		p := usecase.NewProductUsecase(mockProductRepo, time.Second*2)

		err := p.Create(context.TODO(), &mockProduct)

		assert.Error(t, err)

		mockProductRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockProductRepo := new(mocks.Repository)
	mockProduct := models.Product{
		ID:      64,
		Name:    "product 64",
		SKU:     "sku64",
		Created: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Update", mock.Anything, &mockProduct).Return(nil).Once()

		p := usecase.NewProductUsecase(mockProductRepo, time.Second*2)

		err := p.Update(context.TODO(), &mockProduct)

		assert.NoError(t, err)
		mockProductRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockProductRepo := new(mocks.Repository)
	mockProduct := models.Product{
		Name:    "product 64",
		SKU:     "sku64",
		Created: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockProduct, nil).Once()
		mockProductRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		p := usecase.NewProductUsecase(mockProductRepo, time.Second*2)

		err := p.Delete(context.TODO(), mockProduct.ID)

		assert.NoError(t, err)
		mockProductRepo.AssertExpectations(t)
	})
	t.Run("item is not exist", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil).Once()

		p := usecase.NewProductUsecase(mockProductRepo, time.Second*2)

		err := p.Delete(context.TODO(), mockProduct.ID)

		assert.Error(t, err)
		mockProductRepo.AssertExpectations(t)
	})
}

func TestCompare(t *testing.T) {
	mockProductRepo := new(mocks.Repository)
	mockProduct1 := models.Product{
		ID:      64,
		Name:    "product 64",
		SKU:     "sku64",
		Created: time.Now(),
	}
	mockProduct2 := models.Product{
		ID:      129,
		Name:    "product 129",
		SKU:     "sku129",
		Created: time.Now(),
	}

	compareProduct := []*models.Product{
		&mockProduct1, &mockProduct2,
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockProduct1, nil).Once()
		mockProductRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockProduct2, nil).Once()

		p := usecase.NewProductUsecase(mockProductRepo, time.Second*2)

		products, err := p.Compare(context.TODO(), mockProduct1.ID, mockProduct2.ID)

		assert.NoError(t, err)
		assert.Len(t, products, 2)
		assert.Equal(t, compareProduct, products)

		mockProductRepo.AssertExpectations(t)
	})
}
