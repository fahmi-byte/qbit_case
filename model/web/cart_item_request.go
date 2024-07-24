package web

type CartItemRequest struct {
	UserId    int     `json:"user_id"`
	ProductId int     `validate:"required" json:"product_id"`
	Quantity  int     `validate:"required" json:"quantity"`
	Price     float32 `validate:"required" json:"price"`
	SubTotal  float32 `validate:"required" json:"sub_total"`
}
