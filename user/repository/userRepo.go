package repository

import (
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

func (r *UserRepo) FindAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) ReduceBalance(userId int, amount int) error {
	user := model.User{}
	if err := r.DB.First(&user, userId).Error; err != nil {
		return err
	}
	user.Balance -= amount
	if err := r.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
