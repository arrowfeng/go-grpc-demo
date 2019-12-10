/*
 * File : main.go
 * Author : arrowfeng
 * Date : 2019/12/10
 */
package main

import (
	"log"
	"net"

	"github.com/arrowfeng/go-grpc-demo/helloworld/pb"

	"github.com/arrowfeng/go-grpc-demo/helloworld/service"

	"google.golang.org/grpc"
)

const (
	port = ":50001"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	srv := service.NewHelloWorldService()

	pb.RegisterHelloWorldServer(s, srv)

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
