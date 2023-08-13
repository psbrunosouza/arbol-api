package entities

type Iteration struct {
	Default
	IsLoop           bool   `json:"isLoop,omitempty" gorm:"default:false;"`
	IterationsNumber int    `json:"iterationsNumber,omitempty"`
	IterationsSpace  int    `json:"iterationsSpace,omitempty"`
	Tasks            []Task `json:"tasks,omitempty"`
}
