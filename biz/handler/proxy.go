// Copyright 2023 The Author stonebird. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/go-redis/redis_rate/v10"
	"github.com/stonebirdjx/topx/biz/dal"
	"github.com/stonebirdjx/topx/biz/utils"
)

const (
	serviceName = "serviceName"
	version     = "version"
	actionName  = "actionName"
)

func (ctrl *Controller) Porxy(ctx context.Context, c *app.RequestContext) {
	serviceName := c.Param(serviceName)
	version := c.Param(version)
	actionName := c.Param(actionName)

	hlog.CtxInfof(ctx, "%s Porxy service_name=%s version=%s action_name=%s",
		utils.GetLogID(ctx),
		serviceName,
		version,
		actionName,
	)

	action, err := ctrl.daler.FindActionByOpt(ctx, dal.FindActionOption{
		ServiceName: serviceName,
		Version:     version,
		ActionName:  actionName,
	})

	if err != nil {
		hlog.CtxErrorf(ctx, "%s Porxy query mongo err=%s",
			utils.GetLogID(ctx),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	res, err := ctrl.daler.Allow(ctx, "proxy:"+actionName, redis_rate.PerMinute(int(action.RateLimit)))
	if err != nil {
		hlog.CtxErrorf(ctx, "%s Porxy query redis limit err=%s",
			utils.GetLogID(ctx),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}
	if res.Allowed == 0 {
		hlog.CtxTracef(ctx, "%s Porxy rate limited.",
			utils.GetLogID(ctx),
		)
		sendOK(ctx, c, okOption{
			statusCode: consts.StatusTooManyRequests,
			obj: map[string]string{
				"message": "too many request",
			},
		})
		return
	}

	proxyReq, porxyRes := &protocol.Request{}, &protocol.Response{}
	c.Request.CopyTo(proxyReq)
	proxyReq.URI().SetScheme(string(action.Proxy[0].Scheme))
	proxyReq.URI().SetHost(action.Proxy[0].Host)
	proxyReq.URI().SetPath(action.Proxy[0].Path)
	hlog.CtxInfof(ctx, "%s proxy uri=%s",
		utils.GetLogID(ctx),
		proxyReq.URI().String(),
	)

	switch action.Proxy[0].Scheme {
	case "http":
		err = utils.HzHTTPClient.Do(ctx, proxyReq, porxyRes)
	case "https":
		err = utils.HzHTTPSClient.Do(ctx, proxyReq, porxyRes)
	}

	if err != nil {
		hlog.CtxErrorf(ctx, "%s Porxy request err=%s",
			utils.GetLogID(ctx),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	porxyRes.CopyTo(&c.Response)
}
