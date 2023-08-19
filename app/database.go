package app

import (
	"database/sql"
	"learn-7-restful-api/helper"
	"time"
)

func NewDB(driver string, dataSourceName string) *sql.DB {
	db, err := sql.Open(driver, dataSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
