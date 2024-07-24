package domain

type OrderItem struct {
	Id        int     `json:"id"`
	OrderId   int     `json:"order_id"`
	ProductId int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
	SubTotal  float32 `json:"sub_total"`
}
