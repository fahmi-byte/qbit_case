package domain

import "time"

type Order struct {
	Id              int         `json:"id"`
	UserId          int         `json:"user_id"`
	OrderNumber     string      `json:"order_number"`
	TotalAmount     float32     `json:"total_amount"`
	Status          string      `json:"status"`
	PaymentStatus   string      `json:"payment_status"`
	PaymentDate     time.Time   `json:"payment_date"`
	OrderDate       time.Time   `json:"order_date"`
	DeliveryAddress string      `json:"delivery_address"`
	OrderItems      []OrderItem `json:"order_items"`
}
