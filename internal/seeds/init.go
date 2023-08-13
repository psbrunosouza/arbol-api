package seeds

import (
	"loop-notes-api/internal/databases"
)

func Init() {
	//Instancias de seed
	statusSeed := NewStatusSeedDatabase(databases.Database)

	//Seeds aplicadas
	statusSeed.StatusSeed()
}