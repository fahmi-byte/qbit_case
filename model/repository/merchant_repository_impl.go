package repository

import (
	"context"
	"database/sql"
	"fmt"
	"qbit_case/helper"
	"qbit_case/model/web"
	"strconv"
)

type MerchantRepositoryImpl struct {
}

func NewMerchantRepository() *MerchantRepositoryImpl {
	return &MerchantRepositoryImpl{}
}

func (repository *MerchantRepositoryImpl) CreateMerchant(ctx context.Context, tx *sql.Tx, request web.RegisterMerchantRequest) error {
	SQL := "INSERT INTO merchants (user_id, name, address, city) VALUES($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, SQL, request.UserId, request.Name, request.Address, request.City)
	helper.PanicIfError(err)
	return nil
}

func (repository *MerchantRepositoryImpl) FindMerchantIdByUserId(ctx context.Context, tx *sql.Tx, userId int) string {
	fmt.Println("Masuk sini")
	fmt.Println(userId, "user id nya ")
	var merchantId int
	SQL := "SELECT id from merchants WHERE user_id = $1 LIMIT 1"
	err := tx.QueryRowContext(ctx, SQL, userId).Scan(&merchantId)
	if err != nil {
		return ""
	}
	return strconv.Itoa(merchantId)
}
