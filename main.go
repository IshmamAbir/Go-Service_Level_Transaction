package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"main.go/database"
	_ProductUsecase "main.go/product/usecase"
	_ShoppingCartUsecase "main.go/shopping_cart/usecase"
	_UserUsecase "main.go/user/usecase"

	_ProductHandler "main.go/product/deliver/http"
	_ShoppingCartHandler "main.go/shopping_cart/deliver/http"
	_UserHandler "main.go/user/deliver/http"

	_ProductRepo "main.go/product/repository"
	_ShoppingCartRepo "main.go/shopping_cart/repository"
	_UserRepo "main.go/user/repository"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		println(err)
		return
	}

	router := mux.NewRouter()

	userRepo := _UserRepo.NewUserRepo(db)
	productRepo := _ProductRepo.NewProductsRepo(db)
	shoppingCartRepo := _ShoppingCartRepo.NewShoppingCartRepo(db)

	productUsecase := _ProductUsecase.NewProductUsecase(*productRepo)
	shoppingCartUsecase := _ShoppingCartUsecase.NewShoppingCartUsecase(*shoppingCartRepo)
	userUsecase := _UserUsecase.NewUserUsecase(*userRepo, *productUsecase, *shoppingCartUsecase)

	_ProductHandler.NewProductHandler(*productUsecase, router)
	_ShoppingCartHandler.NewShoppingCartHandler(*shoppingCartUsecase, router)
	_UserHandler.NewUserHandler(*userUsecase, router, db)

	println("server running")
	println("--------------------")
	http.ListenAndServe(":8080", router)
}
