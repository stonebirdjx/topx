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
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/stonebirdjx/topx/biz/config"
	"github.com/stonebirdjx/topx/biz/dal"
	"github.com/stonebirdjx/topx/biz/utils"
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

	daler, err := dal.NewDaler(
		dal.DalerOption{
			MongoDBURI:  cfg.GetMongDBURI(),
			MongoDBName: cfg.GetMongDBName(),
			RedisURI:    cfg.GetRedisURI(),
		},
	)
	if err != nil {
		return nil, err
	}

	limiter := rate.NewLimiter(rate.Limit(cfg.GetRateLimte()), cfg.GetBurst())
	return &Controller{
		limiter: limiter,
		daler:   daler,
	}, nil
}

type BaseMsg struct {
	Message string `json:"message"`
}

type ResponseMetadata struct {
	RequestID string `json:"request_id"`
	Error     *Error `json:"error,omitempty"`
}

type Error struct {
	Code int `json:"code"`
	BaseMsg
}

type errorResponse struct {
	ResponseMetadata ResponseMetadata `json:"metadata"`
}

type errorOption struct {
	statusCode int
	err        error
}

func sendError(ctx context.Context, c *app.RequestContext, opt errorOption) {
	res := &errorResponse{
		ResponseMetadata: ResponseMetadata{
			RequestID: utils.GetLogID(ctx),
			Error: &Error{
				Code: opt.statusCode,
				BaseMsg: BaseMsg{
					Message: opt.err.Error(),
				},
			},
		},
	}
	c.JSON(opt.statusCode, res)
}

type okResponse struct {
	ResponseMetadata ResponseMetadata `json:"metadata"`
	Result           any              `json:"result"`
}
type okOption struct {
	statusCode int
	obj        any
}

func sendOK(ctx context.Context, c *app.RequestContext, opt okOption) {
	res := &okResponse{
		ResponseMetadata: ResponseMetadata{
			RequestID: utils.GetLogID(ctx),
		},
		Result: opt.obj,
	}
	c.JSON(opt.statusCode, res)
}
