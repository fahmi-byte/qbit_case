package web

import "qbit_case/model/domain"

type OrderRequest struct {
	UserId          int                `json:"user_id"`
	TotalAmount     float32            `json:"total_amount"`
	DeliveryAddress string             `json:"delivery_address"`
	OrderItems      []domain.OrderItem `json:"order_items"`
}
