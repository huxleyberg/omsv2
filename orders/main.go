package main

import (
	"context"
	"log"
	"net"

	common "github.com/huxleyberg/omsv2commons"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("ORDER_SERVER_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to dial server : %v", err)
	}
	defer listener.Close()

	store := NewStore()
	svc := NewService(store)

	NewGRPCHandler(grpcServer)

	svc.CreateOrder(context.Background())

	log.Println("GRPC server started at ", grpcAddr)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to start grpc server : %v", err)
	}
}
