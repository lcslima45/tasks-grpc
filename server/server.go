package server

import (
	"context"
	"fmt"

	tasks "github.com/lcslima45/tasks-grpc/protos/tasks"
)

type TaskServer struct {
	tasksMap map[int32]*tasks.Tasks
}

func NewTasksServer() *TaskServer {
	mapa := make(map[int32]*tasks.Tasks)
	return &TaskServer{
		tasksMap: mapa,
	}
}

func (t *TaskServer) AddTask(ctx context.Context, task *tasks.Tasks) (*tasks.TaskResponse, error) {
	if _, existentValue := t.tasksMap[task.Id]; !existentValue {
		t.tasksMap[task.Id] = task
		return &tasks.TaskResponse{Ok: false}, nil
	}
	err := fmt.Errorf("the id=%d already exists in the database", task.Id)
	return &tasks.TaskResponse{Ok: false}, err
}

func (t *TaskServer) MarkTaskAsCompleted(ctx context.Context, taskRequest *tasks.TaskRequest) (*tasks.TaskResponse, error) {
	if _, existentValue := t.tasksMap[taskRequest.Id]; existentValue {
		t.tasksMap[taskRequest.Id].Completed = true
		return &tasks.TaskResponse{Ok: true}, nil
	}
	err := fmt.Errorf("the id=%d does not exist in the database", taskRequest.Id)
	return &tasks.TaskResponse{}, err
}

func (t *TaskServer) ListTaks(empty *tasks.Empty, stream tasks.TaskService_ListTaksServer) error {
	// Implemente a lógica para listar as tarefas e enviar pelo stream
	for _, task := range t.tasksMap {
		if err := stream.Send(task); err != nil {
			return err
		}
	}
	return nil
}
