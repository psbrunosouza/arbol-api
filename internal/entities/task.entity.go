package entities

type Task struct {
	Default
	Name        string `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Description string `json:"description,omitempty" gorm:"not null;default:null;type:varchar(255)"`
}