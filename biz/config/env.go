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

package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/stonebirdjx/topx/biz/utils"
)

// Env 预置环境变量.
const (
	envMongoDBURI              = "MONGODB_URI"
	envMongoDBName             = "MONGODB_DB"
	envRateLimit               = "RATE_LIMIT"
	envBurst                   = "BURST"
	envRedisURI                = "REDIS_URI"
	defaultMongoDBName         = "topx"
	defaultRetaLimit   float64 = 5
	defaultBurst       int     = 5
)

type Env struct {
	mongoDBURI  string
	mongoDBName string
	rateLimit   float64
	burst       int
	redisURI    string
}

func (e *Env) validate() error {
	if utils.IsEmptyString(e.mongoDBURI) {
		return fmt.Errorf("env %s is empty", envMongoDBURI)
	}

	if utils.IsEmptyString(e.mongoDBName) {
		return fmt.Errorf("env %s is empty", envMongoDBName)
	}

	if utils.IsEmptyString(e.redisURI) {
		return fmt.Errorf("env %s is empty", envRedisURI)
	}

	return nil
}

func (e *Env) GetMongDBURI() string {
	return e.mongoDBURI
}

func (e *Env) GetMongDBName() string {
	return e.mongoDBName
}

func (e *Env) GetRedisURI() string {
	return e.redisURI
}

func (e *Env) GetRateLimte() float64 {
	return e.rateLimit
}

func (e *Env) GetBurst() int {
	return e.burst
}

// readFromEnv Read configuration from environment variables.
func readFromEnv() *Env {
	ratelimit, err := strconv.ParseFloat(os.Getenv(envRateLimit), 64)
	if err != nil {
		ratelimit = defaultRetaLimit
	}

	burst, err := strconv.Atoi(os.Getenv(envBurst))
	if err != nil {
		burst = defaultBurst
	}

	return &Env{
		mongoDBURI:  os.Getenv(envMongoDBURI),
		mongoDBName: os.Getenv(envMongoDBName),
		rateLimit:   ratelimit,
		burst:       burst,
		redisURI:    os.Getenv(envRedisURI),
	}
}

// initEnvConfiger init configer by env.
func initEnvConfiger() (Configer, error) {
	env := readFromEnv()
	if err := env.validate(); err != nil {
		return nil, err
	}

	return env, nil
}
