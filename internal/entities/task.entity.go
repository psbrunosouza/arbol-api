package entities

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Default
	Name           string     `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Description    string     `json:"description,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	IsFavorite     bool       `json:"isFavorite" gorm:"default:false"`
	IterationID    uint       `json:"iterationId,omitempty" gorm:"default:null;"`
	Iteration      *Iteration `json:"iteration,omitempty" gorm:"default:null;constraint:OnDelete:CASCADE;"`
	StatusID       uint       `json:"statusId,omitempty" gorm:"default:null;"`
	Status         *Status    `json:"status,omitempty" gorm:"default:null;"`
	EasinessFactor float64    `json:"easinessFactor,omitempty" gorm:"not null;default:null"`
	ReviewVersion  int        `json:"reviewVersion,omitempty" gorm:"not null;default:null"`
	NextIteration  *time.Time `json:"nextIteration,omitempty" gorm:"not null;default:null"`
	Score          *Score     `json:"score,omitempty" gorm:"default:null;"`
	ScoreID        uint       `json:"scoreId,omitempty" gorm:"default:null;"`
}

func calculatePeriod(review int, easinessFactor float64, basePeriod float64, isBad bool) float64 {
	if isBad {
		return (float64(review) * basePeriod) * easinessFactor
	} else {
		return (float64(review) * basePeriod) / easinessFactor
	}
}

func (task *Task) BeforeCreate(db *gorm.DB) error {
	lastTask := &Task{}

	if result := db.Order("id desc").Last(lastTask); result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		task.ReviewVersion = 1
		task.EasinessFactor = 0.5
		nextIterationDate := time.Now()
		task.NextIteration = &nextIterationDate
	} else {
		task.ReviewVersion = lastTask.ReviewVersion + 1
		task.EasinessFactor = lastTask.EasinessFactor
		period := calculatePeriod(task.ReviewVersion, task.EasinessFactor, 20, true)
		nextIterationDate := lastTask.NextIteration.Add(time.Duration(int(period) * int(time.Minute)))
		task.NextIteration = &nextIterationDate
	}

	return nil
}
