package entities

import (
	"time"
)

type Task struct {
	Default
	Name          string `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Description   string `json:"description,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	IsFavorite    bool   `json:"isFavorite" gorm:"default:false"`
	NextIteration *time.Time `json:"nextIteration,omitempty" gorm:"default:null;"`
	IterationID 	int `json:"iterationId,omitempty" gorm:"default:null;"`
  Iteration   	*Iteration `json:"iteration,omitempty" gorm:"default:null;"`
	StatusID 			int	`json:"statusId,omitempty" gorm:"default:null;"`
	Status 				*Status `json:"status,omitempty" gorm:"default:null;"`
}

