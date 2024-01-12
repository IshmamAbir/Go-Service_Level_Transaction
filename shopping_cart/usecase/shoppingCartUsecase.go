package usecase

import (
	"context"

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

func (u ShoppingCartUsecase) AddToShoppingCart(ctx context.Context, shoppingCart *model.ShoppingCart) error {
	return u.ShoppingCartRepo.Create(ctx, shoppingCart)
}

func (u ShoppingCartUsecase) FindAll(ctx context.Context) ([]*model.ShoppingCart, error) {
	return u.ShoppingCartRepo.FindAll(ctx)
}
