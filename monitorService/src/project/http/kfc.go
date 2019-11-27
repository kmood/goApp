package http

import (
	bm "github.com/DazzlingSun/monitorService/src/basic/net/http/blademaster"
	kfcmdl "github.com/DazzlingSun/monitorService/src/project/model/kfc"
)

func kfcList(c *bm.Context) {
	arg := new(kfcmdl.ListParams)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(kfcSrv.List(c, arg))
}
