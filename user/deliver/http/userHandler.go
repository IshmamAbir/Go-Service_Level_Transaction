package http

import (
	"encoding/json"
	"net/http"

	CommonError "main.go/errors"
	"main.go/response"

	"github.com/gorilla/mux"
	"main.go/requests"
	"main.go/user/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase, router *mux.Router) {
	handler := &UserHandler{
		userUsecase: userUsecase,
	}
	subroute := router.PathPrefix("/users").Subrouter()
	subroute.HandleFunc("", handler.GetAll).Methods("GET")
	subroute.HandleFunc("/order-product", handler.PurchaseProduct).Methods("POST")
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.FindAll(r.Context())
	if err != nil {
		response.Render(w, CommonError.ErrInternalServerError, nil)
		return
	}
	response.Render(w, nil, users)
}

func (h *UserHandler) PurchaseProduct(w http.ResponseWriter, r *http.Request) {
	orderRequest := requests.OrderRequest{}
	json.NewDecoder(r.Body).Decode(&orderRequest)
	if err := h.userUsecase.PurchaseProduct(r.Context(), orderRequest); err != nil {
		response.Render(w, err, nil)
		return
	}
	response.Render(w, nil, "order taken successfully !")
}
