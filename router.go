package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soerjadi/exam/utils"
)

// RegisterRouter --
func RegisterRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/info", HelloWorld).Methods("GET")

	return router
}

// HelloWorld --
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, "success")
}
