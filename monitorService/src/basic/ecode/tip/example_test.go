package tip_test

import (
	"time"

	"github.com/DazzlingSun/monitorService/src/basic/ecode/tip"
	xhttp "github.com/DazzlingSun/monitorService/src/basic/net/http/blademaster"
	"github.com/DazzlingSun/monitorService/src/basic/net/netutil/breaker"
	xtime "github.com/DazzlingSun/monitorService/src/basic/time"
)

func ExampleInit() {
	conf := &tip.Config{
		Domain: "172.16.33.248:6401",
		Diff:   xtime.Duration(5 * time.Minute),
		ClientConfig: &xhttp.ClientConfig{
			App: &xhttp.App{
				Key:    "test",
				Secret: "e6c4c252dc7e3d8a90805eecd7c73396",
			},
			Dial:      xtime.Duration(time.Millisecond * 100),
			Timeout:   xtime.Duration(time.Second * 2),
			KeepAlive: xtime.Duration(time.Second * 2),
			Breaker: &breaker.Config{
				Window:  xtime.Duration(time.Millisecond * 10),
				Sleep:   xtime.Duration(time.Second * 10),
				Bucket:  10,
				Ratio:   0.5,
				Request: 100,
			},
		},
	}
	tip.Init(conf)
}
