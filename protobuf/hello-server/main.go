package main

import (
	greet "GolangStudy/protobuf/hello-server/service"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (srv *server) SayHello(ctx context.Context, req *greet.HelloRequest) (*greet.HelloReply, error) {
	return &greet.HelloReply{message: "hello" + req.name}
}

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("err listen")
	}

	s := grpc.NewServer()

	greet.ResisterGreeterServer(s, &server{})

	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
