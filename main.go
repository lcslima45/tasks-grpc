package main

import (
	"log"
	"net"

	models "github.com/lcslima45/tasks-grpc/models"
	tasks "github.com/lcslima45/tasks-grpc/protos/tasks"
	repository "github.com/lcslima45/tasks-grpc/repository"
	server "github.com/lcslima45/tasks-grpc/server"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "admin:admin@tcp(localhost:3306)/tasks_manager?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.TaskModel{})
	taskRepo := repository.NewRepository(db)
	taskServer := server.NewTasksServer(taskRepo)
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	tasks.RegisterTaskServiceServer(grpcServer, taskServer)
	log.Println("Server is running at port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
