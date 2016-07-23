// Copyright (c) 2016 Kazumasa Kohtaka. All rights reserved.
// This file is available under the MIT license.

package main

import (
	"log"

	"github.com/kkohtaka/godep-trial/glide/server/pkg/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	c, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer c.Close()
	client := hello.NewHelloClient(c)

	r, err := client.Hello(
		context.Background(),
		&hello.HelloRequest{Name: "glide-trial"})
	if err != nil {
		log.Fatalf("Failed to request: %v", err)
	}
	log.Printf("message: %s", r.Message)
}
