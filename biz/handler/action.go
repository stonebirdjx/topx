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
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/stonebirdjx/topx/biz/dal"
	"github.com/stonebirdjx/topx/biz/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	acitonID = "actionid"
)

type CreateActionRequest struct {
	*dal.Action
}

type CreateActionResponse struct {
	BaseMsg
}

// CreateAction create new action request.
func (ctrl *Controller) CreateAction(ctx context.Context, c *app.RequestContext) {
	req := &CreateActionRequest{}
	if err := c.BindAndValidate(req); err != nil {
		hlog.CtxErrorf(ctx, `%s CreateActions BindAndValidate request parameters err="%s"`,
			utils.GetLogID(ctx),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	hlog.CtxTracef(ctx, `%s CreateActions BindAndValidate request parameters="%v"`,
		utils.GetLogID(ctx),
		req.Action,
	)

	fmt.Printf("111%v", req.Action)
	if err := req.Action.Validate(); err != nil {
		hlog.CtxErrorf(ctx, `%s CreateActions Validate request parameters err="%s"`,
			utils.GetLogID(ctx),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	fmt.Printf("daler %v", ctrl.daler)

	if err := ctrl.daler.CreateAction(ctx, req.Action); err != nil {
		hlog.CtxErrorf(ctx, `%s CreateActions Request daler failed err="%s"`,
			utils.GetLogID(ctx),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	res := &CreateActionResponse{
		BaseMsg: BaseMsg{
			Message: "CreateActions create new action successfully",
		},
	}

	sendOK(ctx, c, okOption{statusCode: consts.StatusOK, obj: res})
}

type ListActionsRequest struct {
	PapeSize int64 `query:"page_size"`
	PageNum  int64 `query:"page_num"`
}

func (l *ListActionsRequest) validate() error {
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

type ListActionsResponse struct {
	Actions *[]dal.Action `json:"actions"`
	Totals  int           `json:"totals"`
}

// ListActions list some actions.
func (ctrl *Controller) ListActions(ctx context.Context, c *app.RequestContext) {
	req := &ListActionsRequest{}
	if err := c.BindAndValidate(req); err != nil {
		hlog.CtxErrorf(ctx, `%s ListActions BindAndValidate request parameters err="%s"`,
			utils.GetLogID(ctx),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	if err := req.validate(); err != nil {
		hlog.CtxErrorf(ctx, `%s ListActions validate request parameters err="%s"`,
			utils.GetLogID(ctx),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	hlog.CtxTracef(ctx, `%s ListActions Request parameters="%#v"`, req)

	actions, counter, err := ctrl.daler.ListActions(ctx, dal.ListActionOption{PapeSize: req.PapeSize, PageNum: req.PageNum})
	if err != nil {
		hlog.CtxErrorf(ctx, `%s ListActions request daler err="%s"`,
			utils.GetLogID(ctx),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	res := &ListActionsResponse{
		Actions: actions,
		Totals:  counter,
	}

	sendOK(ctx, c, okOption{statusCode: consts.StatusOK, obj: res})
}

type DeleteActionResponse struct {
	BaseMsg
}

// DeleteAction
func (ctrl *Controller) DeleteAction(ctx context.Context, c *app.RequestContext) {
	acitonID := c.Param(acitonID)
	id, err := primitive.ObjectIDFromHex(acitonID)
	if err != nil {
		hlog.CtxErrorf(ctx, `%s DeleteAction get primitive objectid err="%s"`,
			c.Response.Header.Get(utils.RequestID),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	hlog.CtxTracef(ctx, `%s DeleteAction actionID="%s"`, id)

	if err := ctrl.daler.DeleteActionByID(ctx, id); err != nil {
		hlog.CtxErrorf(ctx, `%s DeleteAction delete action err="%s"`,
			c.Response.Header.Get(utils.RequestID),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	res := &DeleteActionResponse{
		BaseMsg: BaseMsg{
			Message: "delete action success",
		},
	}

	sendOK(ctx, c, okOption{statusCode: consts.StatusOK, obj: res})
}

type DeleteActionsRequest struct {
	IDs []primitive.ObjectID `json:"ids"`
}

type DeleteActionsResponse struct {
	BaseMsg
}

// DeleteActions
func (ctrl *Controller) DeleteActions(ctx context.Context, c *app.RequestContext) {
	req := &DeleteActionsRequest{}
	if err := c.BindAndValidate(req); err != nil {
		hlog.CtxErrorf(ctx, `%s DeleteActions BindAndValidate request parameters err="%s"`,
			utils.GetLogID(ctx),
			err.Error(),
		)
		return
	}

	hlog.CtxTracef(ctx, `%s DeleteActions actionIDs="%+v"`, utils.GetLogID(ctx), req.IDs)

	if err := ctrl.daler.DeleteActions(ctx, req.IDs); err != nil {
		hlog.CtxErrorf(ctx, `%s DeleteActions delete actions err="%s"`,
			utils.GetLogID(ctx),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	res := &DeleteActionsResponse{
		BaseMsg: BaseMsg{
			Message: "delete action success",
		},
	}

	sendOK(ctx, c, okOption{statusCode: consts.StatusOK, obj: res})
}

type GetActionResponse struct {
	*dal.Action
}

// GetAction
func (ctrl *Controller) GetAction(ctx context.Context, c *app.RequestContext) {
	acitonID := c.Param(acitonID)
	id, err := primitive.ObjectIDFromHex(acitonID)
	if err != nil {
		hlog.CtxErrorf(ctx, "%s GetAction get objectid err=%s",
			c.Response.Header.Get(utils.RequestID),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	action, err := ctrl.daler.FindActionByID(ctx, id)
	if err != nil {
		hlog.CtxErrorf(ctx, `%s GetAction find action err="%s"`,
			utils.GetLogID(ctx),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusInternalServerError, err: err})
		return
	}

	res := &GetActionResponse{
		Action: action,
	}

	sendOK(ctx, c, okOption{statusCode: consts.StatusOK, obj: res})
}

type UpdateActionRequest struct {
	*dal.Action
}

type UpdateActionResponse struct {
	BaseMsg
}

// UpdateAction
func (ctrl *Controller) UpdateAction(ctx context.Context, c *app.RequestContext) {
	acitonID := c.Param(acitonID)

	id, err := primitive.ObjectIDFromHex(acitonID)
	if err != nil {
		hlog.CtxErrorf(ctx, "%s UpdateAction get objectid err=%s",
			c.Response.Header.Get(utils.RequestID),
			err.Error(),
		)
		sendError(ctx, c, errorOption{statusCode: consts.StatusBadRequest, err: err})
		return
	}

	req := &UpdateActionRequest{}
	if err := c.BindAndValidate(req); err != nil {
		hlog.CtxErrorf(ctx, `%s UpdateAction BindAndValidate request parameters err="%s"`,
			utils.GetLogID(ctx),
			err.Error(),
		)
		return
	}

	// req validate
	req.ID = id

	if err := ctrl.daler.UpdateAction(ctx, req.Action); err != nil {
		if err != nil {
			hlog.CtxErrorf(ctx, `%s UpdateAction update action err="%s"`,
				utils.GetLogID(ctx),
				err.Error(),
			)
			sendError(ctx, c, errorOption{statusCode: consts.StatusInternalServerError, err: err})
			return
		}
	}

	res := &UpdateActionResponse{
		BaseMsg: BaseMsg{
			Message: "update action success",
		},
	}

	sendOK(ctx, c, okOption{statusCode: consts.StatusOK, obj: res})
}
