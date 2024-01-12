package repository

import (
	"context"

	"gorm.io/gorm"
	CommonError "main.go/errors"
	"main.go/model"
	"main.go/transaction"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductsRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{
		DB: db,
	}
}

func (r *ProductRepo) FindAll(ctx context.Context) ([]*model.Product, error) {
	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = r.DB
	}
	var products []*model.Product
	if err := tx.Find(&products).Error; err != nil {
		return nil, CommonError.ErrInternalServerError
	}
	return products, nil
}

func (r *ProductRepo) FindById(ctx context.Context, productId int) (*model.Product, error) {
	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = r.DB
	}
	var product model.Product
	if err := tx.First(&product, productId).Error; err != nil {
		return nil, CommonError.ErrNotFound
	}
	return &product, nil
}

func (r ProductRepo) ReduceStockAmount(ctx context.Context, productId int, amount int) error {
	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = r.DB
	}
	product := model.Product{}
	if err := tx.First(&product, productId).Error; err != nil {
		return CommonError.ErrNotFound
	}
	product.Stock -= amount
	if err := tx.Save(&product).Error; err != nil {
		return CommonError.ErrInternalServerError
	}
	return nil
}
