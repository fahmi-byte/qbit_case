package domain

import "time"

type CartItem struct {
	Id          int `json:"cart_item_id"`
	CartID      int `json:"cart_id"`
	ProductID   int `json:"product_id"`
	ProductName string
	Quantity    int       `json:"quantity"`
	Price       float32   `json:"price"`
	Subtotal    float32   `json:"subtotal"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
