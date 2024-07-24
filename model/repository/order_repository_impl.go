package repository

import (
	"context"
	"database/sql"
	"fmt"
	"qbit_case/constant"
	"qbit_case/helper"
	"qbit_case/model/domain"
	"strconv"
	"strings"
	"time"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) PaymentOrderSuccess(ctx context.Context, tx *sql.Tx, orderId int, amount float32) {
	SQL := "UPDATE orders SET status = $1, payment_status = $2, payment_date = $3 WHERE id = $4"
	_, err := tx.ExecContext(ctx, SQL, constant.SHIPPED, constant.PAYMENT_STATUS[constant.SUCCESS], time.Now(), orderId)
	helper.PanicIfError(err)
	fmt.Println(amount, "berapa")

	var productId int
	SQLGet := "SELECT product_id from order_items WHERE order_id = $1 LIMIT 1"
	err = tx.QueryRowContext(ctx, SQLGet, orderId).Scan(&productId)

	SQLUpdateMerchantBalance := `UPDATE merchants
SET balance = $1
    FROM products
WHERE merchants.id = products.merchant_id
  AND products.id = $2
`
	_, err = tx.ExecContext(ctx, SQLUpdateMerchantBalance, amount, productId)
	helper.PanicIfError(err)
}

func (repository *OrderRepositoryImpl) PaymentOrderFailed(ctx context.Context, tx *sql.Tx, orderId int) {
	SQL := "UPDATE orders SET status = $1, payment_status = $2, payment_date = $3 WHERE id = $4"
	_, err := tx.ExecContext(ctx, SQL, constant.CANCELED, constant.PAYMENT_STATUS[constant.FAILED], time.Now(), orderId)
	helper.PanicIfError(err)

	SQLGet := "SELECT product_id, quantity from order_items WHERE order_id = $1"
	rows, err := tx.QueryContext(ctx, SQLGet, orderId)
	helper.PanicIfError(err)
	defer rows.Close()

	orderItems := []domain.OrderItem{}
	for rows.Next() {
		var orderItem domain.OrderItem
		err = rows.Scan(&orderItem.ProductId, &orderItem.Quantity)
		helper.PanicIfError(err)
		orderItems = append(orderItems, orderItem)
	}

	// Membuat query UPDATE dinamis
	var caseStatements []string
	var ids []string
	var args []interface{}

	for i, item := range orderItems {
		// Membuat CASE untuk setiap produk
		caseStatements = append(caseStatements, fmt.Sprintf("WHEN product_id = $%d THEN $%d", i*2+1, i*2+2))
		args = append(args, item.ProductId, item.Quantity)
		ids = append(ids, fmt.Sprintf("$%d", i*2+1))
	}

	// Membuat query lengkap
	query := fmt.Sprintf(`
		UPDATE products
		SET stock_quantity = stock_quantity + CASE
		%s
		ELSE 0
		END
		WHERE product_id IN (%s)
	`, strings.Join(caseStatements, " "), strings.Join(ids, ", "))

	// Menampilkan query untuk debugging
	fmt.Println("Executing query:", query)

	_, err = tx.ExecContext(ctx, query, args...)
	helper.PanicIfError(err)

}

func (repository *OrderRepositoryImpl) CreateOrder(ctx context.Context, tx *sql.Tx, request domain.Order) (string, error) {
	var orderId int
	SQL := "INSERT INTO orders (user_id, order_number, total_amount, status, payment_status, order_date, delivery_address) VALUES($1, generate_new_order_number(), $2, $3, $4, $5, $6) RETURNING id"
	err := tx.QueryRowContext(ctx, SQL, request.UserId, request.TotalAmount, request.Status, request.PaymentStatus, request.OrderDate, request.DeliveryAddress).Scan(&orderId)
	helper.PanicIfError(err)

	fmt.Println(request.OrderItems)

	var values []string
	var args []interface{}

	for i, item := range request.OrderItems {
		placeholders := fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5)
		values = append(values, placeholders)
		args = append(args, orderId, item.ProductId, item.Quantity, item.Price, item.SubTotal)
	}

	query := fmt.Sprintf("INSERT INTO order_items (order_id, product_id, quantity, price, subtotal) VALUES %s",
		strings.Join(values, ", "))

	fmt.Println(query)

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		panic(err)
	}

	orderNumber := "ORD-" + strconv.Itoa(orderId)

	return orderNumber, nil
}

func (repository *OrderRepositoryImpl) FindAllUserOrders(ctx context.Context, tx *sql.Tx, userId int) []domain.Order {
	SQL := "SELECT order_number, total_amount, status, payment_status, order_date, delivery_address FROM orders"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var orders []domain.Order
	for rows.Next() {
		order := domain.Order{}
		err = rows.Scan(&order.OrderNumber, &order.TotalAmount, &order.Status, &order.PaymentStatus, &order.OrderDate, &order.DeliveryAddress)
		helper.PanicIfError(err)
		orders = append(orders, order)
	}

	return orders
}
