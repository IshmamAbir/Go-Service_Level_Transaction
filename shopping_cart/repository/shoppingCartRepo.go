package repository

import (
	"context"

	"gorm.io/gorm"
	CommonError "main.go/errors"
	"main.go/model"
	"main.go/transaction"
)

type ShoppingCartRepo struct {
	DB *gorm.DB
}

func NewShoppingCartRepo(db *gorm.DB) *ShoppingCartRepo {
	return &ShoppingCartRepo{
		DB: db,
	}
}

func (r ShoppingCartRepo) Create(ctx context.Context, shoppingCart *model.ShoppingCart) error {
	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = r.DB
	}
	if err := tx.Create(&shoppingCart).Error; err != nil {
		return CommonError.ErrInternalServerError
	}
	return nil
}

func (r ShoppingCartRepo) FindAll(ctx context.Context) ([]*model.ShoppingCart, error) {
	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = r.DB
	}
	var shoppingCarts []*model.ShoppingCart
	if err := tx.Find(&shoppingCarts).Error; err != nil {
		return nil, CommonError.ErrInternalServerError
	}
	return shoppingCarts, nil
}
