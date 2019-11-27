package aqm

import (
	"context"

	"github.com/DazzlingSun/monitorService/src/basic/container/queue/aqm"
	"github.com/DazzlingSun/monitorService/src/basic/ecode"
	bm "github.com/DazzlingSun/monitorService/src/basic/net/http/blademaster"
	"github.com/DazzlingSun/monitorService/src/basic/rate"
	"github.com/DazzlingSun/monitorService/src/basic/rate/limit"
	"github.com/DazzlingSun/monitorService/src/basic/stat/prom"
)

const (
	_family = "blademaster"
)

var (
	stats = prom.New().WithState("go_active_queue_mng", []string{"family", "title"})
)

// AQM aqm midleware.
type AQM struct {
	limiter rate.Limiter
}

// New return a ratelimit midleware.
func New(conf *aqm.Config) (s *AQM) {
	return &AQM{
		limiter: limit.New(conf),
	}
}

// Limit return a bm handler func.
func (a *AQM) Limit() bm.HandlerFunc {
	return func(c *bm.Context) {
		done, err := a.limiter.Allow(c)
		if err != nil {
			stats.Incr(_family, c.Request.URL.Path[1:])
			// TODO: priority request.
			// c.JSON(nil, err)
			// c.Abort()
			return
		}
		defer func() {
			if c.Error != nil && !ecode.Deadline.Equal(c.Error) && c.Err() != context.DeadlineExceeded {
				done(rate.Ignore)
				return
			}
			done(rate.Success)
		}()
		c.Next()
	}
}
