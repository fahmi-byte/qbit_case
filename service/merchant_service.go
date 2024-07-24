package service

import (
	"context"
	"qbit_case/model/web"
)

type MerchantService interface {
	RegisterMerchant(ctx context.Context, request web.RegisterMerchantRequest) error
	AddProductData(ctx context.Context, request web.AddProductRequest) error
}
