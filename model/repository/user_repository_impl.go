package repository

import (
	"context"
	"database/sql"
	"errors"
	"qbit_case/constant"
	"qbit_case/helper"
	"qbit_case/model/domain"
	"qbit_case/model/web"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) CreateUser(ctx context.Context, tx *sql.Tx, request web.RegisterRequest) (int, error) {
	var userId int
	SQL := "INSERT INTO users (username, email, password, phone_number, role_id, address, city, full_name) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	err := tx.QueryRowContext(ctx, SQL, request.Username, request.Email, request.Password, request.PhoneNumber, request.RoleId, request.Address, request.City, request.FullName).Scan(&userId)
	helper.PanicIfError(err)
	return userId, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	var user domain.User
	rows, err := tx.QueryContext(ctx, "SELECT u.id, email, username, password, full_name, phone_number, r.name as role FROM users u JOIN roles r ON u.role_id = r.id WHERE email=$1 LIMIT 1", email)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.FullName, &user.PhoneNumber, &user.Role)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("Wrong Email!")
	}
}

func (repository *UserRepositoryImpl) UpdateToMerchant(ctx context.Context, tx *sql.Tx, userId int) error {
	SQL := "UPDATE users SET role_id = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, SQL, constant.MERCHANT, userId)
	helper.PanicIfError(err)
	return nil
}
