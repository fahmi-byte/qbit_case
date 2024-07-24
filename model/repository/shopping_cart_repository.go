package repository

import (
	"context"
	"database/sql"
	"qbit_case/model/domain"
)

type ShoppingCartRepository interface {
	CreateCart(ctx context.Context, tx *sql.Tx, userId int) error
	GetCartIdByUserId(ctx context.Context, tx *sql.Tx, userId int) int
	GetAllItemsCartByShoppingCartId(ctx context.Context, tx *sql.Tx, shoppingCartId int) []domain.CartItem
	AddCartItem(ctx context.Context, tx *sql.Tx, request domain.CartItem) error
	DeleteCartItem(ctx context.Context, tx *sql.Tx, request domain.CartItem) error
	DeleteBatchCartItem(ctx context.Context, tx *sql.Tx, products []int, cartId int) error
	UpdateCartItem(ctx context.Context, tx *sql.Tx, request domain.CartItem) error
}
