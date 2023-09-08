package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/stonebirdjx/topx/biz/config"
	"golang.org/x/time/rate"
)

var limiter *rate.Limiter

type LimiterOptions struct {
	R rate.Limit
	B int
}

func NewLimiter(ctx context.Context, l LimiterOptions) {
	limiter = rate.NewLimiter(l.R, l.B)
}

// RetaLimit 全局速率中间件.
func RetaLimit(ctx context.Context, c *app.RequestContext) {
	total := limiter.Limit()
	tokens := limiter.Tokens()
	hlog.CtxTracef(ctx, "%s system_total_rate=%fQPS sytem_now_rate=%fQPS",
		c.Response.Header.Get(config.RequestID),
		total,
		total-rate.Limit(tokens),
	)

	if !limiter.Allow() {
		// 429 StatusTooManyRequests.
		c.AbortWithMsg("Request rate limit", consts.StatusTooManyRequests)
	}
	c.Next(ctx)
}
