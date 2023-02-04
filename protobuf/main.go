package main

import (
	"context"
	"fmt"
	"golang/GolangStudy/protobuf/service"
	pb "golang/GolangStudy/protobuf/service"
	"net"

	"reflect"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type server struct {
	pb.UnimplementedProdServiceServer
}

func (srv *server) GetProductStock(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{Prod_stock: 1 + req.Prod_id}, nil
}

func main() {
	user := &service.Student{
		Name: "张三",
		Male: false,
	}
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	//反序列化

	newUser := &service.Student{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}

	fmt.Println(newUser, reflect.TypeOf(newUser))
	l, _ := net.Listen("tcp", ":9090")
	grpcServer := grpc.NewServer()
	pb.RegisterProdServiceServer(grpcServer, &server{})

	err = grpcServer.Serve(l)
	if err != nil {
		panic(err)
	}

}
