package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/go-redis/redis_rate/v10"
	"github.com/stonebirdjx/topx/biz/dal"
	"github.com/stonebirdjx/topx/biz/model"
	"github.com/stonebirdjx/topx/biz/util"
)

const (
	serviceName = "serviceName"
	version     = "version"
	actionName  = "actionName"
)

func Porxy(ctx context.Context, c *app.RequestContext) {
	serviceName := c.Param(serviceName)
	version := c.Param(version)
	actionName := c.Param(actionName)

	action := &model.Action{
		Name:        actionName,
		ServiceName: serviceName,
		Version:     version,
	}

	hlog.CtxInfof(ctx, "%s Porxy service_name=%s version=%s action_name=%s",
		util.GetLogID(ctx),
		serviceName,
		version,
		actionName,
	)

	if err := action.FindOne(ctx); err != nil {
		hlog.CtxErrorf(ctx, "%s Porxy query mongo err=%s",
			util.GetLogID(ctx),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	res, err := dal.RedisLimiter.Allow(ctx, "proxy:"+actionName, redis_rate.PerMinute(int(action.RateLimit)))
	if err != nil {
		hlog.CtxErrorf(ctx, "%s Porxy query redis limit err=%s",
			util.GetLogID(ctx),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	if res.Allowed == 0 {
		hlog.CtxTracef(ctx, "%s Porxy rate limited.",
			util.GetLogID(ctx),
		)
		sendOk(c, okOption{
			statusCode: consts.StatusTooManyRequests,
			obj: utils.H{
				"message": "too many request",
			},
		})
		return
	}

	proxyReq, porxyRes := &protocol.Request{}, &protocol.Response{}
	c.Request.CopyTo(proxyReq)
	proxyReq.URI().SetScheme(string(action.Scheme))
	proxyReq.URI().SetHost(action.Host)
	proxyReq.URI().SetPath(action.Path)
	hlog.CtxInfof(ctx, "%s proxy uri=%s",
		util.GetLogID(ctx),
		proxyReq.URI().String(),
	)

	switch action.Scheme {
	case "http":
		err = util.HzHTTPClient.Do(ctx, proxyReq, porxyRes)
	case "https":
		err = util.HzHTTPSClient.Do(ctx, proxyReq, porxyRes)
	}

	if err != nil {
		hlog.CtxErrorf(ctx, "%s Porxy request err=%s",
			util.GetLogID(ctx),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	porxyRes.CopyTo(&c.Response)
}
