package middleware

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func TestLimiter(t *testing.T) {
	ctx := context.Background()
	l := NewLimiter(ctx, LimiterOption{60, 5})
	assert.Equal(t, l.limiter.Limit(), rate.Limit(60), "they should be equal")
}
