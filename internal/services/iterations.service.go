package services

import (
	"loop-notes-api/internal/entities"
	"loop-notes-api/internal/repositories"
)

type IterationService interface{
	ListIterationsService(iterations *[]entities.Iteration) error
	CreateIterationService(iteration *entities.Iteration) error
	FindIterationService(iteration *entities.Iteration) error
	DeleteIterationService(iteration *entities.Iteration) error
	UpdateIterationService(iteration *entities.Iteration) error
}

type iterationService struct {
	repository repositories.IterationRepository
}

func NewIterationService(repository repositories.IterationRepository) *iterationService{
	return &iterationService{
		repository: repository,
	}	
}

func (service *iterationService) ListIterationsService(interations *[]entities.Iteration) error {
	if result := service.repository.ListIterationsRepository(interations); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *iterationService) CreateIterationService(interation *entities.Iteration) error {
	if result := service.repository.CreateIterationRepository(interation); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}


func (service *iterationService) FindIterationService(interation *entities.Iteration) error {
	if result := service.repository.FindIterationRepository(interation); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *iterationService) DeleteIterationService(interation *entities.Iteration) error {
	if result := service.repository.DeleteIterationRepository(interation); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *iterationService) UpdateIterationService(interation *entities.Iteration) error {
	if result := service.repository.UpdateIterationRepository(interation); result != nil && result.Error != nil {
		return result.Error
	}
	return nil
}