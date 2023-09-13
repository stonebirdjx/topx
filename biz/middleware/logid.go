package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/stonebirdjx/topx/biz/config"
)

func SetLogID(ctx context.Context, c *app.RequestContext) {
	logCtx := context.WithValue(ctx, config.Key(config.LogID), c.Response.Header.Get(config.RequestID))
	c.Next(logCtx)
}
