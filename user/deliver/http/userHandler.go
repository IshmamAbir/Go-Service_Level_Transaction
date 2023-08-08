package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"main.go/requests"
	"main.go/transaction"
	"main.go/user/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
	Db          *gorm.DB
}

func NewUserHandler(userUsecase usecase.UserUsecase, router *mux.Router, db *gorm.DB) {
	handler := &UserHandler{
		userUsecase: userUsecase,
		Db:          db,
	}
	subroute := router.PathPrefix("/users").Subrouter()
	subroute.HandleFunc("/", handler.GetAll).Methods("GET")
	subroute.HandleFunc("/order_product", transaction.DBTransactionMiddleware(handler.Db, handler.OrderProduct)).Methods("POST")
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) OrderProduct(w http.ResponseWriter, r *http.Request) {
	println("order product")

	orderRequest := requests.OrderRequest{}
	json.NewDecoder(r.Body).Decode(&orderRequest)
	err := h.userUsecase.OrderProduct(orderRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("order successful !")
}
