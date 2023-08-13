package services

import (
	"loop-notes-api/internal/entities"
	"loop-notes-api/internal/repositories"
)

type TaskService interface {
	ListTaskService(tasks *[]entities.Task) error
	CreateTaskService(task *entities.Task) error
	FindTaskService(task *entities.Task) error
	DeleteTaskService(task *entities.Task) error
	UpdateTaskService(task *entities.Task) error
	MarkTaskAsFavoriteService(task *entities.Task) error
}

type taskService struct {
	repository repositories.TaskRepository
}

func NewTaskService(repository repositories.TaskRepository) *taskService{
	return &taskService{
		repository: repository,
	}	
}

func (service *taskService) ListTaskService(tasks *[]entities.Task) error {
	if result := service.repository.ListTasksRepository(tasks); result != nil && result.Error != nil {
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


func (service *taskService) FindTaskService(task *entities.Task) error {
	if result := service.repository.FindTaskRepository(task); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *taskService) DeleteTaskService(task *entities.Task) error {
	if result := service.repository.DeleteTaskRepository(task); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
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