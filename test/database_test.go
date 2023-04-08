package test

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/NaylaDeLis/Go-6-Database/services"

	_ "github.com/go-sql-driver/mysql"
)

// Setting a database pool
func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:Boardmaker^19@tcp(localhost:3306)/go_6_database")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(2 ^ time.Second)
	db.SetConnMaxLifetime(50 * time.Minute)

	fmt.Println(db)
}

// Query a SQL with no result
func TestExecSQL(t *testing.T) {
	db := services.GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO Customer(id,name) VALUES ('Eko', 'eko')")
	if err != nil {
		panic(err)
	}

	fmt.Println("Success inserting to Customer")
}

// Query a SQL with result and reading the result
func TestQuerySQL(t *testing.T) {
	db := services.GetConnection()
	defer db.Close()
	ctx := context.Background()

	queryCommand := "SELECT id, name, email, balance, rating, created_at, birth_date, married FROM Customer"

	rows, err := db.QueryContext(ctx, queryCommand)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int
		var rating float64
		var created_at time.Time
		var birth_date sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID	:", id)
		fmt.Println("Name	:", name)
		fmt.Println("Email	:", email.Valid)
		fmt.Println("Balance	:", balance)
		fmt.Println("Rating	:", rating)
		fmt.Println("createdAt:", created_at)
		fmt.Println("birthDate:", birth_date)
		fmt.Println("Married	:", married)
		fmt.Println("-----------------------------------------")
	}
}

// Preventing SQL injection
func TestSQLInjection(t *testing.T) {
	db := services.GetConnection()
	defer db.Close()

	username := "admin'; #"
	password := "salah"

	ctx := context.Background()
	queryCommand := "SELECT username, password FROM sqlinjection WHERE username=? LIMIT 1"

	var usernameRes string
	var passwordRes string

	err := db.QueryRowContext(ctx, queryCommand, username).Scan(&usernameRes, &passwordRes)
	if err == sql.ErrNoRows {
		fmt.Println("No Username was found")
	} else if passwordRes != password {
		fmt.Println("Password salah")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Login", username, "Success!")
	}
}

// Get the ID from last time querying
func TestLastID(t *testing.T) {
	db := services.GetConnection()
	defer db.Close()

	ctx := context.Background()
	queryCommand := "INSERT INTO comments(name, comment) VALUES(?, ?)"
	name := "Riverwater"
	comment := "There is no such things"
	result, err := db.ExecContext(ctx, queryCommand, name, comment)
	if err != nil {
		panic(err)
	}

	lastInsID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(lastInsID)

	DelID, err1 := db.ExecContext(ctx, "DELETE FROM comments WHERE id=?", lastInsID)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(DelID.RowsAffected())
}

// Doing query for several time with same command
func TestPrepareStatement(t *testing.T) {
	db := services.GetConnection()
	defer db.Close()

	ctx := context.Background()
	queryCommand := "INSERT INTO comments(email, comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, queryCommand)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "joko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Lorem Ipsum"

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		lastID, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println(lastID)
	}
}
