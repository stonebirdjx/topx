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
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stonebirdjx/topx/biz/utils"
)

// Metrics .
func Metrics(ctx context.Context, c *app.RequestContext) {
	h := promhttp.Handler()
	req, err := adaptor.GetCompatRequest(&c.Request)
	if err != nil {
		hlog.CtxTracef(ctx, "%s hertz adaptor http request err=%s",
			utils.GetLogID(ctx),
			err.Error(),
		)
		return
	}

	// caution: don't pass in c.GetResponse() as it return a copy of response
	writer := adaptor.GetCompatResponseWriter(&c.Response)
	h.ServeHTTP(writer, req)
}
