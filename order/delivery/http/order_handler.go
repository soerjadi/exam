package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/order"
	"github.com/soerjadi/exam/product"
	price "github.com/soerjadi/exam/product_price"
	"github.com/soerjadi/exam/utils"
)

type newOrder struct {
	ProductID int64   `json:"product_id"`
	Amount    int64   `json:"amount"`
	Price     float64 `json:"price"`
	Status    int     `json:"status"`
}

// OrderHandler represent the http handler for order
type OrderHandler struct {
	OrderUsecase   order.Usecase
	ProductUsecase product.Usecase
	PriceUsecase   price.Usecase
}

var logger = utils.LogBuilder(true)

// NewOrderHandler initialize product resource endpoint
func NewOrderHandler(router *mux.Router, usecase order.Usecase, productUsecase product.Usecase) *mux.Router {
	handler := &OrderHandler{
		OrderUsecase:   usecase,
		ProductUsecase: productUsecase,
	}

	p := router.PathPrefix("/v1/order").Subrouter()
	p.HandleFunc("/add", handler.CreateOrder).Methods("POST")
	p.HandleFunc("/list", handler.GetList).Methods("GET")
	p.HandleFunc("/delete", handler.Delete).Methods("GET")

	return p
}

// CreateOrder endpoint for create an order
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newOrder newOrder
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &newOrder)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	product, err := h.ProductUsecase.GetByID(ctx, newOrder.ProductID)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if product == nil {
		utils.Error(w, http.StatusBadRequest, models.ErrNotFound.Error())
		return
	}

	order := models.Order{
		ProductID: newOrder.ProductID,
		Price:     newOrder.Price,
		Amount:    newOrder.Amount,
		Status:    newOrder.Status,
	}

	err = h.OrderUsecase.Create(ctx, &order)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, order)

}

// GetList endpoint for get list an order
func (h *OrderHandler) GetList(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
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

	if limit == 0 {
		limit = 10
	}

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	orders, found, err := h.OrderUsecase.GetList(ctx, offset, limit)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	entriesResult := &utils.EntriesResponse{
		Data:  orders,
		Found: found,
	}

	utils.JSON(w, http.StatusOK, entriesResult)

}

// Delete endpoint to delete an order
func (h *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	err = h.OrderUsecase.Delete(ctx, id)

	utils.JSON(w, http.StatusOK, "success")
}
