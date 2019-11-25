package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/soerjadi/exam/category"
	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/utils"
	"gopkg.in/guregu/null.v3"
)

type newCategory struct {
	Name     string   `json:"name"`
	ParentID null.Int `json:"parent_id"`
}

type updateCategoryData struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	ParentID null.Int `json:"parent_id"`
}

// CategoryHandler represent the http handler for category
type CategoryHandler struct {
	CategoryUsecase category.Usecase
}

var logger = utils.LogBuilder(true)

// NewCategoryHandler initialize category resource endpoint
func NewCategoryHandler(router *mux.Router, usecase category.Usecase) *mux.Router {
	handler := &CategoryHandler{
		CategoryUsecase: usecase,
	}

	c := router.PathPrefix("/v1/category").Subrouter()
	c.HandleFunc("/add", handler.AddCategory).Methods("POST")
	c.HandleFunc("/update", handler.UpdateCategory).Methods("POST")
	c.HandleFunc("/detail", handler.GetByID).Methods("GET")
	c.HandleFunc("/delete", handler.Delete).Methods("GET")

	return c
}

// AddCategory will add category from given body to DB
func (h *CategoryHandler) AddCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCategory newCategory
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &newCategory)
	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	category := models.Category{
		Name:     newCategory.Name,
		ParentID: newCategory.ParentID,
	}

	err = h.CategoryUsecase.Create(ctx, &category)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, category)
}

// UpdateCategory will update category from specific id
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateCategory updateCategoryData
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &updateCategory)
	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	category := models.Category{
		ID:       updateCategory.ID,
		Name:     updateCategory.Name,
		ParentID: updateCategory.ParentID,
	}

	err = h.CategoryUsecase.Update(ctx, &category)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, category)
}

// GetByID get detail category from given ID
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, err := strconv.ParseInt(params.Get("id"), 0, 64)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	category, err := h.CategoryUsecase.GetByID(ctx, id)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, category)
}

// Delete will delete category by given ID
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, err := strconv.ParseInt(params.Get("id"), 0, 64)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = h.CategoryUsecase.Delete(ctx, id)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, "success")
}
