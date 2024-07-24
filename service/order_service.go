package service

import (
	"context"
	"qbit_case/model/web"
)

type OrderService interface {
	CreateNewOrder(ctx context.Context, request web.OrderRequest) (string, error)
	GetUserOrdersData(ctx context.Context, userId int) []web.OrderResponse
	PaymentOrderCallback(ctx context.Context, orderId int, status string, amount float32)
}
