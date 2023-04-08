package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/NaylaDeLis/Go-6-Database/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repo.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(lastID)
	return comment, nil
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id=?"
	comment := entity.Comment{}

	err := repo.DB.QueryRowContext(ctx, script, int64(id)).Scan(&comment.Id, &comment.Email, &comment.Comment)
	if err == sql.ErrNoRows {
		return comment, errors.New("Comment with id: " + strconv.Itoa(int(id)) + " are not found")
	} else if err != nil {
		return comment, err
	} else {
		return comment, err
	}
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT * FROM comments"

	var arrComment []entity.Comment

	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return arrComment, err
	}

	for rows.Next() {
		var commentEntity entity.Comment
		err := rows.Scan(&commentEntity.Id, &commentEntity.Email, &commentEntity.Comment)
		if err != nil {
			return arrComment, err
		}

		arrComment = append(arrComment, commentEntity)
	}

	return arrComment, nil
}
