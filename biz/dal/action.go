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

package dal

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/stonebirdjx/topx/biz/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Scheme string

type Proxy struct {
	Scheme Scheme `json:"scheme" bson:"scheme"`
	Host   string `json:"host" bson:"host"`
	Path   string `json:"path" bson:"path"`
	Weight int    `json:"weight" bson:"weight"`
}

type Action struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `json:"name" bson:"name"`
	ServiceName string             `json:"service_name" bson:"service_name"`
	Description string             `json:"description" bson:"description"`
	RateLimit   float64            `json:"rate_limit" bson:"rate_limit"`
	IsAuth      bool               `json:"is_auth" bson:"is_auth"`
	Proxy       []Proxy            `json:"proxy" bson:"proxy"`
	Timeout     int                `json:"timeout" bson:"timeout"`
	Version     string             `json:"version" bson:"version"`
}

func (a *Action) Validate() error {
	return nil
}

func (b *Basic) CreateAction(ctx context.Context, action *Action) error {
	action.ID = primitive.NewObjectID()
	_, err := b.topxCol.InsertOne(ctx, action)
	return err
}

// ListActions(ctx context.Context, opt ListActionOption) ([]*Action, error)
func (b *Basic) ListActions(ctx context.Context, opt ListActionOption) (*[]Action, int, error) {
	filter := bson.M{}
	skip := (opt.PageNum - 1) * opt.PapeSize
	findOpt := &options.FindOptions{
		Limit: &opt.PapeSize,
		Skip:  &skip,
	}

	cursor, err := b.topxCol.Find(ctx, filter, findOpt)
	if err != nil {
		hlog.CtxErrorf(ctx, "%s topcol find err=%s",
			utils.GetLogID(ctx),
			err.Error(),
		)
		return nil, 0, err
	}

	actions := &[]Action{}
	if err := cursor.All(ctx, actions); err != nil {
		hlog.CtxErrorf(ctx, "%s cursor actions err=%s",
			utils.GetLogID(ctx),
			err.Error(),
		)
		return nil, 0, err
	}

	return actions, len(*actions), nil
}

func (b *Basic) DeleteActionByID(ctx context.Context, id primitive.ObjectID) error {
	_, err := b.topxCol.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (b *Basic) DeleteActions(ctx context.Context, ids []primitive.ObjectID) error {
	filter := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}

	_, err := b.topxCol.DeleteMany(ctx, filter)
	return err
}

func (b *Basic) FindActionByID(ctx context.Context, id primitive.ObjectID) (*Action, error) {
	a := &Action{}
	fliter := bson.M{
		"_id": id,
	}

	err := b.topxCol.FindOne(ctx, fliter).Decode(a)
	return a, err
}

func (b *Basic) UpdateAction(ctx context.Context, action *Action) error {
	fliter := bson.M{
		"_id": action.ID,
	}

	_, err := b.topxCol.ReplaceOne(ctx, fliter, action)
	return err
}

func (b *Basic) FindActionByOpt(ctx context.Context, opt FindActionOption) (*Action, error) {
	a := &Action{}
	fliter := bson.M{
		"service_name": opt.ServiceName,
		"name":         opt.ActionName,
		"version":      opt.Version,
	}

	err := b.topxCol.FindOne(ctx, fliter).Decode(a)
	return a, err
}
