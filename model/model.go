package model

type User struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Balance int    `json:"balance"`
}

func (User) TableName() string {
	return "user"
}

type Product struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func (Product) TableName() string {
	return "product"
}

type ShoppingCart struct {
	Id            int `json:"id" gorm:"primaryKey"`
	UserId        int `json:"user_id"`
	ProductId     int `json:"product_id"`
	ProductAmount int `json:"product_amount"`
}

func (ShoppingCart) TableName() string {
	return "shopping_cart"
}
