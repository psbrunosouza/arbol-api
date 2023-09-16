package entities

type Score struct {
	Default
	Description string  `json:"description,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Value       float64 `json:"value,omitempty" gorm:"not null;"`
}