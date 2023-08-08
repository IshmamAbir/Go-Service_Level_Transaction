package usecase

import (
	"main.go/model"
	"main.go/product/usecase"
	"main.go/requests"
	shoppingCartUsecase "main.go/shopping_cart/usecase"
	"main.go/user/repository"
)

type UserUsecase struct {
	UserRepo            repository.UserRepo
	ProductUsecase      usecase.ProductUsecase
	ShoppingCartUsecase shoppingCartUsecase.ShoppingCartUsecase
}

func NewUserUsecase(userRepo repository.UserRepo,
	productUsecase usecase.ProductUsecase,
	shoppingCartUsecase shoppingCartUsecase.ShoppingCartUsecase) *UserUsecase {
	return &UserUsecase{
		UserRepo:       userRepo,
		ProductUsecase: productUsecase,
	}
}

func (u *UserUsecase) FindAll() ([]*model.User, error) {
	return u.UserRepo.FindAll()
}

func (u *UserUsecase) ReduceBalance(userId int, productId int, productQuantity int) error {
	product, err := u.ProductUsecase.FindById(productId)
	if err != nil {
		return err
	}
	amount := product.Price * productQuantity
	return u.UserRepo.ReduceBalance(userId, amount)
}

// -------------------

func (u *UserUsecase) OrderProduct(orderRequest requests.OrderRequest) error {
	// step 1: add to shopping cart
	// step 2: reduce product stock
	// step 3: reduce user balance
	// step 4: create order
	return nil
}
