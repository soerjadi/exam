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
	"github.com/soerjadi/exam/product"
	cat "github.com/soerjadi/exam/product_category"
	price "github.com/soerjadi/exam/product_price"
	t "github.com/soerjadi/exam/types"
	"github.com/soerjadi/exam/utils"
)

type productPrice struct {
	Amount int64   `json:"amount"`
	Price  float64 `json:"price"`
}

type newProduct struct {
	Name       string         `json:"name"`
	SKU        string         `json:"sku"`
	CategoryID []int64        `json:"category_id"`
	Price      []productPrice `json:"price"`
}

type updateProductData struct {
	ID         int64          `json:"id"`
	Name       string         `json:"name"`
	SKU        string         `json:"sku"`
	CategoryID []int64        `json:"category_id"`
	Price      []productPrice `json:"price"`
}

// ProductHandler represent the http handler for product
type ProductHandler struct {
	ProductUsecase    product.Usecase
	ProductCatUsecase cat.Usecase
	CategoryUsecase   category.Usecase
	PriceUsecase      price.Usecase
}

var logger = utils.LogBuilder(true)

// NewProductHandler initialize product resource endpoint
func NewProductHandler(router *mux.Router, usecase product.Usecase, catUsecase cat.Usecase, categoryUsecase category.Usecase) *mux.Router {
	handler := &ProductHandler{
		ProductUsecase:    usecase,
		ProductCatUsecase: catUsecase,
		CategoryUsecase:   categoryUsecase,
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

	for _, catID := range newProduct.CategoryID {
		cat := models.ProductCategory{
			ProductID:  product.ID,
			CategoryID: catID,
		}

		err = h.ProductCatUsecase.Create(ctx, &cat)

		if err != nil {
			utils.Error(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	for _, price := range newProduct.Price {
		_price := models.ProductPrice{
			Amount:    price.Amount,
			Price:     price.Price,
			ProductID: product.ID,
		}

		err = h.PriceUsecase.Create(ctx, &_price)

		if err != nil {
			utils.Error(w, http.StatusBadRequest, err.Error())
			return
		}
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

	origProduct, err := h.ProductUsecase.GetByID(ctx, updateProduct.ID)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, models.ErrNotFound.Error())
		return
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

	err = h.ProductCatUsecase.DeleteByProductID(ctx, product.ID)

	// Revert when get an error update link product category
	if err != nil {
		_ = h.ProductUsecase.Update(ctx, origProduct)
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.PriceUsecase.DeleteByProductID(ctx, product.ID)
	// Revert when get an error update price
	if err != nil {
		_ = h.ProductUsecase.Update(ctx, origProduct)
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	for _, cat := range updateProduct.CategoryID {
		cat := &models.ProductCategory{
			ProductID:  updateProduct.ID,
			CategoryID: cat,
		}

		err = h.ProductCatUsecase.Create(ctx, cat)
		if err != nil {
			utils.Error(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	for _, price := range updateProduct.Price {
		_price := models.ProductPrice{
			Amount:    price.Amount,
			Price:     price.Price,
			ProductID: product.ID,
		}

		err = h.PriceUsecase.Create(ctx, &_price)

		if err != nil {
			utils.Error(w, http.StatusBadRequest, err.Error())
			return
		}
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
	cats := make([]*models.Category, 0)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	pCats, err := h.ProductCatUsecase.GetByProductID(ctx, product.ID)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	for _, pCat := range pCats {
		cat, err := h.CategoryUsecase.GetByID(ctx, pCat.CategoryID)
		if err != nil {
			continue
		}

		cats = append(cats, cat)
	}

	result := t.Product{
		ID:       product.ID,
		Name:     product.Name,
		SKU:      product.SKU,
		Category: cats,
	}

	utils.JSON(w, http.StatusOK, result)
}

// CompareProduct detail product
func (h *ProductHandler) CompareProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
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

	// escape error because there is no connection between product and category already
	_ = h.ProductCatUsecase.DeleteByProductID(ctx, id)

	utils.JSON(w, http.StatusOK, "success")
}
