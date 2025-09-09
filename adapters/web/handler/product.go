package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/leomoritz/fullcycle-arq-hexagonal-app/adapters/web/dto"
	"github.com/leomoritz/fullcycle-arq-hexagonal-app/application"
)

func MakeProductHandlers(router *mux.Router, middleware *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", middleware.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	router.Handle("/product", middleware.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	router.Handle("/product/{id}/enable", middleware.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PUT", "OPTIONS")

	router.Handle("/product/{id}/disable", middleware.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PUT", "OPTIONS")
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

		err = json.NewEncoder(response).Encode(product)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")

		var productDto = dto.NewProductDto()
		err := json.NewDecoder(request.Body).Decode(&productDto)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(JsonError(err.Error()))
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write(JsonError(err.Error()))
			return
		}

		err = json.NewEncoder(response).Encode(product)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(JsonError(err.Error()))
			return
		}
		response.WriteHeader(http.StatusOK)
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
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

		result, err := service.Enable(product)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write(JsonError(err.Error()))
			return
		}

		err = json.NewEncoder(response).Encode(result)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
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

		result, err := service.Disable(product)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write(JsonError(err.Error()))
			return
		}

		err = json.NewEncoder(response).Encode(result)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
