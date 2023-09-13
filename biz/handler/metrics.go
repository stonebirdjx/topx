package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stonebirdjx/topx/biz/util"
)

// Metrics .
func Metrics(ctx context.Context, c *app.RequestContext) {
	h := promhttp.Handler()
	req, err := adaptor.GetCompatRequest(&c.Request)
	if err != nil {
		hlog.CtxTracef(ctx, "%s hertz adaptor http request err=%s",
			util.GetLogID(ctx),
			err.Error(),
		)
		return
	}

	// caution: don't pass in c.GetResponse() as it return a copy of response
	writer := adaptor.GetCompatResponseWriter(&c.Response)
	h.ServeHTTP(writer, req)
}
