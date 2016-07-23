// Copyright (c) 2016 Kazumasa Kohtaka. All rights reserved.
// This file is available under the MIT license.

package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kkohtaka/godep-trial/none/server/pkg/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Hello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{
		Message: fmt.Sprintf("Hello, %s", in.Name),
	}, nil
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloServer(s, &server{})
	s.Serve(l)
}
