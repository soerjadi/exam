package usecase_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/soerjadi/exam/category/mocks"
	"github.com/soerjadi/exam/category/usecase"
	"github.com/soerjadi/exam/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/guregu/null.v3"
)

func TestSearch(t *testing.T) {
	mockCategoryRepo := new(mocks.Repository)
	mockCategory := &models.Category{
		Name:    "category 1",
		Created: time.Now(),
	}

	mockListCategory := make([]*models.Category, 0)
	mockListCategory = append(mockListCategory, mockCategory)
	searchQuery := strings.ToLower("category")

	t.Run("success", func(t *testing.T) {
		mockCategoryRepo.On("Search", mock.Anything, &searchQuery,
			mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).
			Return(mockListCategory, int64(1), nil).Once()

		p := usecase.NewCategoryUsecase(mockCategoryRepo, time.Second*2)
		categories, found, err := p.Search(context.TODO(), &searchQuery, int64(0), int64(10))

		assert.NoError(t, err)
		assert.Equal(t, int64(1), found)
		assert.Len(t, categories, len(mockListCategory))

		mockCategoryRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockCategoryRepo.On("Search", mock.Anything, &searchQuery,
			mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).
			Return(nil, int64(0), errors.New("Unexpected Error")).Once()

		p := usecase.NewCategoryUsecase(mockCategoryRepo, time.Second*2)
		categories, found, err := p.Search(context.TODO(), &searchQuery, int64(0), int64(10))

		assert.Error(t, err)
		assert.Len(t, categories, 0)
		assert.Equal(t, int64(0), found)

		mockCategoryRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	mockCategoryRepo := new(mocks.Repository)
	mockCategory := &models.Category{
		ID:   64,
		Name: "category 64",
		ParentID: null.NewInt(
			int64(0), true,
		),
		Created: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockCategory, nil).Once()

		p := usecase.NewCategoryUsecase(mockCategoryRepo, time.Second*2)

		product, err := p.GetByID(context.TODO(), mockCategory.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockCategory, product)

		mockCategoryRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected errors")).Once()

		p := usecase.NewCategoryUsecase(mockCategoryRepo, time.Second*2)

		product, err := p.GetByID(context.TODO(), mockCategory.ID)

		assert.Error(t, err)
		assert.Nil(t, product)

		mockCategoryRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockCategoryRepo := new(mocks.Repository)
	mockCategory := models.Category{
		Name: "category 1",
		ParentID: null.NewInt(
			int64(0), true,
		),
		Created: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		tmpMockProduct := mockCategory
		tmpMockProduct.ID = 0
		mockCategoryRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Category")).Return(nil).Once()

		p := usecase.NewCategoryUsecase(mockCategoryRepo, time.Second*2)

		err := p.Create(context.TODO(), &tmpMockProduct)

		assert.NoError(t, err)
		assert.Equal(t, mockCategory.Name, tmpMockProduct.Name)

		mockCategoryRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockCategoryRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Category")).Return(errors.New("Unexpected error")).Once()

		p := usecase.NewCategoryUsecase(mockCategoryRepo, time.Second*2)

		err := p.Create(context.TODO(), &mockCategory)

		assert.Error(t, err)

		mockCategoryRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockCategoryRepo := new(mocks.Repository)
	mockCategory := models.Category{
		ID:   64,
		Name: "category 64",
		ParentID: null.NewInt(
			int64(0), true,
		),
		Created: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockCategoryRepo.On("Update", mock.Anything, &mockCategory).Return(nil).Once()

		p := usecase.NewCategoryUsecase(mockCategoryRepo, time.Second*2)

		err := p.Update(context.TODO(), &mockCategory)

		assert.NoError(t, err)
		mockCategoryRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockCategoryRepo := new(mocks.Repository)
	mockCategory := models.Category{
		Name: "category 1",
		ParentID: null.NewInt(
			int64(0), true,
		),
		Created: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockCategory, nil).Once()
		mockCategoryRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		p := usecase.NewCategoryUsecase(mockCategoryRepo, time.Second*2)

		err := p.Delete(context.TODO(), mockCategory.ID)

		assert.NoError(t, err)
		mockCategoryRepo.AssertExpectations(t)
	})
	t.Run("item is not exist", func(t *testing.T) {
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil).Once()

		p := usecase.NewCategoryUsecase(mockCategoryRepo, time.Second*2)

		err := p.Delete(context.TODO(), mockCategory.ID)

		assert.Error(t, err)
		mockCategoryRepo.AssertExpectations(t)
	})
}
