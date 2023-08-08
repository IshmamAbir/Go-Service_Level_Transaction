package requests

type OrderRequest struct {
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
