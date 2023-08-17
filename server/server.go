package server

import (
	"context"

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

func (t *TaskServer) AddTask(ctx context.Context, task *tasks.Tasks) (*tasks.Empty, error) {
	return &tasks.Empty{}, nil
}

func (t *TaskServer) MarkTaskAsCompleted(ctx context.Context, task *tasks.TaskRequest) (*tasks.TaskResponse, error) {
	return &tasks.TaskResponse{}, nil
}

func (t *TaskServer) ListTaks(ctx context.Context, task *tasks.Empty) error {
	return nil
}
