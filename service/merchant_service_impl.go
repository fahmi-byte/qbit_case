package service

import (
	"context"
	"database/sql"
	"qbit_case/helper"
	"qbit_case/model/domain"
	"qbit_case/model/repository"
	"qbit_case/model/web"
)

type MerchantServiceImpl struct {
	DB                 *sql.DB
	merchantRepository repository.MerchantRepository
	userRepository     repository.UserRepository
	productRepository  repository.ProductRepository
}

func NewMerchantService(DB *sql.DB, merchantRepository repository.MerchantRepository, userRepository repository.UserRepository, productRepository repository.ProductRepository) *MerchantServiceImpl {
	return &MerchantServiceImpl{DB: DB, merchantRepository: merchantRepository, userRepository: userRepository, productRepository: productRepository}
}

func (service *MerchantServiceImpl) RegisterMerchant(ctx context.Context, request web.RegisterMerchantRequest) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.merchantRepository.CreateMerchant(ctx, tx, request)
	helper.PanicIfError(err)

	err = service.userRepository.UpdateToMerchant(ctx, tx, request.UserId)
	helper.PanicIfError(err)

	return nil
}

func (service *MerchantServiceImpl) AddProductData(ctx context.Context, request web.AddProductRequest) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	discountPrice := request.Price * (request.Discount / 100)

	finalPrice := request.Price - discountPrice

	product := domain.Product{
		MerchantId:  request.MerchantId,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Discount:    request.Discount,
		FinalPrice:  finalPrice,
		Stock:       request.Stock,
		CategoryId:  request.CategoryId,
	}

	err = service.productRepository.CreateProduct(ctx, tx, product)

	return nil
}
