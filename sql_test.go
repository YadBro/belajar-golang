package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

var db *sql.DB
var ctx context.Context

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	db = GetConnection()

	ctx = context.Background()

	m.Run()

	db.Close()
	fmt.Println("AFTER UNIT TEST")
}

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('joko', 'Joko')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string

		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID:", id)
		fmt.Println("NAME:", name)
	}

}

func TestQuerySqlComplex(t *testing.T) {
	script := "SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("=================")
		fmt.Println("ID:", id)
		fmt.Println("NAME:", name)
		if email.Valid {
			fmt.Println("EMAIL:", email)
		}
		fmt.Println("BALANCE:", balance)
		fmt.Println("RATING:", rating)
		fmt.Println("CREATED_AT:", created_at)
		if birth_date.Valid {
			fmt.Println("BIRTH_DATE:", birth_date)
		}
		fmt.Println("MARRIED:", married)
		fmt.Println("=================\n\n.")
	}
}

func TestSqlInjection(t *testing.T) {
	var username, password string

	username = "admin'; #"
	password = "salah"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {

		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses login!", username)
	} else {
		fmt.Println("Gagal login!")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	var username, password string

	username = "admin"
	password = "admin"

	script := "SELECT username FROM user WHERE username=? AND password=? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {

		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses login!", username)
	} else {
		fmt.Println("Gagal login!")
	}
}

func TestExecSqlSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	username := "yadi'; DROP TABLE user; #"
	password := "yadi"

	ctx := context.Background()

	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	email := "jena@gmail.com"
	comment := "This is my jena comment."

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", id)
}

func TestPrepareStatement(t *testing.T) {
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i <= 10; i++ {
		email := "yadi" + strconv.Itoa(i) + "@gmail.com"
		comment := "Comment " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("New comment with id", id)
	}

	fmt.Println("Success")
}

func TestTransaction(t *testing.T) {
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	for i := 0; i < 10; i++ {
		email := "yadi" + strconv.Itoa(i) + "@gmail.com"
		comment := "Comment " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("New comment with id", id)
	}

	// err = tx.Commit()
	err = tx.Rollback()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Success")
}
