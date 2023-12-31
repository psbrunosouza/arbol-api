package main

import (
	"log"
	"loop-notes-api/internal/databases"
	"loop-notes-api/internal/entities"
	"loop-notes-api/internal/routes"
	"loop-notes-api/internal/seeds"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	if dotEnvError := godotenv.Load(); dotEnvError != nil {
		log.Fatal("Error loading .env file")
	}

	databases.InitPgGorm(
		&entities.Status{},
		&entities.Score{},
		&entities.Task{},
		&entities.Iteration{},
	)

	seeds.Init()
}

func main() {
	echo := echo.New()

	echo.Use(middleware.CORS())
	
	api := echo.Group("/v1")

	routes.TasksRoutes(api)
	routes.IterationRoutes(api)

	echo.Logger.Fatal(echo.Start(os.Getenv("PORT")))
}