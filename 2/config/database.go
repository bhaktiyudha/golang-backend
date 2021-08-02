package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&timeout=5000ms",
		db_username, db_password, db_host, db_port, db_database))

	if err != nil {
		log.Fatal("Error")
	}

	return db, err
}
