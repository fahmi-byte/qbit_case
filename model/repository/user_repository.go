package repository

import (
	"context"
	"database/sql"
	"qbit_case/model/domain"
	"qbit_case/model/web"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx *sql.Tx, request web.RegisterRequest) (int, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	UpdateToMerchant(ctx context.Context, tx *sql.Tx, userId int) error
}
