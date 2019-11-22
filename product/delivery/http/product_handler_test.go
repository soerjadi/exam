package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/soerjadi/exam/models"
	productHttp "github.com/soerjadi/exam/product/delivery/http"
	"github.com/soerjadi/exam/product/mocks"
)

func TestCreate(t *testing.T) {
	mockProduct := models.Product{
		Name: "product",
		SKU:  "sku",
	}

	tmpMockProduct := mockProduct
	tmpMockProduct.ID = 0
	mockUsecase := new(mocks.Usecase)

	mockUsecase.On("Create", mock.Anything, mock.AnythingOfType("*models.Product")).Return(nil)

	j, err := json.Marshal(mockProduct)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/v1/product/add", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := productHttp.ProductHandler{
		ProductUsecase: mockUsecase,
	}

	handler.AddProduct(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}
