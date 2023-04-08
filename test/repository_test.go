package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/NaylaDeLis/Go-6-Database/entity"
	"github.com/NaylaDeLis/Go-6-Database/repository"
	"github.com/NaylaDeLis/Go-6-Database/services"
)

func TestCommentInsert(t *testing.T) {
	db := services.GetConnection()
	commentRepository := repository.NewCommentRepository(db)

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "joko11@gmail.com",
		Comment: "Lorem Ipsum",
	}

	comment, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestCommentFindByID(t *testing.T) {
	db := services.GetConnection()
	commentRepository := repository.NewCommentRepository(db)

	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 11)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	db := services.GetConnection()
	commentRepository := repository.NewCommentRepository(db)

	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
