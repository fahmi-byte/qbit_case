package repository

import (
	"context"
	"database/sql"
	"fmt"
	"qbit_case/helper"
	"qbit_case/model/domain"
	"qbit_case/model/web"
	"strings"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) CreateProduct(ctx context.Context, tx *sql.Tx, request domain.Product) error {
	SQL := "INSERT INTO products (merchant_id, name, description, price, final_price, discount, stock, category_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := tx.ExecContext(ctx, SQL, request.MerchantId, request.Name, request.Description, request.Price, request.FinalPrice, request.Discount, request.Stock, request.CategoryId)
	helper.PanicIfError(err)
	return nil
}

func (repository *ProductRepositoryImpl) GetAllProducts(ctx context.Context, tx *sql.Tx, params web.ProductParams) ([]domain.Product, error) {
	SQL := "SELECT p.id, m.name as merchant_name, m.city as merchant_city, p.name, p.description, price, final_price, discount, stock, c.name as category_name FROM products p " +
		"JOIN merchants m ON p.merchant_id = m.id " +
		"JOIN categories c ON p.category_id = c.id"

	var conditions []string
	var args []interface{}

	if params.Name != "" {
		conditions = append(conditions, "p.name ILIKE $1")
		args = append(args, "%"+params.Name+"%")
	}
	if params.City != "" {
		conditions = append(conditions, "m.city = $2")
		args = append(args, params.City)
	}
	if params.StartPrice > 0 {
		conditions = append(conditions, "p.final_price >= $3")
		args = append(args, params.StartPrice)
	}
	if params.EndPrice > 0 {
		conditions = append(conditions, "p.final_price <= $4")
		args = append(args, params.EndPrice)
	}

	if len(conditions) > 0 {
		SQL += " WHERE " + strings.Join(conditions, " AND ")
	}

	fmt.Println(SQL)

	rows, err := tx.QueryContext(ctx, SQL, args...)
	helper.PanicIfError(err)
	fmt.Println("Errornya disini")
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.MerchantName, &product.MerchantCity, &product.Name, &product.Description, &product.Price, &product.FinalPrice, &product.Discount, &product.Stock, &product.CategoryName); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
