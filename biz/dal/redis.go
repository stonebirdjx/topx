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
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

var (
	RedisRDB     *redis.Client
	RedisLimiter *redis_rate.Limiter
)

type RedisOption struct {
	URI string // "redis://<user>:<pass>@localhost:6379/<db>"
}

func RedisInit(opt RedisOption) error {
	o, err := redis.ParseURL(opt.URI)
	if err != nil {
		return err
	}

	RedisRDB = redis.NewClient(o)

	RedisLimiter = redis_rate.NewLimiter(RedisRDB)
	return nil
}
