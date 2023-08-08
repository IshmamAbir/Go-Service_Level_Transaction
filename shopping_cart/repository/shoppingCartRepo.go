package repository

import (
	"gorm.io/gorm"
	"main.go/model"
)

type ShoppingCartRepo struct {
	DB *gorm.DB
}

func NewShoppingCartRepo(db *gorm.DB) *ShoppingCartRepo {
	return &ShoppingCartRepo{
		DB: db,
	}
}

func (r *ShoppingCartRepo) Create(shoppingCart *model.ShoppingCart) error {
	if err := r.DB.Create(shoppingCart).Error; err != nil {
		return err
	}
	return nil
}

func (r *ShoppingCartRepo) FindAll() ([]*model.ShoppingCart, error) {
	var shoppingCarts []*model.ShoppingCart
	if err := r.DB.Find(&shoppingCarts).Error; err != nil {
		return nil, err
	}
	return shoppingCarts, nil
}
