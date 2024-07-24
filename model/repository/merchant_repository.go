package repository

import (
	"context"
	"database/sql"
	"qbit_case/model/web"
)

type MerchantRepository interface {
	CreateMerchant(ctx context.Context, tx *sql.Tx, request web.RegisterMerchantRequest) error
	FindMerchantIdByUserId(ctx context.Context, tx *sql.Tx, userId int) string
}
