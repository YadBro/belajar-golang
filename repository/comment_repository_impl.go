package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImplementation struct {
	*sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImplementation{DB: db}
}

func (c *commentRepositoryImplementation) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := c.DB.ExecContext(ctx, script, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (c *commentRepositoryImplementation) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id=? LIMIT 1"

	rows, err := c.DB.QueryContext(ctx, script, id)

	comment := entity.Comment{}

	if err != nil {
		return comment, err
	}

	defer rows.Close()
	if rows.Next() {
		// ada
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		if err != nil {
			return comment, err
		}

		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found!")
	}
}

func (c *commentRepositoryImplementation) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"

	rows, err := c.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []entity.Comment

	for rows.Next() {
		comment := entity.Comment{}
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		if err != nil {
			return comments, err
		}

		comments = append(comments, comment)

		return comments, nil
	}

	return comments, nil
}
