package handler

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/leomoritz/fullcycle-arq-hexagonal-app/application"
)

func MakeProductHandlers(router *mux.Router, middleware *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", middleware.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(request) // pega todas as variáveis da requisição
		id := vars["id"]

		product, err := service.Get(id)

		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			return
		}

		if product == nil {
			response.WriteHeader(http.StatusNotFound)
			return
		}
	})
}
