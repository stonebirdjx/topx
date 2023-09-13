package util

import (
	"crypto/tls"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/network/standard"
)

var (
	HzHTTPClient  *client.Client
	HzHTTPSClient *client.Client
)

//ProxyClientHTTPInit.
func ProxyClientHTTPInit() error {
	c, err := client.NewClient()
	if err != nil {
		return err
	}
	HzHTTPClient = c
	return nil
}

//
func ProxyClientHTTPSInit() error {
	clientCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	c, err := client.NewClient(
		client.WithTLSConfig(clientCfg),
		client.WithDialer(standard.NewDialer()),
	)
	if err != nil {
		return err
	}
	HzHTTPSClient = c
	return nil
}
