package usecase

import (
	"gorm.io/gorm"
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

func (u UserUsecase) ReduceBalance(userId int, productId int, productQuantity int) error {
	product, err := u.ProductUsecase.FindById(productId)
	if err != nil {
		return err
	}
	amount := product.Price * productQuantity
	return u.UserRepo.ReduceBalance(userId, amount)
}

// -------------------

func (u UserUsecase) WithTx(txHandle *gorm.DB) UserUsecase {
	u.UserRepo = u.UserRepo.WithTx(txHandle)
	u.ProductUsecase = u.ProductUsecase.WithTx(txHandle)
	u.ShoppingCartUsecase = u.ShoppingCartUsecase.WithTx(txHandle)
	return u
}

func (u UserUsecase) OrderProduct(orderRequest requests.OrderRequest) error {
	// step 1 (business logic): add to shopping cart
	shoppingCart := model.ShoppingCart{}
	shoppingCart.UserId = orderRequest.UserId
	shoppingCart.ProductId = orderRequest.ProductId
	shoppingCart.ProductAmount = orderRequest.Quantity
	if err := u.ShoppingCartUsecase.AddToShoppingCart(&shoppingCart); err != nil {
		return err
	}
	// step 2 (business logic): reduce product stock
	if err := u.ProductUsecase.ReduceStockAmount(orderRequest.ProductId, orderRequest.Quantity); err != nil {
		return err
	}
	// step 3 (business logic): reduce user balance
	product, err := u.ProductUsecase.FindById(orderRequest.ProductId)
	if err != nil {
		return err
	}
	if err := u.UserRepo.ReduceBalance(orderRequest.UserId, (product.Price * orderRequest.Quantity)); err != nil {
		return err
	}
	return nil
}
