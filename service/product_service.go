package service

import (
	"context"
	"qbit_case/model/web"
)

type ProductService interface {
	GetAllProductsData(ctx context.Context, params web.ProductParams) []web.ProductResponse
}
