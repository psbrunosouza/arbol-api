package seeds

import (
	"loop-notes-api/internal/databases"
)



func Init() {
	//Instancias de seed
	statusSeed := NewStatusSeedDatabase(databases.Database)
	scoreSeed := NewScoreSeedDatabase(databases.Database)

	//Seeds aplicadas
	statusSeed.StatusSeed()
	scoreSeed.ScoreSeed()
}