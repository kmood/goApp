package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDB() {
	sql.Open()
}
