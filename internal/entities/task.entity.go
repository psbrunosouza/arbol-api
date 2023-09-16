package entities

import (
	"time"
)

type Task struct {
	Default
	Name           string     `json:"name,omitempty" gorm:"not null;default:null;type:varchar(64)"`
	Description    string     `json:"description,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	IsFavorite     bool       `json:"isFavorite" gorm:"default:false"`
	IterationID    uint       `json:"iterationId,omitempty" gorm:"default:null;"`
	Iteration      *Iteration `json:"iteration,omitempty" gorm:"default:null;"`
	StatusID       uint       `json:"statusId,omitempty" gorm:"default:null;"`
	Status         *Status    `json:"status,omitempty" gorm:"default:null;"`
	ReviewVersion  int        `json:"reviewVersion,omitempty" gorm:"not null;default:null"`
	NextIteration  *time.Time `json:"nextIteration,omitempty" gorm:"not null;default:null"`
	Score          *Score     `json:"score,omitempty" gorm:"default:null;"`
	ScoreID        uint       `json:"scoreId,omitempty" gorm:"default:null;"`
}



