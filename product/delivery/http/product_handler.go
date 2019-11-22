package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/product"
	"github.com/soerjadi/exam/utils"
)

type newProduct struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

// ProductHandler represent the http handler for product
type ProductHandler struct {
	ProductUsecase product.Usecase
}

// NewProductHandler initialize product resource endpoint
func NewProductHandler(router *mux.Router, usecase product.Usecase) *mux.Router {
	handler := &ProductHandler{
		ProductUsecase: usecase,
	}

	p := router.PathPrefix("/v1/product").Subrouter()
	p.HandleFunc("/add", handler.AddProduct).Methods("POST")

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
