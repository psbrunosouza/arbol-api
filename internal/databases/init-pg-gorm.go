package databases

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitPgGorm(entities ...interface{}) {
	database, databaseError := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		//  Logger: logger.Default.LogMode(logger.Info),
	})
	
	if(databaseError != nil){
		log.Fatalln("Error during database initialization")
	}

	if autoMigrateError := database.AutoMigrate(entities...); autoMigrateError != nil {
		log.Fatalln("Error during auto migration")
	}

	Database = database
}