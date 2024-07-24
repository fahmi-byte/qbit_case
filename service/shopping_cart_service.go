package service

import (
	"context"
	"qbit_case/model/web"
)

type ShoppingCartService interface {
	GetAllItemsCartData(ctx context.Context, userId int) []web.CartItemResponse
	AddCartItemData(ctx context.Context, request web.CartItemRequest) error
	DeleteCartItemData(ctx context.Context, request web.DeleteCartItemRequest) error
	UpdateCartItemData(ctx context.Context, request web.UpdateCartItemRequest) error
}
