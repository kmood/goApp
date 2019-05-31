package dao

import (
	"context"
	"library/conf"

	//account "go-common/app/service/main/account/api"
	//archive "go-common/app/service/main/archive/api"

	"library/database/orm"
	"github.com/jinzhu/gorm"
)

// Dao is the appeal database access object
type Dao struct {
	ORM     *gorm.DB
}

// New will create a new appeal Dao instance
func New() (d *Dao) {
	config:= conf.GetConfig()
	d = &Dao{
		ORM:       orm.NewMySQL(config.Orm),
	}
	d.initORM()
	return
}

func (d *Dao) initORM() {
	d.ORM.LogMode(true)
}

// Close close dao.
func (d *Dao) Close() {
	if d.ORM != nil {
		d.ORM.Close()
	}
}

// Ping ping cpdb
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.ORM.DB().PingContext(c); err != nil {
		return
	}
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
