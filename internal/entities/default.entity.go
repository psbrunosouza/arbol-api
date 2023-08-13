package entities

import (
	"time"

	"gorm.io/gorm"
)

type Default struct {
	Id        uint  `json:"id" gorm:"primarykey"`
	CreatedAt *time.Time `json:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
}
