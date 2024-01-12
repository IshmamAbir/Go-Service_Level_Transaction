package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"main.go/response"
	"main.go/shopping_cart/usecase"
)

type ShoppingCartHandler struct {
	ShoppingCartUsecase usecase.ShoppingCartUsecase
}

func NewShoppingCartHandler(shoppingCartUsecase usecase.ShoppingCartUsecase, router *mux.Router) {
	handler := &ShoppingCartHandler{
		ShoppingCartUsecase: shoppingCartUsecase,
	}
	subroute := router.PathPrefix("/shopping-cart").Subrouter()
	subroute.HandleFunc("", handler.GetAll).Methods("GET")
}

func (h *ShoppingCartHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	shoppingCarts, err := h.ShoppingCartUsecase.FindAll(r.Context())
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		response.Render(w, err, nil)
		return
	}
	// json.NewEncoder(w).Encode(shoppingCarts)
	response.Render(w, nil, shoppingCarts)
}
