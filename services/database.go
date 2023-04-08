package services

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:Boardmaker^19@tcp(localhost:3306)/go_6_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 ^ time.Minute)
	db.SetConnMaxLifetime(50 * time.Minute)

	return db
}
