package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"golang.org/x/time/rate"
)

type LimiterManager struct {
	limiter *rate.Limiter
}

type LimiterOption struct {
	R rate.Limit
	B int
}

func NewLimiter(ctx context.Context, l LimiterOption) *LimiterManager {
	limiter := rate.NewLimiter(l.R, l.B)
	return &LimiterManager{
		limiter: limiter,
	}
}

// RetaLimit 全局速率中间件.
func RetaLimit(ctx context.Context, c *app.RequestContext) {
	c.Next(ctx)
}
