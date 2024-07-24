package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"qbit_case/app"
	"qbit_case/config"
	"qbit_case/controller"
	"qbit_case/helper"
	"qbit_case/middleware"
	"qbit_case/model/repository"
	"qbit_case/service"
)

func NewValidator() *validator.Validate {
	return validator.New()
}

func NewServer(c config.Server, authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    c.Address(),
		Handler: authMiddleware,
	}
}

func main() {
	fmt.Println("Starting Server")
	cfg := config.LoadConfig()

	db := app.NewDB(cfg.Database)
	validate := NewValidator()

	cryptoRepositoryImpl := repository.NewCryptoRepository(cfg.Auths())
	userRepositoryImpl := repository.NewUserRepository()
	merchantRepositoryImpl := repository.NewMerchantRepository()
	productRepositoryImpl := repository.NewProductRepository()
	shoppingCartRepositoryImpl := repository.NewShoppingCartRepository()
	orderRepositoryImpl := repository.NewOrderRepository()

	authServiceImpl := service.NewAuthService(db, cryptoRepositoryImpl, userRepositoryImpl, merchantRepositoryImpl, shoppingCartRepositoryImpl)
	merchantServiceImpl := service.NewMerchantService(db, merchantRepositoryImpl, userRepositoryImpl, productRepositoryImpl)
	productServiceImpl := service.NewProductService(db, productRepositoryImpl)
	shoppingCartServiceImpl := service.NewShoppingCartService(db, shoppingCartRepositoryImpl)
	orderServiceImpl := service.NewOrderService(db, orderRepositoryImpl, shoppingCartRepositoryImpl)

	authControllerImpl := controller.NewAuthController(validate, authServiceImpl)
	merchantControllerImpl := controller.NewMerchantController(validate, merchantServiceImpl)
	productControllerImpl := controller.NewProductController(productServiceImpl)
	shoppingCartControllerImpl := controller.NewShoppingCartController(validate, shoppingCartServiceImpl)
	orderControllerImpl := controller.NewOrderController(validate, orderServiceImpl, cfg.AuthPayment())

	router := app.NewRouter(
		authControllerImpl,
		merchantControllerImpl,
		productControllerImpl,
		shoppingCartControllerImpl,
		orderControllerImpl,
	)

	authMiddleware := middleware.NewAuthMiddleware(router, db, cryptoRepositoryImpl)

	server := NewServer(cfg.Server, authMiddleware)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
