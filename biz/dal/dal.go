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

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	topxColName = "topx"
)

// Daler Collection of some database methods.
type Daler interface {
	Allow(ctx context.Context, key string, limit redis_rate.Limit) (*redis_rate.Result, error)
}

type DalerOption struct {
	MongoDBURI  string
	MongoDBName string
	RedisURI    string
}

// Basic implement Daler functions.
type Basic struct {
	mongoClient  *mongo.Client
	topxCol      *mongo.Collection
	redisRDB     *redis.Client
	redisLimiter *redis_rate.Limiter
}

func NewDaler(opt DalerOption) (Daler, error) {
	basic := &Basic{}
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(opt.MongoDBURI))
	if err != nil {
		return nil, err
	}
	basic.mongoClient = mongoClient

	topxCol := basic.mongoClient.Database(opt.MongoDBName).Collection(topxColName)
	basic.topxCol = topxCol
	if basic.topColInit(context.Background()) != nil {
		return nil, err
	}

	redisOption, err := redis.ParseURL(opt.RedisURI)
	if err != nil {
		return nil, err
	}

	basic.redisRDB = redis.NewClient(redisOption)
	basic.redisLimiter = redis_rate.NewLimiter(basic.redisRDB)
	return basic, nil
}

func (b *Basic) topColInit(ctx context.Context) error {
	index := mongo.IndexModel{
		// bson.M there is no order and must be used bson.D
		Keys: bson.D{
			{Key: "name", Value: 1},
			{Key: "service_name", Value: 1},
			{Key: "version", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}

	_, err := b.topxCol.Indexes().CreateOne(ctx, index)
	return err
}
