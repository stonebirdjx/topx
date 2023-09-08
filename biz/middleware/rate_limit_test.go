package middleware

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func TestLimiter(t *testing.T) {
	NewLimiter(context.Background(), LimiterOptions{60, 5})
	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")
	assert.Equal(t, limiter.Limit(), rate.Limit(60), "they should be equal")
}
