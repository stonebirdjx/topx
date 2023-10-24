package config

import (
	"fmt"
	"os"
	"strconv"
)

// Env 预置环境变量.
const (
	MongoDBURIEnv = "MONGODB_URI"
	MongoDBDBEnv  = "MONGODB_DB"
	RateLimitEnv  = "RATE_LIMIT"
	BurstEnv      = "BURST"
	RedisURIEnv   = "REDIS_URI"
)

// Global 全局环境变量.
type Global struct {
	MongoDBURI string
	MongoDBDB  string
	RateLimit  float64
	Burst      int
	RedisURI   string
}

func ReadFromEnv() *Global {
	ratelimit, err := strconv.ParseFloat(os.Getenv(RateLimitEnv), 64)
	if err != nil {
		ratelimit = defaultRetaLimit
	}

	burst, err := strconv.Atoi(os.Getenv(BurstEnv))
	if err != nil {
		burst = defaultBurst
	}

	return &Global{
		MongoDBURI: os.Getenv(MongoDBURIEnv),
		MongoDBDB:  os.Getenv(MongoDBDBEnv),
		RateLimit:  ratelimit,
		Burst:      burst,
		RedisURI:   os.Getenv(RedisURIEnv),
	}
}

func (g *Global) Validate() error {
	if g.MongoDBURI == "" {
		return fmt.Errorf("env=%s value can not be nil,may be not set", MongoDBURIEnv)
	}

	if g.MongoDBDB == "" {
		return fmt.Errorf("env=%s value can not be nil,may be not set", MongoDBDBEnv)
	}

	if g.RedisURI == "" {
		return fmt.Errorf("env=%s value can not be nil,may be not set", RedisURIEnv)
	}

	return nil
}

type EnvPart struct{}

func (e *EnvPart) Read() error {
	return nil
}

func (e *EnvPart) Validate() error {
	return nil
}
