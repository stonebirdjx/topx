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

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	handler "github.com/stonebirdjx/topx/biz/handler"
	"github.com/stonebirdjx/topx/biz/utils"
)

func main() {
	hlog.Infof(utils.BlessProgram())

	ctrl, err := handler.NewController()
	if err != nil {
		hlog.Fatalf(`new controller failed err="%s"`, err.Error())
	}

	// hertz
	h := server.Default(server.WithHostPorts(":6789"))
	register(ctrl, h)
	h.Spin()
	// hlog.Infof(utils.Thankyou())
}
