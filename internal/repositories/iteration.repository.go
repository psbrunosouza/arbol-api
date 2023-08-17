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
	return repository.db.Preload("Tasks").Find(iterations)
}

func (repository *iterationRepository) CreateIterationRepository(iteration *entities.Iteration) *gorm.DB {
	return repository.db.Create(iteration)
}

func (repository *iterationRepository) FindIterationRepository(iteration *entities.Iteration) *gorm.DB {
	return repository.db.Preload("Tasks").First(iteration)
}

func (repository *iterationRepository) DeleteIterationRepository(iteration *entities.Iteration) *gorm.DB {
	return repository.db.Delete(iteration)
}


func (repository *iterationRepository) UpdateIterationRepository(iteration *entities.Iteration) *gorm.DB {
	iterationMapped := map[string]interface{}{
		"IsLoop": iteration.IsLoop,
	}

	result := repository.db.Model(iteration).Where("id = ?", iteration.Id).Updates(iteration).First(iteration)

	result = repository.db.Model(iteration).Where("id = ?", iteration.Id).Updates(iterationMapped).First(iteration)
	
	return result
}
