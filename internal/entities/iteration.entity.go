package entities

import (
	"time"

	"gorm.io/gorm"
)

type Iteration struct {
	Default
	IsLoop          bool   `json:"isLoop,omitempty" gorm:"default:false;"`
	IterationsSpace int    `json:"iterationsSpace,omitempty"`
	Tasks           []Task `json:"tasks,omitempty"`
}

func (iteration *Iteration) BeforeCreate(db *gorm.DB) (err error) {
	var spaceInDays int = 0
	var iterationDate *time.Time

	if len(iteration.Tasks) >= 1 {
		iteration.IsLoop = true
	}

	spaceInDays = iteration.IterationsSpace / len(iteration.Tasks)

	currentTime := time.Now()
	iterationDate = &currentTime


if iteration.IterationsSpace != 0 && iteration.IsLoop {
   if iteration.IterationsSpace != 0 && iteration.IsLoop {
    for i := range iteration.Tasks {
			if i == 0 {
				taskIterationDate := time.Now()
				iteration.Tasks[i].NextIteration = &taskIterationDate
				iterationDate = &taskIterationDate
			}else {
				taskIterationDate := iterationDate.AddDate(0, 0, spaceInDays)
        iteration.Tasks[i].NextIteration = &taskIterationDate
        iterationDate = &taskIterationDate
			}  
    }
}
}

	return nil
}