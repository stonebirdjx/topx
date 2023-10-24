// Copyright 2023 The Author stonebird. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	router "github.com/stonebirdjx/topx/biz/router"
)

// register registers all routers.
func register(r *server.Hertz) {

	router.GeneratedRegister(r)

	customizedRegister(r)
}
