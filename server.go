package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"practice-grpc/proto"
)

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &proto.Server{}

	grpcServer := grpc.NewServer()

	proto.RegisterChatServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
