package usecase

import (
	"context"

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

func (u *ProductUsecase) FindAll(ctx context.Context) ([]*model.Product, error) {
	return u.ProductRepo.FindAll(ctx)
}

func (u *ProductUsecase) FindById(ctx context.Context, productId int) (*model.Product, error) {
	return u.ProductRepo.FindById(ctx, productId)
}

func (u ProductUsecase) ReduceStockAmount(ctx context.Context, productId int, amount int) error {
	return u.ProductRepo.ReduceStockAmount(ctx, productId, amount)
}
