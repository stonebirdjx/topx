package util

import (
	"context"

	"github.com/stonebirdjx/topx/biz/config"
)

const (
	sysLogid = "0000000-0000-0000-0000-000000000000"
)

func GetLogID(ctx context.Context) string {
	logid, ok := ctx.Value(config.Key(config.LogID)).(string)
	if !ok {
		logid = sysLogid
	}
	return logid
}
