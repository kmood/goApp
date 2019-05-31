package sql

import (
	"github.com/prometheus/common/log"
	"testing"
)

func TestNewMySQL(t *testing.T) {
	config := &Config{
		Addr:        "test",
		DSN:         "root:123456@tcp(localhost:3307)/test?charset=utf8",
		Active:      10,
		Idle:        5,
		IdleTimeout :1000,
		TranTimeout : 1000,
		ExecTimeout : 1000,
		QueryTimeout : 1000,
	}
	db := NewMySQL(config)
	defer  db.Close()
	e := db.Ping()
	if e != nil {
		log.Error(e)
	}
	//println(ping)
}