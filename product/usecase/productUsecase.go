package usecase

import (
	"main.go/model"
	"main.go/product/repository"
)

type ProductUsecase struct {
	ProductRepo repository.ProductRepo
}

func NewProductUsecase(productRepo repository.ProductRepo) *ProductUsecase {
	return &ProductUsecase{
		ProductRepo: productRepo,
	}
}

func (u *ProductUsecase) FindAll() ([]*model.Product, error) {
	return u.ProductRepo.FindAll()
}

func (u *ProductUsecase) FindById(productId int) (*model.Product, error) {
	return u.ProductRepo.FindById(productId)
}

func (u *ProductUsecase) ReduceStockAmount(productId int, amount int) error {
	return u.ProductRepo.ReduceStockAmount(productId, amount)
}
