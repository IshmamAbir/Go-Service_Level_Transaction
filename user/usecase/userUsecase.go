package usecase

// app 1

import (
	"context"

	"main.go/model"
	pr "main.go/product/repository"
	"main.go/product/usecase"
	"main.go/requests"
	shoppingCartUsecase "main.go/shopping_cart/usecase"
	"main.go/transaction"
	"main.go/user/repository"
)

type UserUsecase struct {
	UserRepo            repository.UserRepo
	ProductUsecase      usecase.ProductUsecase
	ProductRepo         pr.ProductRepo
	UW                  transaction.UoW
	ShoppingCartUsecase shoppingCartUsecase.ShoppingCartUsecase
}

func NewUserUsecase(userRepo repository.UserRepo,
	productUsecase usecase.ProductUsecase,
	prepo pr.ProductRepo,
	uw transaction.UoW,
	shoppingCartUsecase shoppingCartUsecase.ShoppingCartUsecase,
) *UserUsecase {
	return &UserUsecase{
		UserRepo:       userRepo,
		ProductUsecase: productUsecase,
		ProductRepo:    prepo,
		UW:             uw,
	}
}

func (u *UserUsecase) FindAll(ctx context.Context) ([]*model.User, error) {
	return u.UserRepo.FindAll(ctx)
}

func (u UserUsecase) ReduceBalance(ctx context.Context, userId int, productId int, productQuantity int) error {
	product, err := u.ProductUsecase.FindById(ctx, productId)
	if err != nil {
		return err
	}
	amount := product.Price * productQuantity
	return u.UserRepo.ReduceBalance(ctx, userId, amount)
}

// -------------------

func (u UserUsecase) PurchaseProduct(ctx context.Context, orderRequest requests.OrderRequest) error {

	v, err := u.UW.WithTx(ctx, func(ctx context.Context) (interface{}, error) {

		// step 1 (business logic): add to shopping cart
		shoppingCart := model.ShoppingCart{}
		shoppingCart.UserId = orderRequest.UserId
		shoppingCart.ProductId = orderRequest.ProductId
		shoppingCart.ProductAmount = orderRequest.Quantity
		if err := u.ShoppingCartUsecase.AddToShoppingCart(ctx, &shoppingCart); err != nil {
			return nil, err
		}

		// step 2 (business logic): reduce product stock
		if err := u.ProductUsecase.ReduceStockAmount(ctx, orderRequest.ProductId, orderRequest.Quantity); err != nil {
			return nil, err
		}
		// step 3 (business logic): reduce user balance
		product, err := u.ProductUsecase.FindById(ctx, orderRequest.ProductId)
		if err != nil {
			return nil, err
		}
		if err := u.UserRepo.ReduceBalance(ctx, orderRequest.UserId, (product.Price * orderRequest.Quantity)); err != nil {
			return nil, err
		}

		return "product purchased", nil

	})
	if err != nil {
		return err
	}

	println(v)
	return err
}
