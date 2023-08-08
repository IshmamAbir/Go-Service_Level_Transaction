package repository

import (
	"gorm.io/gorm"
	"main.go/model"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductsRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{
		DB: db,
	}
}

func (r *ProductRepo) FindAll() ([]*model.Product, error) {
	var products []*model.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepo) FindById(productId int) (*model.Product, error) {
	var product model.Product
	if err := r.DB.First(&product, productId).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepo) ReduceStockAmount(productId int, amount int) error {
	product := model.Product{}
	if err := r.DB.First(&product, productId).Error; err != nil {
		return err
	}
	product.Stock -= amount
	if err := r.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}
