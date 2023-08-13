package entities

import (
	"time"

	"gorm.io/gorm"
)

type Iteration struct {
	Default
	IsLoop          bool   `json:"isLoop,omitempty" gorm:"default:false;"`
	IterationSpaceInDays int    `json:"iterationsSpace,omitempty"`
	Tasks           []Task `json:"tasks,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
}

func (iteration *Iteration) BeforeCreate(db *gorm.DB) (err error) {
	var spaceInDays int = 0
	var iterationDate *time.Time

	if len(iteration.Tasks) >= 1 {
		iteration.IsLoop = true
	}

	spaceInDays = iteration.IterationSpaceInDays / len(iteration.Tasks)

	currentTime := time.Now()
	iterationDate = &currentTime

	if iteration.IterationSpaceInDays != 0 && iteration.IsLoop {
		for i := range iteration.Tasks {
			if i == 0 {
				taskIterationDate := time.Now()
				iteration.Tasks[i].NextIteration = &taskIterationDate
				iterationDate = &taskIterationDate
			} else {
				taskIterationDate := iterationDate.AddDate(0, 0, spaceInDays)
				iteration.Tasks[i].NextIteration = &taskIterationDate
				iterationDate = &taskIterationDate
			}
		}
	}

	return nil
}

func (iteration *Iteration) BeforeUpdate(db *gorm.DB) (err error) {
	if !iteration.IsLoop {
		for _, value := range iteration.Tasks {
			if response := db.Delete(&value); response != nil && response.Error != nil {
				return response.Error
			}
		}
	}

  return nil
}