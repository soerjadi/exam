package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/soerjadi/exam/database"
	"github.com/soerjadi/exam/utils"

	pHttp "github.com/soerjadi/exam/product/delivery/http"
	pRepo "github.com/soerjadi/exam/product/repository"
	pUsecase "github.com/soerjadi/exam/product/usecase"
)

// RegisterRouter --
func RegisterRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/info", HelloWorld).Methods("GET")

	conn := database.RDB().DB()
	timeout := time.Duration(utils.GetEnvInt("CONTEXT_TIMEOUT", 0)) * time.Second

	productRepo := pRepo.NewPGProductRepository(conn)
	productUsecase := pUsecase.NewProductUsecase(productRepo, timeout)
	pHttp.NewProductHandler(router, productUsecase)

	return router
}

// HelloWorld --
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, "success")
}
