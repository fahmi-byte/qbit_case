package repository

import (
	"context"
	"database/sql"
	"fmt"
	"qbit_case/exception"
	"qbit_case/helper"
	"qbit_case/model/domain"
	"strings"
)

type ShoppingCartRepositoryImpl struct {
}

func NewShoppingCartRepository() *ShoppingCartRepositoryImpl {
	return &ShoppingCartRepositoryImpl{}
}

func (repository *ShoppingCartRepositoryImpl) CreateCart(ctx context.Context, tx *sql.Tx, userId int) error {
	SQL := "INSERT INTO shopping_cart (user_id) VALUES($1)"
	_, err := tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	return nil
}

func (repository *ShoppingCartRepositoryImpl) GetCartIdByUserId(ctx context.Context, tx *sql.Tx, userId int) int {
	var cartId int
	SQL := "SELECT id from shopping_cart WHERE user_id = $1 LIMIT 1"
	err := tx.QueryRowContext(ctx, SQL, userId).Scan(&cartId)
	if err != nil {
		panic(exception.NewNotFoundError("Cart Id Not Found"))
	}

	return cartId
}

func (repository *ShoppingCartRepositoryImpl) GetAllItemsCartByShoppingCartId(ctx context.Context, tx *sql.Tx, shoppingCartId int) []domain.CartItem {
	SQL := "SELECT product_id, p.name, quantity, ci.price, subtotal FROM cart_items ci " +
		"JOIN products p ON ci.product_id = p.id WHERE cart_id = $1"

	rows, err := tx.QueryContext(ctx, SQL, shoppingCartId)
	helper.PanicIfError(err)
	defer rows.Close()

	var cartItems []domain.CartItem
	for rows.Next() {
		var cartItem domain.CartItem
		if err := rows.Scan(&cartItem.ProductID, &cartItem.ProductName, &cartItem.Quantity, &cartItem.Price, &cartItem.Subtotal); err != nil {
			return nil
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems
}

func (repository *ShoppingCartRepositoryImpl) AddCartItem(ctx context.Context, tx *sql.Tx, request domain.CartItem) error {
	SQL := "INSERT INTO cart_items (cart_id, product_id, quantity, price, subtotal) VALUES($1, $2, $3, $4, $5)"
	_, err := tx.ExecContext(ctx, SQL, request.CartID, request.ProductID, request.Quantity, request.Price, request.Subtotal)
	helper.PanicIfError(err)
	return nil
}

func (repository *ShoppingCartRepositoryImpl) DeleteCartItem(ctx context.Context, tx *sql.Tx, request domain.CartItem) error {
	SQL := "DELETE FROM cart_items WHERE product_id = $1 AND cart_id = $2"
	_, err := tx.ExecContext(ctx, SQL, request.ProductID, request.CartID)
	helper.PanicIfError(err)
	return nil
}

func (repository *ShoppingCartRepositoryImpl) UpdateCartItem(ctx context.Context, tx *sql.Tx, request domain.CartItem) error {
	SQL := "UPDATE cart_items SET quantity = $1, price = $2, subtotal = $3, updated_at = $4 WHERE product_id = $5 AND cart_id = $6"
	_, err := tx.ExecContext(ctx, SQL, request.Quantity, request.Price, request.Subtotal, request.UpdatedAt, request.ProductID, request.CartID)
	helper.PanicIfError(err)
	return nil
}

func (s ShoppingCartRepositoryImpl) DeleteBatchCartItem(ctx context.Context, tx *sql.Tx, products []int, cartId int) error {
	var placeholders []string
	var args []interface{}

	// Tempatkan cartId sebagai argumen pertama
	args = append(args, cartId)

	// Buat placeholder untuk setiap productId
	for i, id := range products {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+2))
		args = append(args, id)
	}

	query := fmt.Sprintf("DELETE FROM cart_items WHERE cart_id = $1 AND product_id IN (%s)", strings.Join(placeholders, ", "))

	fmt.Println("Executing query:", query)

	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete cart items: %v", err)
	}

	return nil
}
