// Copyright 2023 The Author stonebird. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package config defines the system config.
// Include environment variables and yaml file.
package config

const (
	RequestID = "X-Request-Id"
	LogID     = "Log-Id"
	ReadMode  = "READ_MODE"
)

const (
	defaultRetaLimit float64 = 5
	defaultBurst     int     = 5
)

type Key string
