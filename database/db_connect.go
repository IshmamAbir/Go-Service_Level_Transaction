package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/model"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=democrud port=5433 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	println("connected to database")
	db.AutoMigrate(&model.ShoppingCart{},
		&model.Product{},
		&model.User{})
	return db, nil
}
