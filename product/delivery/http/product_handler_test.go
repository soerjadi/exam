package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/guregu/null.v3"

	categoryMocks "github.com/soerjadi/exam/category/mocks"
	"github.com/soerjadi/exam/models"
	productHttp "github.com/soerjadi/exam/product/delivery/http"
	"github.com/soerjadi/exam/product/mocks"
	catMocks "github.com/soerjadi/exam/product_category/mocks"
	price "github.com/soerjadi/exam/product_price/mocks"
	"github.com/soerjadi/exam/utils"
)

type newProduct struct {
	Name       string  `json:"name"`
	SKU        string  `json:"sku"`
	CategoryID []int64 `json:"category_id"`
}

func TestCreate(t *testing.T) {
	mockProduct := models.Product{
		Name: "product",
		SKU:  "sku",
	}

	inputProduct := newProduct{
		Name:       mockProduct.Name,
		SKU:        mockProduct.SKU,
		CategoryID: []int64{int64(8), int64(88)},
	}

	tmpMockProduct := mockProduct
	tmpMockProduct.ID = 0
	mockUsecase := new(mocks.Usecase)
	mockCatUsecase := new(catMocks.Usecase)

	mockUsecase.On("Create", mock.Anything, mock.AnythingOfType("*models.Product")).Return(nil)
	mockCatUsecase.On("Create", mock.Anything, mock.AnythingOfType("*models.ProductCategory")).Return(nil)

	j, err := json.Marshal(inputProduct)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/v1/product/add", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := productHttp.ProductHandler{
		ProductUsecase:    mockUsecase,
		ProductCatUsecase: mockCatUsecase,
	}

	handler.AddProduct(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestCreateFail(t *testing.T) {
	mockProduct := models.Product{
		Name: "product",
	}

	tmpMockProduct := mockProduct
	tmpMockProduct.ID = 0
	mockUsecase := new(mocks.Usecase)
	mockCatUsecase := new(catMocks.Usecase)

	mockUsecase.On("Create", mock.Anything, mock.AnythingOfType("*models.Product")).Return(models.ErrInternalServerError)

	j, err := json.Marshal(mockProduct)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/v1/product/add", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := productHttp.ProductHandler{
		ProductUsecase:    mockUsecase,
		ProductCatUsecase: mockCatUsecase,
	}

	handler.AddProduct(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockProduct := models.Product{
		ID:   89,
		Name: "product 89",
		SKU:  "sku89",
	}

	mockUsecase := new(mocks.Usecase)
	mockCatUsecase := new(catMocks.Usecase)
	mockCategoryUsecase := new(categoryMocks.Usecase)
	mockPriceUsecase := new(price.Usecase)

	mockUsecase.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockProduct, nil)
	mockUsecase.On("Update", mock.Anything, mock.AnythingOfType("*models.Product")).Return(nil)
	mockCatUsecase.On("DeleteByProductID", mock.Anything, mock.AnythingOfType("int64")).Return(nil)
	mockPriceUsecase.On("DeleteByProductID", mock.Anything, mock.AnythingOfType("int64")).Return(nil)

	j, err := json.Marshal(mockProduct)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/v1/product/update", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := productHttp.ProductHandler{
		ProductUsecase:    mockUsecase,
		ProductCatUsecase: mockCatUsecase,
		CategoryUsecase:   mockCategoryUsecase,
		PriceUsecase:      mockPriceUsecase,
	}

	handler.UpdateProduct(rec, req)

	_expectedResponse := &utils.DefaultResponse{
		Code:    200,
		Message: "success",
		Result:  mockProduct,
	}

	expectedResponse, err := json.Marshal(_expectedResponse)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, string(expectedResponse), rec.Body.String())
	mockUsecase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockProduct models.Product
	err := faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	mockUsecase := new(mocks.Usecase)
	mockCatUsecase := new(catMocks.Usecase)
	mockCategoryUsecase := new(categoryMocks.Usecase)

	id := mockProduct.ID

	cats := make([]*models.ProductCategory, 0)
	cats = append(cats, &models.ProductCategory{
		ID:         1,
		ProductID:  id,
		CategoryID: int64(8),
	})
	category := models.Category{
		ID:       8,
		Name:     "category 8",
		ParentID: null.NewInt(int64(0), true),
	}

	mockUsecase.On("GetByID", mock.Anything, id).Return(&mockProduct, nil)
	mockCatUsecase.On("GetByProductID", mock.Anything, mock.AnythingOfType("int64")).Return(cats, nil)
	mockCategoryUsecase.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&category, nil)

	req, err := http.NewRequest("GET", "/v1/product/detail?id="+strconv.FormatInt(id, 10), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := productHttp.ProductHandler{
		ProductUsecase:    mockUsecase,
		ProductCatUsecase: mockCatUsecase,
		CategoryUsecase:   mockCategoryUsecase,
	}

	handler.GetByID(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestCompareProduct(t *testing.T) {
	var mockProduct1, mockProduct2 models.Product
	err := faker.FakeData(&mockProduct1)
	assert.NoError(t, err)

	err = faker.FakeData(&mockProduct2)
	assert.NoError(t, err)

	mockUsecase := new(mocks.Usecase)
	mockCatUsecase := new(catMocks.Usecase)
	compareProduct := []*models.Product{
		&mockProduct1, &mockProduct2,
	}

	mockUsecase.On("Compare", mock.Anything, mockProduct1.ID, mockProduct2.ID).Return(compareProduct, nil)

	req, err := http.NewRequest("GET", "/v1/product/compare", strings.NewReader(""))
	req = mux.SetURLVars(req, map[string]string{"id_1": strconv.FormatInt(mockProduct1.ID, 10), "id_2": strconv.FormatInt(mockProduct2.ID, 10)})

	handler := productHttp.ProductHandler{
		ProductUsecase:    mockUsecase,
		ProductCatUsecase: mockCatUsecase,
	}

	rec := httptest.NewRecorder()
	handler.CompareProduct(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	// assert.Equal(t, "", rec.Body)

	mockUsecase.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	var mockProduct models.Product
	err := faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	mockUsecase := new(mocks.Usecase)
	mockCatUsecase := new(catMocks.Usecase)

	mockUsecase.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil)
	mockCatUsecase.On("DeleteByProductID", mock.Anything, mock.AnythingOfType("int64")).Return(nil)

	req, err := http.NewRequest("GET", "/v1/product/delete?id="+strconv.FormatInt(mockProduct.ID, 10), strings.NewReader(""))

	handler := productHttp.ProductHandler{
		ProductUsecase:    mockUsecase,
		ProductCatUsecase: mockCatUsecase,
	}

	rec := httptest.NewRecorder()
	handler.DeleteProduct(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}
