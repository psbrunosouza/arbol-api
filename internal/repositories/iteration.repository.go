package repositories

import (
	"loop-notes-api/internal/entities"

	"gorm.io/gorm"
)

type IterationRepository interface {
	ListIterationsRepository(iterations *[]entities.Iteration) *gorm.DB
	FindIterationRepository(iteration *entities.Iteration) *gorm.DB
	CreateIterationRepository(iteration *entities.Iteration) *gorm.DB
	DeleteIterationRepository(iteration *entities.Iteration) *gorm.DB
	UpdateIterationRepository(iteration *entities.Iteration) *gorm.DB
}

type iterationRepository struct {
	db *gorm.DB
}

func NewIterationRepository(db *gorm.DB) *iterationRepository {
	return &iterationRepository{
		db: db,
	}
} 

func (repository *iterationRepository) ListIterationsRepository(iterations *[]entities.Iteration) *gorm.DB {
	result := repository.db.Preload("Tasks").Find(iterations)
	if result.Error != nil {
		return result
	}
	return nil
}

func (repository *iterationRepository) CreateIterationRepository(iteration *entities.Iteration) *gorm.DB {
	if result := repository.db.Create(iteration); result.Error != nil {
		return result
	}
	return nil
}


func (repository *iterationRepository) FindIterationRepository(iteration *entities.Iteration) *gorm.DB {
	if result := repository.db.Preload("Tasks").First(iteration); result.Error != nil {
		return result
	}
	return nil
}

func (repository *iterationRepository) DeleteIterationRepository(iteration *entities.Iteration) *gorm.DB {
	if result := repository.db.Delete(iteration); result.Error != nil {
		return result
	}
	return nil
}


func (repository *iterationRepository) UpdateIterationRepository(iteration *entities.Iteration) *gorm.DB {
	if result := repository.db.Save(iteration); result.Error != nil {
		return result
	}
	return nil
}
