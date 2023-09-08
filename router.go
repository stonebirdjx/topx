// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/requestid"
	handler "github.com/stonebirdjx/topx/biz/handler"
	"github.com/stonebirdjx/topx/biz/middleware"
)

// RESTfulAPi 参考kubernetes: https://kubernetes.io/zh-cn/docs/reference/kubernetes-api/
const (
	apiv1 = "/apis/v1"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {

	// middleware
	r.Use(
		requestid.New(),
		middleware.AccessLog,
		middleware.RetaLimit,
	)

	r.GET("/ping", handler.Ping)
	r.GET("/metrics", handler.Metrics)
	// your code ...

	apiv1Register(r)
}

func apiv1Register(r *server.Hertz) {
	g := r.Group(apiv1)
	g.GET("/actions", handler.ListActions)
	g.POST("/actions", handler.CreateAction)
	g.GET("/actions/:actionid", handler.GetAction)
	g.PATCH("/actions/:actionid", handler.UpdateAction)
	g.PUT("/actions/:actionid", handler.UpdateAction)
	g.DELETE("/actions/:actionid", handler.DeleteAction)
}
