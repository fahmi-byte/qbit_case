package web

type DeleteCartItemRequest struct {
	UserId    int `json:"user_id"`
	ProductId int `validate:"required" json:"product_id"`
}
