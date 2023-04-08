package main

import (
	"context"
	"fmt"

	"github.com/NaylaDeLis/Go-6-Database/services"
)

func main() {
	db := services.GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO Customer(id,name) VALUES ('Eko', eko)")
	if err != nil {
		panic(err)
	}

	fmt.Println("Success inserting to Customer")
}
