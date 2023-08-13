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

	var seedDatabaseError error



for _, status := range statuses {
		var existingStatus entities.Status

		result := statusSeedDatabase.database.Where("description = ?", status.Description).First(&existingStatus)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				if createResult := statusSeedDatabase.database.Create(&status); createResult.Error != nil {
					seedDatabaseError = createResult.Error
				}
			} else {
				seedDatabaseError = result.Error
			}
		}
	}


	return seedDatabaseError
}