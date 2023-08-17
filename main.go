package main

import (
	"log"
	"net"

	tasks "github.com/lcslima45/tasks-grpc/protos/tasks"
	server "github.com/lcslima45/tasks-grpc/server"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	taskServer := server.NewTasksServer()
	grpcServer := grpc.NewServer()
	tasks.RegisterTaskServiceServer(grpcServer, taskServer)
	log.Println("Server is running at port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
