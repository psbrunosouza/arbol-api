package seeds

import (
	"loop-notes-api/internal/entities"

	"gorm.io/gorm"
)

type ScoreSeeder interface {
	ScoreSeed() error
}

type scoreSeedDatabase struct {
	database *gorm.DB
}

func NewScoreSeedDatabase(db *gorm.DB) *scoreSeedDatabase {
	return &scoreSeedDatabase{
		database: db,
	}
}

func (scoreSeedDatabase *scoreSeedDatabase) ScoreSeed() error {
	scores := []entities.Score{
		{
			Description: "very-easy",
		},
		{
			Description: "easy",
		},
		{
			Description: "dificult",
		},
		{
			Description: "very-dificult",
		},
	}

	for _, score := range scores {
		rawScore := entities.Score{}

		if result := scoreSeedDatabase.database.Where("description = ?", score.Description).First(&rawScore); result.Error == nil {
			return nil
		}

		if result := scoreSeedDatabase.database.Create(&score); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
