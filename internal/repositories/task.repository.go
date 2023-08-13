package repositories

import (
	"loop-notes-api/internal/entities"

	"gorm.io/gorm"
)

type TaskRepository interface {
	ListTasksRepository(tasks *[]entities.Task) *gorm.DB
	FindTaskRepository(task *entities.Task) *gorm.DB
	UpdateTaskRepository(task *entities.Task) *gorm.DB
	MarkTaskAsFavorite(task *entities.Task) *gorm.DB
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{
		db: db,
	}
} 

func (repository *taskRepository) ListTasksRepository(tasks *[]entities.Task) *gorm.DB {
	result := repository.db.Find(tasks)
	if result.Error != nil {
		return result
	}
	return nil
}


func (repository *taskRepository) FindTaskRepository(task *entities.Task) *gorm.DB {
	if result := repository.db.First(task); result.Error != nil {
		return result
	}
	return nil
}

func (repository *taskRepository) UpdateTaskRepository(task *entities.Task) *gorm.DB {
	if result := repository.db.Save(task); result.Error != nil {
		return result
	}
	return nil
}

func (repository *taskRepository) MarkTaskAsFavorite(task *entities.Task) *gorm.DB {
	if result := repository.db.Model(task).Where("id = ?", task.Default.Id).Update("is_favorite", task.IsFavorite); result.Error != nil {
		return result
	}

	return nil
}