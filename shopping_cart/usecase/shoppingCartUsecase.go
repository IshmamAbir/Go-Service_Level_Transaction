package usecase

import (
	"gorm.io/gorm"
	"main.go/model"
	"main.go/shopping_cart/repository"
)

type ShoppingCartUsecase struct {
	ShoppingCartRepo repository.ShoppingCartRepo
}

func NewShoppingCartUsecase(shoppingCartRepo repository.ShoppingCartRepo) *ShoppingCartUsecase {
	return &ShoppingCartUsecase{
		ShoppingCartRepo: shoppingCartRepo,
	}
}

func (u ShoppingCartUsecase) AddToShoppingCart(shoppingCart *model.ShoppingCart) error {
	return u.ShoppingCartRepo.Create(shoppingCart)
}

func (u ShoppingCartUsecase) FindAll() ([]*model.ShoppingCart, error) {
	return u.ShoppingCartRepo.FindAll()
}

func (u ShoppingCartUsecase) WithTx(txHandle *gorm.DB) ShoppingCartUsecase {
	u.ShoppingCartRepo = u.ShoppingCartRepo.WithTx(txHandle)
	return u
}
