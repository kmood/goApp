package sql

import (
	"database/sql"
	"github.com/DazzlingSun/monitorService/src/basic/log"
	"github.com/DazzlingSun/monitorService/src/basic/net/netutil/breaker"
	"github.com/DazzlingSun/monitorService/src/basic/stat"
	"github.com/DazzlingSun/monitorService/src/basic/time"

	// database driver
	_ "github.com/go-sql-driver/mysql"
)

var stats = stat.DB

// Config mysql config.
type Config struct {
	Addr         string          // for trace
	DSN          string          // write data source name.
	ReadDSN      []string        // read data source name.
	Active       int             // pool
	Idle         int             // pool
	IdleTimeout  time.Duration   // connect max life time.
	QueryTimeout time.Duration   // query sql timeout
	ExecTimeout  time.Duration   // execute sql timeout
	TranTimeout  time.Duration   // transaction sql timeout
	Breaker      *breaker.Config // breaker
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *DB) {
	if c.QueryTimeout == 0 || c.ExecTimeout == 0 || c.TranTimeout == 0 {
		panic("mysql must be set query/execute/transction timeout")
	}
	db, err := Open(c)
	if err != nil {
		log.Error("open mysql error(%v)", err)
		panic(err)
	}
	return
}
