package repository

import (
	"log"

	"gorm.io/gorm"
	CommonError "main.go/errors"
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

func (r ShoppingCartRepo) WithTx(txHandle *gorm.DB) ShoppingCartRepo {
	if txHandle == nil {
		log.Println("no transaction db found")
		return r
	}
	r.DB = txHandle
	return r
}

func (r ShoppingCartRepo) Create(shoppingCart *model.ShoppingCart) error {
	if err := r.DB.Create(&shoppingCart).Error; err != nil {
		return CommonError.ErrInternalServerError
	}
	return nil
}

func (r ShoppingCartRepo) FindAll() ([]*model.ShoppingCart, error) {
	var shoppingCarts []*model.ShoppingCart
	if err := r.DB.Find(&shoppingCarts).Error; err != nil {
		return nil, CommonError.ErrInternalServerError
	}
	return shoppingCarts, nil
}
