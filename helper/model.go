package helper

import (
	"qbit_case/model/domain"
	"qbit_case/model/web"
)

func ToProductResponse(products []domain.Product) []web.ProductResponse {
	var productsResponse []web.ProductResponse
	for _, product := range products {
		productsResponse = append(productsResponse, web.ProductResponse{
			Id:           product.Id,
			Name:         product.Name,
			MerchantName: product.MerchantName,
			MerchantCity: product.MerchantCity,
			Description:  product.Description,
			Price:        product.Price,
			Discount:     product.Discount,
			FinalPrice:   product.FinalPrice,
			Stock:        product.Stock,
			CategoryName: product.CategoryName,
		})
	}
	return productsResponse
}

func ToCartItemsResponse(cartItems []domain.CartItem) []web.CartItemResponse {
	var productsResponse []web.CartItemResponse
	for _, cartItem := range cartItems {
		productsResponse = append(productsResponse, web.CartItemResponse{
			ProductID:   cartItem.ProductID,
			ProductName: cartItem.ProductName,
			Quantity:    cartItem.Quantity,
			Price:       cartItem.Price,
			Subtotal:    cartItem.Subtotal,
		})
	}
	return productsResponse
}

func ToOrderResponse(orders []domain.Order) []web.OrderResponse {
	var orderResponse []web.OrderResponse
	for _, order := range orders {
		orderResponse = append(orderResponse, web.OrderResponse{
			Id:              order.Id,
			OrderNumber:     order.OrderNumber,
			OrderDate:       order.OrderDate,
			Status:          order.Status,
			PaymentDate:     order.PaymentDate,
			PaymentStatus:   order.PaymentStatus,
			TotalAmount:     order.TotalAmount,
			DeliveryAddress: order.DeliveryAddress,
		})
	}
	return orderResponse
}
