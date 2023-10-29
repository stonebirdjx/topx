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

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/requestid"
	handler "github.com/stonebirdjx/topx/biz/handler"
)

// RESTfulAPi 参考kubernetes: https://kubernetes.io/zh-cn/docs/reference/kubernetes-api/
const (
	appv1   = "/apis/apps/v1"
	proxyv1 = "/apis/proxy/v1"
)

// customizeRegister registers customize routers.
func customizedRegister(ctrl *handler.Controller, r *server.Hertz) {
	// middleware
	r.Use(
		requestid.New(),
		ctrl.SetLogID,
		ctrl.AccessLog,
		ctrl.RetaLimit,
	)

	r.GET("/ping", handler.Ping)
	r.GET("/metrics", handler.Metrics)

	// your code ...
	appv1Register(ctrl, r)
	proxyv1Register(ctrl, r)
}

func proxyv1Register(ctrl *handler.Controller, r *server.Hertz) {
	g := r.Group(proxyv1)
	//	http://ip:port/apis/proxy/v1/iva/2022-05-13/TestPing
	g.Any("/:serviceName/:version/:actionName", ctrl.Porxy)
}

func appv1Register(ctrl *handler.Controller, r *server.Hertz) {
	g := r.Group(appv1)
	g.POST("/actions", ctrl.CreateAction)
	g.GET("/actions", ctrl.ListActions)
	g.DELETE("/actions/:actionid", ctrl.DeleteAction)
	g.DELETE("/actions", ctrl.DeleteActions)
	g.GET("/actions/:actionid", ctrl.GetAction)
	g.PATCH("/actions/:actionid", ctrl.UpdateAction)
	g.PUT("/actions/:actionid", ctrl.UpdateAction)
}
