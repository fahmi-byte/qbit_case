package service

import (
	"context"
	"database/sql"
	"qbit_case/exception"
	"qbit_case/helper"
	"qbit_case/model/repository"
	"qbit_case/model/web"
)

type AuthServiceImpl struct {
	DB                     *sql.DB
	cryptoRepository       repository.CryptoRepository
	userRepository         repository.UserRepository
	merchantRepository     repository.MerchantRepository
	shoppingCartRepository repository.ShoppingCartRepository
}

func NewAuthService(DB *sql.DB, cryptoRepository repository.CryptoRepository, userRepository repository.UserRepository, merchantRepository repository.MerchantRepository, shoppingCartRepository repository.ShoppingCartRepository) *AuthServiceImpl {
	return &AuthServiceImpl{DB: DB, cryptoRepository: cryptoRepository, userRepository: userRepository, merchantRepository: merchantRepository, shoppingCartRepository: shoppingCartRepository}
}

func (service *AuthServiceImpl) AuthRegister(ctx context.Context, request web.RegisterRequest) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.userRepository.FindByEmail(ctx, tx, request.Email)

	if err == nil {
		panic(exception.NewBadRequestError("Email already exist!"))
	}

	hashedPassword := service.cryptoRepository.GenerateHashPassword(request.Password)
	request.Password = string(hashedPassword)

	userId, err := service.userRepository.CreateUser(ctx, tx, request)
	if err != nil {
		return err
	}

	err = service.shoppingCartRepository.CreateCart(ctx, tx, userId)
	helper.PanicIfError(err)

	return nil
}

func (service *AuthServiceImpl) AuthLogin(ctx context.Context, request web.LoginRequest) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.userRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	checkPassword := service.cryptoRepository.CompareHashAndPassword(user.Password, request.Password)

	if !checkPassword {
		panic(exception.NewBadRequestError("wrong password!"))
	}

	merchantId := service.merchantRepository.FindMerchantIdByUserId(ctx, tx, user.Id)

	token, err := service.cryptoRepository.GenerateJWT(user, merchantId)
	helper.PanicIfError(err)

	return token
}
