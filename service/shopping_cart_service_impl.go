package service

import (
	"context"
	"database/sql"
	"qbit_case/helper"
	"qbit_case/model/domain"
	"qbit_case/model/repository"
	"qbit_case/model/web"
	"time"
)

type ShoppingCartServiceImpl struct {
	DB                     *sql.DB
	shoppingCartRepository repository.ShoppingCartRepository
}

func NewShoppingCartService(DB *sql.DB, shoppingCartRepository repository.ShoppingCartRepository) *ShoppingCartServiceImpl {
	return &ShoppingCartServiceImpl{DB: DB, shoppingCartRepository: shoppingCartRepository}
}

func (service *ShoppingCartServiceImpl) GetAllItemsCartData(ctx context.Context, userId int) []web.CartItemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cartId := service.shoppingCartRepository.GetCartIdByUserId(ctx, tx, userId)

	cartItems := service.shoppingCartRepository.GetAllItemsCartByShoppingCartId(ctx, tx, cartId)

	return helper.ToCartItemsResponse(cartItems)
}

func (service *ShoppingCartServiceImpl) AddCartItemData(ctx context.Context, request web.CartItemRequest) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cartId := service.shoppingCartRepository.GetCartIdByUserId(ctx, tx, request.UserId)

	cartItem := domain.CartItem{
		CartID:    cartId,
		ProductID: request.ProductId,
		Quantity:  request.Quantity,
		Price:     request.Price,
		Subtotal:  request.SubTotal,
	}

	err = service.shoppingCartRepository.AddCartItem(ctx, tx, cartItem)
	helper.PanicIfError(err)

	return nil
}

func (service *ShoppingCartServiceImpl) DeleteCartItemData(ctx context.Context, request web.DeleteCartItemRequest) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cartId := service.shoppingCartRepository.GetCartIdByUserId(ctx, tx, request.UserId)

	cartItem := domain.CartItem{
		CartID:    cartId,
		ProductID: request.ProductId,
	}

	err = service.shoppingCartRepository.DeleteCartItem(ctx, tx, cartItem)
	helper.PanicIfError(err)

	return nil
}

func (service *ShoppingCartServiceImpl) UpdateCartItemData(ctx context.Context, request web.UpdateCartItemRequest) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cartId := service.shoppingCartRepository.GetCartIdByUserId(ctx, tx, request.UserId)

	subTotal := request.Price * float32(request.Quantity)

	cartItem := domain.CartItem{
		CartID:    cartId,
		ProductID: request.ProductId,
		Quantity:  request.Quantity,
		Price:     request.Price,
		Subtotal:  subTotal,
		UpdatedAt: time.Now(),
	}

	err = service.shoppingCartRepository.UpdateCartItem(ctx, tx, cartItem)
	helper.PanicIfError(err)

	return nil
}
