package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "localhost:@tcp(localhost:3306)/belajar_golang_database")

	if err != nil {
		panic(err)
	}

	defer db.Close()

}
