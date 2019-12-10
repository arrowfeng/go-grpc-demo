package main

import (
	"context"
	"log"
	"time"

	"github.com/arrowfeng/go-grpc-demo/helloworld/pb"
	"google.golang.org/grpc"
)

const (
	address = "localhost: 50001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatal("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewHelloWorldClient(conn)

	name := "zdf"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not sayhello: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
