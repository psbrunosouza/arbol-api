package entities

type Task struct {
	Default
	Name           string     `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Description    string     `json:"description,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	IsFavorite     bool       `json:"isFavorite" gorm:"default:false"`
	IterationID    int        `json:"iterationId,omitempty" gorm:"default:null;"`
	Iteration      *Iteration `json:"iteration,omitempty" gorm:"default:null;constraint:OnDelete:CASCADE;"`
	StatusID       int        `json:"statusId,omitempty" gorm:"default:null;"`
	Status         *Status    `json:"status,omitempty" gorm:"default:null;"`
	EasinessFactor int        `json:"easinessFactor,omitempty" gorm:"not null;default:null"`
	ReviewVersion  int        `json:"reviewVersion,omitempty" gorm:"not null;default:null"`
}
