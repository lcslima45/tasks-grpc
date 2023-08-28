package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/lcslima45/tasks-grpc/models"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{
		db: db,
	}
}

func (repo *taskRepository) AddNewTask(ctx context.Context, id int32, title string, description string, completed bool) (bool, error) {
	newTask := &models.TaskModel{
		NumTask:     id,
		Title:       title,
		Description: description,
		Completed:   completed,
	}

	result := repo.db.Create(newTask)
	if result.Error == nil {
		return true, nil
	}
	log.Println("Error:", result.Error)
	return false, result.Error
}

func (repo *taskRepository) MarkTaskAsCompleted(ctx context.Context, id int32, completed bool) (bool, error) {
	var taskToUpdate models.TaskModel

	if err := repo.db.Where("num_task = ?", id).First(&taskToUpdate).Error; err != nil {
		return false, err
	}

	taskToUpdate.Completed = completed

	if err := repo.db.Save(&taskToUpdate).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (repo *taskRepository) ListTasks() ([]models.TaskModel, error) {
	var tasks []models.TaskModel
	if err := repo.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repo *taskRepository) DeleteTask(ctx context.Context, id int32) (bool, error) {
	result := repo.db.Delete(&models.TaskModel{}, id)
	if result.Error != nil {
		err := fmt.Errorf("error on deleting: %v", result.Error)
		return false, err
	}

	return true, nil
}
