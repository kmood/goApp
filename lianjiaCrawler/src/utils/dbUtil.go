package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3307)/test?charset=utf8")
	return db, err
}
