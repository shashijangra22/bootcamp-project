package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// fire the gRPC Server
	fmt.Println("Hello from grpc server.")

	lis, err := net.Listen("tcp", "0.0.0.0:5051")

	if err != nil {
		log.Fatalf("Sorry failed to load server %v:", err)
	}

	s := grpc.NewServer()

	// orderpb.RegisterQueryServiceServer(s, &server{})

	if s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
