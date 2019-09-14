package main

import (
	"fmt"
	"go-envoy-boilerplate/todo"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := todo.Server{}
	grpcServer := grpc.NewServer()

	todo.RegisterTodoServiceServer(grpcServer, &s)

	log.Printf("Server started on port 8080\n")

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
