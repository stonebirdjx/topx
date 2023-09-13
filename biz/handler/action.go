package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/stonebirdjx/topx/biz/config"
	"github.com/stonebirdjx/topx/biz/model"
	"github.com/stonebirdjx/topx/biz/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	acitonID = "actionid"
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
	Actions []*model.Action `json:"actions"`
}

type CreateActionsRes struct {
	Message string `json:"message"`
}

func (c *CreateActionsReq) validate(ctx context.Context) error {
	if len(c.Actions) == 0 {
		err := fmt.Errorf("create api actions can not be nil")
		hlog.CtxErrorf(ctx, "%s action len is zero, err=%s",
			util.GetLogID(ctx),
			err.Error(),
		)
		return err
	}

	for _, action := range c.Actions {
		if err := action.Validate(ctx); err != nil {
			hlog.CtxErrorf(ctx, "%s action validate args err=%s",
				util.GetLogID(ctx),
				err.Error(),
			)
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
			util.GetLogID(ctx),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	hlog.CtxTracef(ctx, "%s CreateActions request info=%+v",
		util.GetLogID(ctx),
		req,
	)

	if err := req.validate(ctx); err != nil {
		hlog.CtxErrorf(ctx, "%s CreateActions request body check err=%s",
			util.GetLogID(ctx),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	for idx, action := range req.Actions {
		if err := action.InsertOne(ctx); err != nil {
			hlog.CtxErrorf(ctx, "%s CreateActions Error writing to database element index=%d err=%s",
				util.GetLogID(ctx),
				idx,
				err.Error(),
			)
			sendError(c, errOption{statusCode: consts.StatusInternalServerError, err: err})
			return
		}
	}

	res := &CreateActionsRes{
		Message: "creat actions success",
	}

	sendOk(c, okOption{statusCode: consts.StatusCreated, obj: res})
}

type DeleteActionsReq struct {
	IDs []primitive.ObjectID `json:"ids"`
}

type DeleteActionsRes struct {
	Message string `json:"message"`
}

// DeleteActions .
func DeleteActions(ctx context.Context, c *app.RequestContext) {
	req := &DeleteActionsReq{}
	if err := c.BindAndValidate(req); err != nil {
		hlog.CtxErrorf(ctx, "%s DeleteActions BindAndValidate request err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	if err := model.DeleteByIDs(ctx, req.IDs); err != nil {
		hlog.CtxErrorf(ctx, "%s DeleteActions delte mongo err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	res := &DeleteActionsRes{
		Message: "delete success",
	}

	sendOk(c, okOption{statusCode: consts.StatusOK, obj: res})
}

type GetActionRes struct {
	*model.Action
}

type GetActionResNoDocuments struct {
	Message string `json:"message"`
}

// GetAction return an acitons.
func GetAction(ctx context.Context, c *app.RequestContext) {
	acitonID := c.Param(acitonID)
	id, err := primitive.ObjectIDFromHex(acitonID)
	if err != nil {
		hlog.CtxErrorf(ctx, "%s GetAction get objectid err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}
	action := &model.Action{
		ID: id,
	}

	err = action.GetAction(ctx)
	if err != nil {
		getActionErrHandler(ctx, c, err)
		return
	}
	res := GetActionRes{
		Action: action,
	}

	sendOk(c, okOption{statusCode: consts.StatusOK, obj: res})
}

func getActionErrHandler(ctx context.Context, c *app.RequestContext, err error) {
	if err != mongo.ErrNoDocuments {
		hlog.CtxErrorf(ctx, "%s GetAction query mongo err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	res := &GetActionResNoDocuments{
		Message: "the action no document",
	}
	sendOk(c, okOption{statusCode: consts.StatusOK, obj: res})
}

type UpdateActionReq struct {
	model.Action
}

func (u *UpdateActionReq) validate(ctx context.Context) error {
	return u.Action.Validate(ctx)
}

type UpdateActionRes struct {
	Message string `json:"message"`
}

// UpdateAction update aciton information.
func UpdateAction(ctx context.Context, c *app.RequestContext) {
	acitonID := c.Param(acitonID)
	id, err := primitive.ObjectIDFromHex(acitonID)
	if err != nil {
		hlog.CtxErrorf(ctx, "%s UpdateAction get objectid err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	req := &UpdateActionReq{}
	if err := c.BindAndValidate(req); err != nil {
		hlog.CtxErrorf(ctx, "%s UpdateAction BindAndValidate request err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	req.ID = id

	if err := req.validate(ctx); err != nil {
		hlog.CtxErrorf(ctx, "%s UpdateAction update mongo err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	if err := req.Action.UpdateAction(ctx); err != nil {
		hlog.CtxErrorf(ctx, "%s UpdateAction update mongo err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	res := &UpdateActionRes{
		Message: "update success",
	}

	sendOk(c, okOption{statusCode: consts.StatusOK, obj: res})
}

// DeleteAction.
func DeleteAction(ctx context.Context, c *app.RequestContext) {
	acitonID := c.Param(acitonID)
	id, err := primitive.ObjectIDFromHex(acitonID)
	if err != nil {
		hlog.CtxErrorf(ctx, "%s DeleteAction get objectid err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	if err := model.DeleteByID(ctx, id); err != nil {
		hlog.CtxErrorf(ctx, "%s DeleteAction delete mongo err=%s",
			c.Response.Header.Get(config.RequestID),
			err.Error(),
		)
		sendError(c, errOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	res := &DeleteActionsRes{
		Message: "delete success",
	}

	sendOk(c, okOption{statusCode: consts.StatusOK, obj: res})
}
