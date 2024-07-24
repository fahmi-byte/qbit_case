package repository

import (
	"context"
	"database/sql"
	"qbit_case/model/domain"
)

type OrderRepository interface {
	PaymentOrderSuccess(ctx context.Context, tx *sql.Tx, orderId int, amount float32)
	PaymentOrderFailed(ctx context.Context, tx *sql.Tx, orderId int)
	CreateOrder(ctx context.Context, tx *sql.Tx, request domain.Order) (string, error)
	FindAllUserOrders(ctx context.Context, tx *sql.Tx, userId int) []domain.Order
}
