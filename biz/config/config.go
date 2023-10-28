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

// Package config defines the system config.
// Include environment variables and yaml file.
package config

import "os"

const (
	envCfgMode = "CONF_MODE"
	yamlMode   = "yaml"
)

// Configer define some functions to obtain system built-in parameters.
type Configer interface {
	GetMongDBURI() string
	GetMongDBName() string
	GetRedisURI() string
	GetRateLimte() float64 // RateLimte限制定义了某些事件的最大频率。 限制表示为每秒的事件数。 零限制不允许发生任何事件。
	GetBurst() int
}

// InitConfiger initialize system parameter configuration.
func InitConfiger() (Configer, error) {
	switch os.Getenv(envCfgMode) {
	case yamlMode:
		return initYamlConfiger()
	default:
		return initEnvConfiger()
	}
}
