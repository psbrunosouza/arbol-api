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

type service struct {
	repository repositories.TaskRepository
}

func NewTaskService(repository repositories.TaskRepository) *service{
	return &service{
		repository: repository,
	}	
}

func (service *service) ListTaskService(tasks *[]entities.Task) error {
	if result := service.repository.List(tasks); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *service) CreateTaskService(task *entities.Task) error {
	if result := service.repository.Create(task); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}


func (service *service) FindTaskService(task *entities.Task) error {
	if result := service.repository.Find(task); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *service) DeleteTaskService(task *entities.Task) error {
	if result := service.repository.Delete(task); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *service) UpdateTaskService(task *entities.Task) error {
	if result := service.repository.Update(task); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *service) MarkTaskAsFavoriteService(task *entities.Task) error {
	if result := service.repository.MarkAsFavorite(task); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}