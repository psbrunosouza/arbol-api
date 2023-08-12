package entities

import (
	"time"

	"gorm.io/gorm"
)

type Default struct {
	Id        uint  `gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
