package web

type CartItemResponse struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
	Subtotal    float32 `json:"subtotal"`
}
