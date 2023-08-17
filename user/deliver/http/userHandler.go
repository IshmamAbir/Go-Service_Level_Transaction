package http

import (
	"encoding/json"
	"net/http"

	CommonError "main.go/errors"
	"main.go/response"

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
	subroute.HandleFunc("/order-product", transaction.DBTransactionMiddleware(handler.Db, handler.PurchaseProduct)).Methods("POST")
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.FindAll()
	if err != nil {
		response.Render(w, CommonError.ErrInternalServerError, nil)
		return
	}
	response.Render(w, nil, users)
}

func (h *UserHandler) PurchaseProduct(w http.ResponseWriter, r *http.Request) {
	println("order product")

	txHandle := transaction.GetTxKey(r.Context())

	orderRequest := requests.OrderRequest{}
	json.NewDecoder(r.Body).Decode(&orderRequest)
	if err := h.userUsecase.WithTx(txHandle).PurchaseProduct(orderRequest); err != nil {
		response.Render(w, err, nil)
		// w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	response.Render(w, nil, "order taken successfully !")
}
