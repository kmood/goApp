package utils

import (
	"github.com/prometheus/common/log"
	"testing"
)

func TestGetDB(t *testing.T) {
	db, i := GetDB()
	if i != nil {
		log.Error(i)
	}
	defer  db.Close()
	rows, e := db.Query("select * from BWJL_GXXX")
	if e != nil {
		log.Error(e)
	}
	strings, _ := rows.Columns()
	println(strings)
	//println(ping)
}