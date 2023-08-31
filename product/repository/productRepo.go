package repository

import (
	"log"

	"gorm.io/gorm"
	CommonError "main.go/errors"
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
		return nil, CommonError.ErrInternalServerError
	}
	return products, nil
}

func (r *ProductRepo) DeleteByIDs(ids []string) error {
	var products []*model.Product
	if err := r.DB.Where("id IN ?", ids).Delete(&products).Error; err != nil {
		return CommonError.ErrInternalServerError
	}
	return nil
}

func (r ProductRepo) WithTx(txHandle *gorm.DB) ProductRepo {
	if txHandle == nil {
		log.Println("no transaction db found")
		return r
	}
	r.DB = txHandle
	return r
}

func (r *ProductRepo) FindById(productId int) (*model.Product, error) {
	var product model.Product
	if err := r.DB.First(&product, productId).Error; err != nil {
		return nil, CommonError.ErrNotFound
	}
	return &product, nil
}

func (r ProductRepo) ReduceStockAmount(productId int, amount int) error {
	product := model.Product{}
	if err := r.DB.First(&product, productId).Error; err != nil {
		return CommonError.ErrNotFound
	}
	product.Stock -= amount
	if err := r.DB.Save(&product).Error; err != nil {
		return CommonError.ErrInternalServerError
	}
	return nil
}
