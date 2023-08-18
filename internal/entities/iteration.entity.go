package entities

type Iteration struct {
	Default
	IsLoop bool   `json:"isLoop,omitempty" gorm:"default:false;"`
	Tasks  []Task `json:"tasks,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
}
