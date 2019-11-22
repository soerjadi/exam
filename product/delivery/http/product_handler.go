package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/product"
	"github.com/soerjadi/exam/utils"
)

type newProduct struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

type updateProductData struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

// ProductHandler represent the http handler for product
type ProductHandler struct {
	ProductUsecase product.Usecase
}

var logger = utils.LogBuilder(true)

// NewProductHandler initialize product resource endpoint
func NewProductHandler(router *mux.Router, usecase product.Usecase) *mux.Router {
	handler := &ProductHandler{
		ProductUsecase: usecase,
	}

	p := router.PathPrefix("/v1/product").Subrouter()
	p.HandleFunc("/add", handler.AddProduct).Methods("POST")
	p.HandleFunc("/update", handler.UpdateProduct).Methods("POST")
	p.HandleFunc("/detail", handler.GetByID).Methods("GET")
	p.HandleFunc("/compare/{ID1}/{ID2}", handler.CompareProduct).Methods("GET")
	p.HandleFunc("/search", handler.SearchProduct).Methods("GET")
	return p
}

// AddProduct will add product from given body to DB
func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newProduct newProduct
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &newProduct)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	product := models.Product{
		Name: newProduct.Name,
		SKU:  newProduct.SKU,
	}
	err = h.ProductUsecase.Create(ctx, &product)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, product)

}

// UpdateProduct will update product from given body to DB
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateProduct updateProductData
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &updateProduct)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	product := models.Product{
		ID:   updateProduct.ID,
		Name: updateProduct.Name,
		SKU:  updateProduct.SKU,
	}

	err = h.ProductUsecase.Update(ctx, &product)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, product)
}

// GetByID get detail product from given ID
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	product, err := h.ProductUsecase.GetByID(ctx, id)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, product)
}

// CompareProduct detail product
func (h *ProductHandler) CompareProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	productID1, err := strconv.ParseInt(vars["id_1"], 0, 64)
	if err != nil {
		logger.Error(err)
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	productID2, err := strconv.ParseInt(vars["id_2"], 0, 64)
	if err != nil {
		logger.Error(err)
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	products, err := h.ProductUsecase.Compare(ctx, productID1, productID2)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, products)
}

// SearchProduct search product from specific query
func (h *ProductHandler) SearchProduct(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	query := params.Get("query")
	limit, err := strconv.ParseInt(params.Get("limit"), 0, 64)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	offset, err := strconv.ParseInt(params.Get("offset"), 0, 64)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if offset == 0 {
		offset = 10
	}

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	products, found, err := h.ProductUsecase.Search(ctx, &query, offset, limit)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	entriesResult := &utils.EntriesResponse{
		Data:  products,
		Found: found,
	}

	utils.JSON(w, http.StatusOK, entriesResult)
}

// DeleteProduct will delete product by given ID
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
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

	err = h.ProductUsecase.Delete(ctx, id)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, "success")
}
