package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/product/usecase"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase, router *mux.Router) {
	handler := &ProductHandler{
		ProductUsecase: productUsecase,
	}
	subroute := router.PathPrefix("/products").Subrouter()
	subroute.HandleFunc("/", handler.GetAll).Methods("GET")
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductUsecase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}
