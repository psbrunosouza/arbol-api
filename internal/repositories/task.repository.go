package repositories

import (
	"errors"
	"loop-notes-api/internal/entities"
	"loop-notes-api/internal/utils"
	_ "loop-notes-api/internal/utils"
	"time"

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
	if result := repository.db.Updates(task); result.Error != nil {
		return result
	}

	return repository.db.Model(task).Where("id = ?", task.Default.Id).Update("is_favorite", task.IsFavorite)
}

func (repository *taskRepository) MarkTaskAsFavorite(task *entities.Task) *gorm.DB {
	return repository.db.Model(task).Where("id = ?", task.Default.Id).Update("is_favorite", task.IsFavorite)
}

func (reposity *taskRepository) CreateTaskRepository(task *entities.Task) *gorm.DB {
	lastTask := &entities.Task{}
	score := &entities.Score{
		Default: entities.Default{
			Id: task.ScoreID,
		},
	}

	if result := reposity.db.Order("id desc").Last(score); result.Error != nil {
		return result
	}

	if result := reposity.db.Order("id desc").Last(lastTask); result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result
		}
		task.ReviewVersion = 1
		nextIterationDate := time.Now()
		task.NextIteration = &nextIterationDate
	}else {
		task.ReviewVersion = lastTask.ReviewVersion + 1

		period := utils.CalculatePeriod(task.ReviewVersion, *score)

		nextIterationDate := lastTask.NextIteration.Add(time.Duration(int(period) * int(time.Minute)))
		task.NextIteration = &nextIterationDate
	}

	return reposity.db.Create(task)
}
