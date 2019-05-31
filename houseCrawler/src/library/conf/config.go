package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
	"library/database/orm"
	csql "library/database/sql"
	clog "library/log"
	"path/filepath"
	"sync"
)
type  config struct {
	Orm          *orm.Config
	Mysql        *csql.Config
	Log 		 *clog.Config
}
var confPath string
var conf config
var once sync.Once
//单例
func GetConfig() config {
	once.Do(func() {
		_, e := toml.DecodeFile(confPath, &conf)
		if e != nil {
			//记录日志
			panic(e)
		}
	})
	return conf
}

func init() {
	flag.StringVar(&confPath, "conf", "", "配置文件路径")
	if confPath == "" {
		if cp_, e := filepath.Abs("../../app/config.toml");e == nil{
			confPath = cp_
		}
	}

}

