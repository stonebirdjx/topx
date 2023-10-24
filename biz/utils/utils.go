// Copyright 2023 The Author stonebird. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package utils defines the system common functions.
package utils

import (
	"context"

	"github.com/stonebirdjx/topx/biz/config"
)

const (
	sysLogid = "0000000-0000-0000-0000-000000000000"
)

// GetLogID get the ctx Log-ID value.
func GetLogID(ctx context.Context) string {
	logid, ok := ctx.Value(config.Key(config.LogID)).(string)
	if !ok {
		logid = sysLogid
	}
	return logid
}
