package services

import (
	"loop-notes-api/internal/entities"
	"loop-notes-api/internal/repositories"
)

type TaskService interface {
	UpdateTaskService(task *entities.Task) error
	MarkTaskAsFavoriteService(task *entities.Task) error
	CreateTaskService(task *entities.Task) error
}

type taskService struct {
	repository repositories.TaskRepository
}

func NewTaskService(repository repositories.TaskRepository) *taskService{
	return &taskService{
		repository: repository,
	}	
}

func (service *taskService) UpdateTaskService(task *entities.Task) error {
	if result := service.repository.UpdateTaskRepository(task); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *taskService) MarkTaskAsFavoriteService(task *entities.Task) error {
	if result := service.repository.MarkTaskAsFavorite(task); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *taskService) CreateTaskService(task *entities.Task) error {
	if result := service.repository.CreateTaskRepository(task); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}