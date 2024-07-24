package service

import (
	"context"
	"database/sql"
	"qbit_case/helper"
	"qbit_case/model/repository"
	"qbit_case/model/web"
)

type ProductServiceImpl struct {
	DB                *sql.DB
	productRepository repository.ProductRepository
}

func NewProductService(DB *sql.DB, productRepository repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{DB: DB, productRepository: productRepository}
}

func (service *ProductServiceImpl) GetAllProductsData(ctx context.Context, params web.ProductParams) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products, err := service.productRepository.GetAllProducts(ctx, tx, params)
	helper.PanicIfError(err)

	return helper.ToProductResponse(products)
}
