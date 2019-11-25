package main

import (
	"time"

	"github.com/gorilla/mux"
	"github.com/soerjadi/exam/database"
	"github.com/soerjadi/exam/utils"

	pHttp "github.com/soerjadi/exam/product/delivery/http"
	pRepo "github.com/soerjadi/exam/product/repository"
	pUsecase "github.com/soerjadi/exam/product/usecase"

	cHttp "github.com/soerjadi/exam/category/delivery/http"
	cRepo "github.com/soerjadi/exam/category/repository"
	cUsecase "github.com/soerjadi/exam/category/usecase"
)

// RegisterRouter --
func RegisterRouter(router *mux.Router) *mux.Router {
	// router.HandleFunc("/v1/info", HelloWorld).Methods("GET")

	conn := database.RDB().DB()
	timeout := time.Duration(utils.GetEnvInt("CONTEXT_TIMEOUT", 0)) * time.Second

	productRepo := pRepo.NewPGProductRepository(conn)
	productUsecase := pUsecase.NewProductUsecase(productRepo, timeout)
	pHttp.NewProductHandler(router, productUsecase)

	categoryRepo := cRepo.NewPGCategoryRepository(conn)
	categoryUsecase := cUsecase.NewCategoryUsecase(categoryRepo, timeout)
	cHttp.NewCategoryHandler(router, categoryUsecase)

	return router
}
