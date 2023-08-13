package entities

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Default
	Name          string `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Description   string `json:"description,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	IsFavorite    bool   `json:"isFavorite" gorm:"default:false"`
	IsLoop        bool   `json:"isLoop" gorm:"default:false"`
	NextIteration *time.Time `json:"nextIteration,omitempty" gorm:"default:null;"`
	IterationID int `json:"iterationId,omitempty"`
  Iteration   *Iteration `json:"Iteration,omitempty"`
}

func (task *Task) BeforeCreate(db *gorm.DB) error {
	if (!task.NextIteration.IsZero()){
		task.IsLoop = true
		return nil
	}

	return db.Error
}
