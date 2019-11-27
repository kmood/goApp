package conf

import (
	"errors"
	"flag"

	"github.com/DazzlingSun/monitorService/src/basic/conf"
	"github.com/DazzlingSun/monitorService/src/basic/database/orm"
	"github.com/DazzlingSun/monitorService/src/basic/log"
	bm "github.com/DazzlingSun/monitorService/src/basic/net/http/blademaster"
	"github.com/DazzlingSun/monitorService/src/basic/net/http/blademaster/middleware/permit"
	"github.com/DazzlingSun/monitorService/src/basic/net/rpc"
	"github.com/DazzlingSun/monitorService/src/basic/net/rpc/warden"
	"github.com/DazzlingSun/monitorService/src/basic/net/trace"

	"github.com/BurntSushi/toml"
)

var (
	// config
	confPath string
	client   *conf.Client
	// Conf .
	Conf = &Config{}
)

// Config def.
type Config struct {
	Auth       *permit.Config
	HTTPServer *bm.ServerConfig
	HTTPClient *bm.ClientConfig
	ORM        *orm.Config
	Log        *log.Config
	Tracer     *trace.Config
	Host       *Host
	// tag rpc client
	TagRPC *rpc.ClientConfig
	//article rpc client
	ArticlrRPC *rpc.ClientConfig
	ArcClient  *warden.ClientConfig
	AccClient  *warden.ClientConfig
}

// Host remote host
type Host struct {
	API string
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func remote() (err error) {
	if client, err = conf.New(); err != nil {
		return
	}
	if err = load(); err != nil {
		return
	}
	go func() {
		for range client.Event() {
			log.Info("config reload")
			if load() != nil {
				log.Error("config reload error (%v)", err)
			}
		}
	}()
	return
}

func load() (err error) {
	var (
		s       string
		ok      bool
		tmpConf *Config
	)
	if s, ok = client.Toml2(); !ok {
		return errors.New("load config center error")
	}
	if _, err = toml.Decode(s, &tmpConf); err != nil {
		return errors.New("could not decode config")
	}
	*Conf = *tmpConf
	return
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init int config
func Init() error {
	if confPath != "" {
		return local()
	}
	return remote()
}
