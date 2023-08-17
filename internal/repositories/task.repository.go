package repositories

import (
	"loop-notes-api/internal/entities"

	"gorm.io/gorm"
)

type TaskRepository interface {
	UpdateTaskRepository(task *entities.Task) *gorm.DB
	MarkTaskAsFavorite(task *entities.Task) *gorm.DB
	CreateTaskRepository(task *entities.Task) *gorm.DB
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{
		db: db,
	}
} 

func (repository *taskRepository) UpdateTaskRepository(task *entities.Task) *gorm.DB {
	return repository.db.Save(task)
}

func (repository *taskRepository) MarkTaskAsFavorite(task *entities.Task) *gorm.DB {
	return repository.db.Model(task).Where("id = ?", task.Default.Id).Update("is_favorite", task.IsFavorite)
}

func (reposity *taskRepository) CreateTaskRepository(task *entities.Task) *gorm.DB {
	return reposity.db.Create(task)
}