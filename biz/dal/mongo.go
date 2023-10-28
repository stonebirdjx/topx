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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	topxColName = "topx"
)

type MongoOption struct {
	URI string
	DB  string
}

var (
	Mongoclient *mongo.Client
	TopCol      *mongo.Collection
)

func MongoInit(ctx context.Context, opt MongoOption) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(opt.URI))
	if err != nil {
		return err
	}

	Mongoclient = client
	TopCol = Mongoclient.Database(opt.DB).Collection(topxColName)
	return topColInit(ctx)
}

func topColInit(ctx context.Context) error {
	index := mongo.IndexModel{
		// bson.M 没有顺序，必须使用 bson.D
		Keys: bson.D{
			{Key: "name", Value: 1},
			{Key: "service_name", Value: 1},
			{Key: "version", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}

	_, err := TopCol.Indexes().CreateOne(ctx, index)
	return err
}
