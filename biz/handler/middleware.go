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
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/stonebirdjx/topx/biz/config"
	"github.com/stonebirdjx/topx/biz/utils"
	"golang.org/x/time/rate"
)

func SetLogID(ctx context.Context, c *app.RequestContext) {
	logCtx := context.WithValue(ctx, config.Key(config.LogID), c.Response.Header.Get(config.RequestID))
	c.Next(logCtx)
}

func AccessLog(ctx context.Context, c *app.RequestContext) {
	start := time.Now()
	c.Next(ctx)
	latency := time.Since(start).Milliseconds()
	hlog.CtxTracef(ctx, "%s status=%d cost=%dms method=%s full_path=%s client_ip=%s host=%s",
		utils.GetLogID(ctx),
		c.Response.StatusCode(),
		latency,
		c.Request.Header.Method(),
		c.Request.URI().PathOriginal(),
		c.ClientIP(),
		c.Request.Host(),
	)
}

var limiter *rate.Limiter

type LimiterOption struct {
	R rate.Limit
	B int
}

func NewLimiter(ctx context.Context, l LimiterOption) {
	limiter = rate.NewLimiter(l.R, l.B)
}

// RetaLimit 全局速率中间件.
func RetaLimit(ctx context.Context, c *app.RequestContext) {
	total := limiter.Limit()
	tokens := limiter.Tokens()
	hlog.CtxTracef(ctx, "%s total_rate=%fQPS, now sytem_rate=%fQPS",
		utils.GetLogID(ctx),
		total,
		total-rate.Limit(tokens),
	)

	if !limiter.Allow() {
		// 429 StatusTooManyRequests.
		c.AbortWithMsg("Request rate limit", consts.StatusTooManyRequests)
	}
	c.Next(ctx)
}
