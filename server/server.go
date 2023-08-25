package server

import (
	"context"
	"log"

	models "github.com/lcslima45/tasks-grpc/models"

	tasks "github.com/lcslima45/tasks-grpc/protos/tasks"
	repository "github.com/lcslima45/tasks-grpc/repository"
)

type TaskServer struct {
	taskRepository repository.TaskRepository
	//tasksMap map[int32]*tasks.Tasks
}

func NewTasksServer(repo repository.TaskRepository) *TaskServer {
	return &TaskServer{
		taskRepository: repo,
	}
}

func (t *TaskServer) AddTask(ctx context.Context, task *tasks.Tasks) (*tasks.TaskResponse, error) {
	ok, err := t.taskRepository.AddNewTask(ctx, task.Id, task.Title, task.Description, task.Completed)
	if ok {
		return &tasks.TaskResponse{Ok: true}, nil
	}
	return &tasks.TaskResponse{Ok: false}, err
}

func (t *TaskServer) MarkTaskAsCompleted(ctx context.Context, taskRequest *tasks.TaskRequest) (*tasks.TaskResponse, error) {
	ok, err := t.taskRepository.MarkTaskAsCompleted(ctx, taskRequest.Id, taskRequest.Completed)
	if ok {
		return &tasks.TaskResponse{Ok: true}, nil
	}
	return &tasks.TaskResponse{Ok: false}, err
}

func (t *TaskServer) ListTaks(empty *tasks.Empty, stream tasks.TaskService_ListTaksServer) error {
	tasksToList, err := t.taskRepository.ListTasks()
	log.Println(tasksToList)

	for _, taskFromDB := range tasksToList {
		taskToGRPC := t.transformTableIntoGRPCTask(taskFromDB)
		if err := stream.Send(taskToGRPC); err != nil {
			return err
		}
	}
	return err
}

func (t *TaskServer) transformTableIntoGRPCTask(taskFromDB models.TaskModel) *tasks.Tasks {
	taskToGRPC := &tasks.Tasks{
		Id:          taskFromDB.NumTask,
		Title:       taskFromDB.Title,
		Description: taskFromDB.Description,
		Completed:   taskFromDB.Completed,
	}
	return taskToGRPC
}

func (t *TaskServer) DeleteTask(ctx context.Context, taskToDelete *tasks.TaskDeleter) (*tasks.TaskResponse, error) {
	ok, err := t.taskRepository.DeleteTask(ctx, taskToDelete.Id)

	if err != nil {
		log.Println("Error :", err)
		return &tasks.TaskResponse{Ok: ok}, err
	}

	return &tasks.TaskResponse{Ok: ok}, err
}
