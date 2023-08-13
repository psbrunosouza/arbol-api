package main

import (
	"log"
	"loop-notes-api/internal/databases"
	"loop-notes-api/internal/entities"
	"loop-notes-api/internal/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if dotEnvError := godotenv.Load(); dotEnvError != nil {
		log.Fatal("Error loading .env file")
	}

	databases.InitPgGorm(
		&entities.Task{},
		&entities.Iteration{},
	)
}

func main() {
	echo := echo.New()
	
	api := echo.Group("/v1")

	routes.TasksRoutes(api)
	routes.IterationRoutes(api)

	echo.Logger.Fatal(echo.Start(os.Getenv("PORT")))
}