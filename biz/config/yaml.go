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

type Yaml struct{}

func (y *Yaml) GetMongDBURI() string {
	return ""
}

func (y *Yaml) GetMongDBName() string {
	return ""
}

func (y *Yaml) GetRedisURI() string {
	return ""
}

func (y *Yaml) GetRateLimte() float64 {
	return 0
}

func (y *Yaml) GetBurst() int {
	return 0
}

// initYamlConfiger init configer by yaml
func initYamlConfiger() (Configer, error) {
	return &Yaml{}, nil
}
