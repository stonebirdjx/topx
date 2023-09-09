package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/stonebirdjx/topx/biz/config"
	"github.com/stonebirdjx/topx/biz/model"
)

// ListActions return all acitons.
func ListActions(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
}

// GetAction return an acitons.
func GetAction(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
}

type CreateActionsReq struct {
	Actions []model.Action `json:"actions"`
}

type CreateActionsRes struct {
	Message string `json:"message"`
}

func (c *CreateActionsReq) validate() error {
	if len(c.Actions) == 0 {
		return fmt.Errorf("create api actions can not be nil")
	}

	for _, action := range c.Actions {
		if err := action.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// CreateActions create some new acitons.
func CreateActions(ctx context.Context, c *app.RequestContext) {
	req := &CreateActionsReq{}
	if err := c.BindAndValidate(req); err != nil {
		hlog.CtxErrorf(ctx, "%s CreateActions BindAndValidate request err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	hlog.CtxTracef(ctx, "%s CreateActions request info=%+v",
		c.Response.Header.Get(config.RequestID),
		req,
	)

	if err := req.validate(); err != nil {
		hlog.CtxErrorf(ctx, "%s CreateActions request body check err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	for idx, action := range req.Actions {
		if err := action.InsertOne(ctx); err != nil {
			hlog.CtxErrorf(ctx, "%s CreateActions Error writing to database element index=%d err=%s",
				c.Response.Header.Get(config.RequestID),
				idx,
				err.Error(),
			)
			sendError(c, errOption{statusCode: consts.StatusInternalServerError, err: err})
			return
		}
	}

	res := &CreateActionsRes{
		Message: "creat some new actions success",
	}

	sendOk(c, okOption{statusCode: consts.StatusCreated, obj: res})
}

// UpdateAction  update aciton information.
func UpdateAction(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
}

// DeleteAction delete a aciton.
func DeleteAction(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
}
