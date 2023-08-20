package entities

type Iteration struct {
	Default
	IsLoop bool   `json:"isLoop" gorm:"default:false;"`
	Tasks  []Task `json:"tasks,omitempty"`
}
