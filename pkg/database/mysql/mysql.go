package mysql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	dsn := os.Getenv("DSN")

	db, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(5)
	return db, nil
}
