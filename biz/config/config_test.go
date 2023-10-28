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

package config_test

import (
	"os"
	"testing"

	"github.com/stonebirdjx/topx/biz/config"
	"github.com/stretchr/testify/assert"
)

func TestConfiger(t *testing.T) {
	if err := os.Setenv("CONF_MODE", "yaml"); err != nil {
		t.Fatalf(`set env err="%s"`, err.Error())
	}
	cfg, err := config.InitConfiger()
	if err != nil {
		t.Fatalf(`init config err="%s"`, err.Error())
	}

	assert.Equal(t, cfg.GetMongDBName(), "", "they should be equal")
	assert.Equal(t, cfg.GetMongDBURI(), "", "they should be equal")
	assert.Equal(t, cfg.GetRedisURI(), "", "they should be equal")
	assert.Equal(t, cfg.GetRateLimte(), float64(0), "they should be equal")
	assert.Equal(t, cfg.GetBurst(), 0, "they should be equal")
}
