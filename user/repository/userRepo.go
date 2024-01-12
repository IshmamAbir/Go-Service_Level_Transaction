package repository

import (
	"context"

	CommonError "main.go/errors"
	"main.go/transaction"

	"gorm.io/gorm"
	"main.go/model"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (r UserRepo) FindAll(ctx context.Context) ([]*model.User, error) {
	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = r.DB
	}
	var users []*model.User
	if err := tx.Find(&users).Error; err != nil {
		return nil, CommonError.ErrInternalServerError
	}
	return users, nil
}

func (r UserRepo) ProductsByID(ctx context.Context, id string) ([]*model.Product, error) {
	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = r.DB
	}
	var prods []*model.Product
	if err := tx.Where("user_id=?", id).Find(&prods).Error; err != nil {
		return nil, CommonError.ErrInternalServerError
	}
	return prods, nil
}

func (r UserRepo) Delete(ctx context.Context, id string) error {
	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = r.DB
	}
	var user *model.User
	if err := tx.Where("id=?", id).Delete(user).Error; err != nil {
		return CommonError.ErrInternalServerError
	}
	return nil
}

func (r UserRepo) ReduceBalance(ctx context.Context, userId int, amount int) error {
	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = r.DB
	}
	user := model.User{}
	if err := r.DB.First(&user, userId).Error; err != nil {
		return CommonError.ErrNotFound
	}
	user.Balance -= amount
	if err := tx.Save(&user).Error; err != nil {
		return CommonError.ErrInternalServerError
	}
	return nil
}
