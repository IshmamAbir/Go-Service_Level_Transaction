package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/shopping_cart/usecase"
)

type ShoppingCartHandler struct {
	ShoppingCartUsecase usecase.ShoppingCartUsecase
}

func NewShoppingCartHandler(shoppingCartUsecase usecase.ShoppingCartUsecase, router *mux.Router) {
	handler := &ShoppingCartHandler{
		ShoppingCartUsecase: shoppingCartUsecase,
	}
	subroute := router.PathPrefix("/shopping_cart").Subrouter()
	subroute.HandleFunc("/", handler.GetAll).Methods("GET")
}

func (h *ShoppingCartHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	shoppingCarts, err := h.ShoppingCartUsecase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(shoppingCarts)
}
