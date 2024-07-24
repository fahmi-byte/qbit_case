package service

import (
	"context"
	"database/sql"
	"qbit_case/constant"
	"qbit_case/helper"
	"qbit_case/model/domain"
	"qbit_case/model/repository"
	"qbit_case/model/web"
	"time"
)

type OrderServiceImpl struct {
	DB                     *sql.DB
	orderRepository        repository.OrderRepository
	shoppingCartRepository repository.ShoppingCartRepository
}

func NewOrderService(DB *sql.DB, orderRepository repository.OrderRepository, shoppingCartRepository repository.ShoppingCartRepository) *OrderServiceImpl {
	return &OrderServiceImpl{DB: DB, orderRepository: orderRepository, shoppingCartRepository: shoppingCartRepository}
}

func (service *OrderServiceImpl) CreateNewOrder(ctx context.Context, request web.OrderRequest) (string, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order := domain.Order{
		UserId:          request.UserId,
		Status:          constant.PENDING,
		PaymentStatus:   constant.PENDING,
		OrderDate:       time.Now(),
		TotalAmount:     request.TotalAmount,
		OrderItems:      request.OrderItems,
		DeliveryAddress: request.DeliveryAddress,
	}

	orderNumber, err := service.orderRepository.CreateOrder(ctx, tx, order)
	helper.PanicIfError(err)

	cartId := service.shoppingCartRepository.GetCartIdByUserId(ctx, tx, request.UserId)
	var products []int

	for _, value := range request.OrderItems {
		products = append(products, value.ProductId)
	}

	err = service.shoppingCartRepository.DeleteBatchCartItem(ctx, tx, products, cartId)

	return orderNumber, nil
}

func (service *OrderServiceImpl) GetUserOrdersData(ctx context.Context, userId int) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := service.orderRepository.FindAllUserOrders(ctx, tx, userId)

	return helper.ToOrderResponse(orders)
}

func (service *OrderServiceImpl) PaymentOrderCallback(ctx context.Context, orderId int, status string, amount float32) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	switch status {
	case constant.SUCCESS:
		service.orderRepository.PaymentOrderSuccess(ctx, tx, orderId, amount)

	case constant.FAILED:
		service.orderRepository.PaymentOrderFailed(ctx, tx, orderId)
	}

}
