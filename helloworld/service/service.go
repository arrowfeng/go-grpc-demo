/*
 * File : service.go
 * Author : arrowfeng
 * Date : 2019/12/10
 */
package service

import (
	"context"

	"github.com/arrowfeng/go-grpc-demo/helloworld/pb"
)

type HelloWorld struct {
}

func NewHelloWorldService() pb.HelloWorldServer {
	return &HelloWorld{}
}

func (s *HelloWorld) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{Message: "Hello : " + in.GetName()}, nil
}
