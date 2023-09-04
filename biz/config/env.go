package config

import (
	"fmt"
	"os"
)

// Env 预置环境变量.
const (
	MongoDBURIEnv = "MONGODB_URI"
	MongoDBDBEnv  = "MONGODB_DB"
)

// Global 全局环境变量.
type Global struct {
	MongoDBURI string
	MongoDBDB  string
}

func ReadFromEnv() *Global {
	return &Global{
		MongoDBURI: os.Getenv(MongoDBURIEnv),
		MongoDBDB:  os.Getenv(MongoDBDBEnv),
	}
}

func (g *Global) Validate() error {
	if g.MongoDBURI == "" {
		return fmt.Errorf("env=%s value can not be nil,may be not set.", MongoDBURIEnv)
	}

	if g.MongoDBDB == "" {
		return fmt.Errorf("env=%s value can not be nil,may be not set.", MongoDBDBEnv)
	}

	return nil
}
