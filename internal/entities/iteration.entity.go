package entities

type Iteration struct {
	Default
	IsLoop          bool   `json:"isLoop,omitempty" gorm:"default:false;"`
	IterationSpaceInDays int    `json:"iterationsSpace,omitempty"`
	Tasks           []Task `json:"tasks,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
}
