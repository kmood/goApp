package orm

import (
	"github.com/prometheus/common/log"
	"time"

	// database driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Config mysql config.
type Config struct {
	DSN         string         // data source name.
	Active      int            // pool
	Idle        int            // pool
	IdleTimeout int64 // connect max life time.
}

type ormLog struct{}


// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	gorm.AddNamingStrategy(&gorm.NamingStrategy{
		DB: func(name string) string {
			return "db_" + name
		},
		Table: func(name string) string {
			return "table_" + name
		},
		Column: func(name string) string {
			return "col_" + name
		},
	})
	db.AutoMigrate()
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	return
}

