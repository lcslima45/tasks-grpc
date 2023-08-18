package repository

import (
	"context"

	"github.com/lcslima45/tasks-grpc/models" // Certifique-se de importar o pacote correto para as definições de modelos
)

type TaskRepository interface {
	AddNewTask(ctx context.Context, id int32, title string, description string, completed bool) (bool, error)
	MarkTaskAsCompleted(ctx context.Context, id int32, completed bool) (bool, error)
	ListTasks() ([]models.TaskModel, error)
}
