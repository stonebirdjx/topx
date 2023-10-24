// Copyright 2023 The Author stonebird. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/stonebirdjx/topx/biz/config"
	"github.com/stonebirdjx/topx/biz/dal"
	"github.com/stonebirdjx/topx/biz/middleware"
	"github.com/stonebirdjx/topx/biz/utils"
	"golang.org/x/time/rate"
)

func Init() error {
	g := config.ReadFromEnv()
	if err := g.Validate(); err != nil {
		return err
	}

	ctx := context.Background()
	middleware.NewLimiter(ctx, middleware.LimiterOption{R: rate.Limit(g.RateLimit), B: g.Burst})

	if err := utils.ProxyClientHTTPInit(); err != nil {
		return err
	}

	if err := utils.ProxyClientHTTPSInit(); err != nil {
		return err
	}

	if err := dal.MongoInit(ctx, dal.MongoOption{URI: g.MongoDBURI, DB: g.MongoDBDB}); err != nil {
		return err
	}

	return dal.RedisInit(dal.RedisOption{URI: g.RedisURI})
}

func main() {
	hlog.Infof(utils.BlessProgram())
	if err := Init(); err != nil {
		panic(err.Error())
	}

	// hertz
	h := server.Default(server.WithHostPorts(":6789"))

	register(h)
	h.Spin()

	// hlog.Infof(config.Thankyou())
}
