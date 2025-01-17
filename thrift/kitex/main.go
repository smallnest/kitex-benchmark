/*
 * Copyright 2021 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"

	"github.com/cloudwego/kitex-benchmark/perf"
	"github.com/cloudwego/kitex-benchmark/thrift/kitex/kitex_gen/echo"
	echosvr "github.com/cloudwego/kitex-benchmark/thrift/kitex/kitex_gen/echo/echoserver"
)

const (
	port = ":8001"
)

var recorder = perf.NewRecorder("KITEX")

// EchoServerImpl implements the last service interface defined in the IDL.
type EchoServerImpl struct{}

// Echo implements the EchoServerImpl interface.
func (s *EchoServerImpl) Echo(ctx context.Context, req *echo.Request) (resp *echo.Response, err error) {
	switch req.Message {
	case "begin":
		recorder.Begin()
	case "end":
		recorder.End()
		recorder.Report()
	}
	return &echo.Response{
		Message: req.Message,
	}, nil
}

func main() {
	address := &net.UnixAddr{Net: "tcp", Name: port}
	svr := echosvr.NewServer(new(EchoServerImpl), server.WithServiceAddr(address))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
