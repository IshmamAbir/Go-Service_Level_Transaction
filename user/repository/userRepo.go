package repository

import (
	"log"

	CommonError "main.go/errors"

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

func (r UserRepo) FindAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, CommonError.ErrInternalServerError
	}
	return users, nil
}

func (r UserRepo) WithTx(txHandle *gorm.DB) UserRepo {
	if txHandle == nil {
		log.Println("no transaction db found")
		return r
	}
	r.DB = txHandle
	return r
}

func (r UserRepo) ReduceBalance(userId int, amount int) error {
	user := model.User{}
	if err := r.DB.First(&user, userId).Error; err != nil {
		return CommonError.ErrNotFound
	}
	user.Balance -= amount
	if err := r.DB.Save(&user).Error; err != nil {
		return CommonError.ErrInternalServerError
	}
	return nil
}
