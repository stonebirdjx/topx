// Copyright 2023 The Author stonebird. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/requestid"
	handler "github.com/stonebirdjx/topx/biz/handler"
	"github.com/stonebirdjx/topx/biz/middleware"
)

// RESTfulAPi 参考kubernetes: https://kubernetes.io/zh-cn/docs/reference/kubernetes-api/
const (
	appv1   = "/apis/apps/v1"
	proxyv1 = "/apis/proxy/v1"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {

	// middleware
	r.Use(
		requestid.New(),
		middleware.SetLogID,
		middleware.AccessLog,
		middleware.RetaLimit,
	)

	r.GET("/ping", handler.Ping)
	r.GET("/metrics", handler.Metrics)
	// your code ...

	appv1Register(r)
	proxyv1Register(r)
}

func proxyv1Register(r *server.Hertz) {
	g := r.Group(proxyv1)
	//	http://ip:port/apis/proxy/v1/iva/2022-05-13/TestPing
	g.Any("/:serviceName/:version/:actionName", handler.Porxy)
}

func appv1Register(r *server.Hertz) {
	g := r.Group(appv1)
	g.POST("/actions", handler.CreateActions)
	g.GET("/actions", handler.ListActions)
	g.DELETE("/actions", handler.DeleteActions)
	g.GET("/actions/:actionid", handler.GetAction)
	g.PATCH("/actions/:actionid", handler.UpdateAction)
	g.PUT("/actions/:actionid", handler.UpdateAction)
	g.DELETE("/actions/:actionid", handler.DeleteAction)
}
