package seeds

import (
	"loop-notes-api/internal/entities"

	"gorm.io/gorm"
)

type StatusSeeder interface {
	StatusSeed() error
}

type statusSeedDatabase struct {
	database *gorm.DB
}

func NewStatusSeedDatabase(db *gorm.DB) *statusSeedDatabase {
	return &statusSeedDatabase{
		database: db,
	}
}

func (statusSeedDatabase *statusSeedDatabase) StatusSeed() error {
	statuses := []entities.Status{
		{
			Description: "progress",
		},
		{
			Description: "done",
		},
	}

	for _, status := range statuses {
		rawStatus := entities.Status{}

		if result := statusSeedDatabase.database.Where("description = ?", status.Description).First(&rawStatus); result.Error == nil {
			return nil
		}

		if result := statusSeedDatabase.database.Create(&status); result.Error != nil {
			return result.Error
		}
	}

	return nil
}