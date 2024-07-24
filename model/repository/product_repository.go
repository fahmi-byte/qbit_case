package repository

import (
	"context"
	"database/sql"
	"qbit_case/model/domain"
	"qbit_case/model/web"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, tx *sql.Tx, request domain.Product) error
	GetAllProducts(ctx context.Context, tx *sql.Tx, params web.ProductParams) ([]domain.Product, error)
}
