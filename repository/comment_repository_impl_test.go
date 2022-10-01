package repository

import (
	db "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var commentRepository CommentRepository
var ctx context.Context

func TestMain(m *testing.M) {
	commentRepository = NewCommentRepository(db.GetConnection())
	ctx = context.Background()

	fmt.Println("DATABASE CONNECTED")

	m.Run()

	fmt.Println("DATABASE CLOSED")
}

func TestCommentInsert(t *testing.T) {

	comment := entity.Comment{
		Email:   "repository3@test.com",
		Comment: "Test Repository3",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	comment, err := commentRepository.FindById(ctx, 39)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
	fmt.Println("Id", comment.Id)
	fmt.Println("Email", comment.Email)
	fmt.Println("Comment", comment.Comment)
}

func TestFindAll(t *testing.T) {
	comments, err := commentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println(comments)

	for key, val := range comments {
		fmt.Println("KEY", key)
		fmt.Println("Id", val.Id)
		fmt.Println("Email", val.Email)
		fmt.Println("Comment", val.Comment)
	}
}
