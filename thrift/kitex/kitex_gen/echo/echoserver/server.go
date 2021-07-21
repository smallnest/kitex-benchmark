// Code generated by Kitex v0.0.1. DO NOT EDIT.
package echoserver

import (
	"github.com/cloudwego/kitex-benchmark/thrift/kitex/kitex_gen/echo"
	"github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler echo.EchoServer, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
