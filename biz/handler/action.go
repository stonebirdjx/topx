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

type ListActionsReq struct {
	PapeSize int64 `query:"page_size"`
	PageNum  int64 `query:"page_num"`
}

func (l *ListActionsReq) validate() error {
	switch {
	case l.PapeSize < 0:
		return fmt.Errorf("page_size value can not lt 0")
	case l.PapeSize > 0:
		if l.PageNum < 1 {
			return fmt.Errorf("when page_size gt 0 page_num value can not lt 1")
		}
	}

	return nil

}

type ListActionsRes struct {
	Actions []model.Action `json:"actions"`
	Totals  int            `json:"totals"`
}

// ListActions return all acitons.
func ListActions(ctx context.Context, c *app.RequestContext) {
	req := &ListActionsReq{}
	if err := c.BindAndValidate(req); err != nil {
		hlog.CtxErrorf(ctx, "%s ListActions BindAndValidate request err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	hlog.CtxTracef(ctx, "%s ListActions request info=%+v",
		c.Response.Header.Get(config.RequestID),
		req,
	)

	if err := req.validate(); err != nil {
		hlog.CtxErrorf(ctx, "%s ListActions request body check err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	actions, total, err := model.ListActions(ctx, model.ListOption{
		PapeSize: req.PapeSize,
		PageNum:  req.PageNum,
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "%s ListActions query mongo err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	res := &ListActionsRes{
		Actions: *actions,
		Totals:  total,
	}

	sendOk(c, okOption{statusCode: consts.StatusOK, obj: res})
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

//
func DeleteActions(ctx context.Context, c *app.RequestContext) {

}

// GetAction return an acitons.
func GetAction(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
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
