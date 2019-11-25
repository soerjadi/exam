package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	categoryHttp "github.com/soerjadi/exam/category/delivery/http"
	"github.com/soerjadi/exam/category/mocks"
	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/guregu/null.v3"
)

func TestCreate(t *testing.T) {
	mockCategory := models.Category{
		Name:     "category 1",
		ParentID: null.NewInt(int64(0), true),
	}

	tmpMockCategory := mockCategory
	tmpMockCategory.ID = 0
	mockUsecase := new(mocks.Usecase)

	mockUsecase.On("Create", mock.Anything, mock.AnythingOfType("*models.Category")).Return(nil)

	j, err := json.Marshal(mockCategory)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/v1/category/add", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := categoryHttp.CategoryHandler{
		CategoryUsecase: mockUsecase,
	}

	handler.AddCategory(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestCreateFail(t *testing.T) {
	mockCategory := models.Category{
		Name:     "category 1",
		ParentID: null.NewInt(int64(0), true),
	}

	tmpMockCategory := mockCategory
	tmpMockCategory.ID = 0
	mockUsecase := new(mocks.Usecase)

	mockUsecase.On("Create", mock.Anything, mock.AnythingOfType("*models.Category")).Return(models.ErrInternalServerError)

	j, err := json.Marshal(mockCategory)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/v1/category/add", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := categoryHttp.CategoryHandler{
		CategoryUsecase: mockUsecase,
	}

	handler.AddCategory(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockCategory models.Category
	err := faker.FakeData(&mockCategory)
	assert.NoError(t, err)

	mockUsecase := new(mocks.Usecase)

	id := mockCategory.ID

	mockUsecase.On("GetByID", mock.Anything, id).Return(&mockCategory, nil)

	req, err := http.NewRequest("GET", "/v1/category/detail?id="+strconv.FormatInt(id, 10), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := categoryHttp.CategoryHandler{
		CategoryUsecase: mockUsecase,
	}

	handler.GetByID(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockCategory := models.Category{
		ID:       56,
		Name:     "category 56",
		ParentID: null.NewInt(int64(0), true),
	}

	mockUsecase := new(mocks.Usecase)

	mockUsecase.On("Update", mock.Anything, mock.AnythingOfType("*models.Category")).Return(nil)

	j, err := json.Marshal(mockCategory)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/v1/category/update", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := categoryHttp.CategoryHandler{
		CategoryUsecase: mockUsecase,
	}

	handler.UpdateCategory(rec, req)

	_expectedResponse := &utils.DefaultResponse{
		Code:    200,
		Message: "success",
		Result:  mockCategory,
	}

	expectedResponse, err := json.Marshal(_expectedResponse)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, string(expectedResponse), rec.Body.String())
	mockUsecase.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	var mockCategory models.Category
	err := faker.FakeData(&mockCategory)
	assert.NoError(t, err)

	mockUsecase := new(mocks.Usecase)

	categoryID := mockCategory.ID

	mockUsecase.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil)

	req, err := http.NewRequest("GET", "/v1/category/delete?id="+strconv.FormatInt(categoryID, 10), strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := categoryHttp.CategoryHandler{
		CategoryUsecase: mockUsecase,
	}

	handler.Delete(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}
