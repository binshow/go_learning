package main

import (
	"context"
	"fmt"
	. "go_learning/gRpc_study/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type greeter struct {
}

func (*greeter) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	fmt.Println(req)
	reply := &HelloReply{Message: "hello"}
	return reply, nil
}

func main() {
	// 创建一个TCP端口监听
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	// 注册 grpcurl 所需的 reflection 服务
	reflection.Register(server)
	// 注册业务服务
	RegisterGreeterServer(server, &greeter{})

	fmt.Println("grpc server start ...")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


