package http

import (
	bm "github.com/DazzlingSun/monitorService/src/basic/net/http/blademaster"
	"github.com/DazzlingSun/monitorService/src/project/model"
)

func archives(c *bm.Context) {
	p := &model.ArchiveParam{}
	if err := c.Bind(p); err != nil {
		return
	}
	c.JSON(actSrv.Archives(c, p))
}
