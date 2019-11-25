package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/soerjadi/exam/models"
	orderHttp "github.com/soerjadi/exam/order/delivery/http"
	"github.com/soerjadi/exam/order/mocks"
	pMocks "github.com/soerjadi/exam/product/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type newOrder struct {
	ProductID int64   `json:"product_id"`
	Amount    int64   `json:"amount"`
	Price     float64 `json:"price"`
	Status    int     `json:"status"`
}

func TestCreate(t *testing.T) {
	mockOrder := models.Order{
		ProductID: int64(9),
		Amount:    int64(10),
		Price:     10000.0,
		Status:    models.OrderPending,
	}

	mockProduct := models.Product{
		ID:   int64(9),
		Name: "product 9",
		SKU:  "sku9",
	}

	inputOrder := newOrder{
		ProductID: mockOrder.ProductID,
		Amount:    mockOrder.Amount,
		Price:     mockOrder.Price,
		Status:    mockOrder.Status,
	}

	tmpMockOrder := mockOrder
	tmpMockOrder.ID = 0

	mockUsecase := new(mocks.Usecase)
	mockProductUsecase := new(pMocks.Usecase)

	mockUsecase.On("Create", mock.Anything, mock.AnythingOfType("*models.Order")).Return(nil)
	mockProductUsecase.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockProduct, nil)

	j, err := json.Marshal(inputOrder)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/v1/order/add", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := orderHttp.OrderHandler{
		OrderUsecase:   mockUsecase,
		ProductUsecase: mockProductUsecase,
	}

	handler.CreateOrder(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	var mockOrder models.Order
	err := faker.FakeData(&mockOrder)
	assert.NoError(t, err)

	mockUsecase := new(mocks.Usecase)

	mockUsecase.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil)

	req, err := http.NewRequest("GET", "/v1/order/delete?id="+strconv.FormatInt(mockOrder.ID, 10), strings.NewReader(""))

	handler := orderHttp.OrderHandler{
		OrderUsecase: mockUsecase,
	}

	rec := httptest.NewRecorder()
	handler.Delete(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}
