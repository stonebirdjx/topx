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

// Package internal Save some status during system initialization
// only for handler paclage use.
package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/stonebirdjx/topx/biz/config"
	"github.com/stonebirdjx/topx/biz/dal"
	"golang.org/x/time/rate"
)

// Controller app global variable controller.
// Dependency injection instead of global variables
type Controller struct {
	daler   dal.Daler
	limiter *rate.Limiter
}

func NewController() (*Controller, error) {
	cfg, err := config.InitConfiger()
	if err != nil {
		return nil, err
	}

	daler, err := dal.NewDaler(dal.DalerOption{})
	if err != nil {
		return nil, err
	}

	limiter := rate.NewLimiter(rate.Limit(cfg.GetRateLimte()), cfg.GetBurst())
	return &Controller{
		limiter: limiter,
		daler:   daler,
	}, nil
}

type errOption struct {
	statusCode int
	err        error
}

func sendError(c *app.RequestContext, opt errOption) {
	c.JSON(opt.statusCode, utils.H{
		"message": opt.err.Error(),
	})
}

type okOption struct {
	statusCode int
	obj        any
}

func sendOk(c *app.RequestContext, opt okOption) {
	c.JSON(opt.statusCode, opt.obj)
}
