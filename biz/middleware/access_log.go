package middleware

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/stonebirdjx/topx/biz/util"
)

func AccessLog(ctx context.Context, c *app.RequestContext) {
	start := time.Now()
	c.Next(ctx)
	latency := time.Since(start).Milliseconds()
	hlog.CtxTracef(ctx, "%s status=%d cost=%dms method=%s full_path=%s client_ip=%s host=%s",
		util.GetLogID(ctx),
		c.Response.StatusCode(),
		latency,
		c.Request.Header.Method(),
		c.Request.URI().PathOriginal(),
		c.ClientIP(),
		c.Request.Host(),
	)
}
