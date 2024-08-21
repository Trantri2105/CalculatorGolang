package main

import (
	endpoints "gokit/endpoint"
	"gokit/proto/pb"
	"gokit/service"
	"gokit/transport"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	mathService := service.NewMathService()
	mathEndpoint := endpoints.NewMathEndpoint(mathService)
	grpcServer := transport.NewGRPCServer(mathEndpoint)

	server := grpc.NewServer()
	pb.RegisterMathServiceServer(server, grpcServer)

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
