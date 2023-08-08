package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"main.go/user/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase, router *mux.Router, db *gorm.DB) {
	handler := &UserHandler{
		userUsecase: userUsecase,
	}
	subroute := router.PathPrefix("/users").Subrouter()
	subroute.HandleFunc("/", handler.GetAll).Methods("GET")
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
