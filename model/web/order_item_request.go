package web

type OrderItemRequest struct {
	ProductId int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
	SubTotal  float32 `json:"sub_total"`
}
