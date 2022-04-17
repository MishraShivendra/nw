package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"server/pkg/protoGo"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	protoGo.UnimplementedGreeterServer
}

func (s *server) SayHello(c context.Context, in *protoGo.HelloRequest) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	fmt.Println("Hello", in.Name)
	return out, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	protoGo.RegisterGreeterServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
