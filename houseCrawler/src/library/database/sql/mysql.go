package sql

import (
	"database/sql"
	"github.com/prometheus/common/log"
	"time"
	_ "github.com/go-sql-driver/mysql"

)
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
	//Breaker      *breaker.Config // breaker
}


// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *sql.DB) {
	//if c.QueryTimeout == 0 || c.ExecTimeout == 0 || c.TranTimeout == 0 {
	//	panic("mysql must be set query/execute/transction timeout")
	//}
	db, err := connect(c,c.DSN)
	if err != nil {
		log.Error(err)
	}
	return db
}

func connect(c *Config, dataSourceName string) (*sql.DB, error) {
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	d.SetMaxOpenConns(c.Active)
	d.SetMaxIdleConns(c.Idle)
	d.SetConnMaxLifetime(time.Duration(c.IdleTimeout))
	return d, nil
}
