package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"qbit_case/config"
	"qbit_case/helper"
)

func NewDB(config config.Database) *sql.DB {
	db, err := sql.Open(config.Driver, config.GetDataSourceName())
	helper.PanicIfError(err)

	db.SetMaxIdleConns(config.MaxIdleConnection)
	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetConnMaxLifetime(config.GetConnectionMaxLifetime())
	db.SetConnMaxIdleTime(config.GetConnectionMaxIdleTime())

	return db
}
