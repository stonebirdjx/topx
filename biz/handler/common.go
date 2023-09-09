package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

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
