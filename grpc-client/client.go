package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lcslima45/tasks-grpc/protos/tasks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Erro ao conectar %v", err)
	}
	defer connection.Close()
	client := tasks.NewTaskServiceClient(connection)
	task := &tasks.Tasks{
		Id:          3,
		Title:       "Tarefa 3",
		Description: "Descrição da tarefa 3",
	}
	response, err := client.AddTask(context.Background(), task)
	if err != nil {
		log.Println("Erro:", err)
	} else if response.Ok {
		log.Println("Tarefa adicionado com sucesso.")
	}
	task2 := &tasks.Tasks{
		Id:          4,
		Title:       "Tarefa 4",
		Description: "Descrição da tarefa 4",
	}
	response, err = client.AddTask(context.Background(), task2)
	if err != nil {
		log.Println("Erro:", err)
	} else if response.Ok {
		log.Println("Tarefa adicionado com sucesso.")
	}

	stream, err := client.ListTaks(context.Background(), &tasks.Empty{})
	if err != nil {
		log.Fatalf("Error dealing to listing the tasks")
	} else {
		for {
			tasks, err := stream.Recv()
			if err != nil {
				break
			}
			fmt.Printf("Task id=%d, Título=%s, Descrição=%s, Completed=%t\n", tasks.Id, tasks.Title, tasks.Description, tasks.Completed)
		}
	}

	responseIfCompleted, err := client.MarkTaskAsCompleted(context.Background(), &tasks.TaskRequest{
		Id:        3,
		Completed: true,
	})

	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("Response Ok:", responseIfCompleted.Ok)
	}

	log.Println("After alter some tasks........................................")

	stream, err = client.ListTaks(context.Background(), &tasks.Empty{})
	if err != nil {
		log.Fatalf("Error dealing to listing the tasks")
	} else {
		for {
			tasks, err := stream.Recv()
			if err != nil {
				break
			}
			fmt.Printf("Task id=%d, Título=%s, Descrição=%s, Completed=%t\n", tasks.Id, tasks.Title, tasks.Description, tasks.Completed)
		}
	}

	response, err = client.DeleteTask(context.Background(), &tasks.TaskDeleter{Id: 4})
	if err != nil {
		log.Fatalf("Error on deleting")
	} else {
		log.Println("Response Ok:", response.Ok)
	}
}
