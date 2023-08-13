package routes

import (
	"loop-notes-api/internal/controllers"
	"loop-notes-api/internal/databases"
	"loop-notes-api/internal/repositories"
	"loop-notes-api/internal/services"

	"github.com/labstack/echo/v4"
)

func IterationRoutes(e *echo.Group) {
	iterationRepository := repositories.NewIterationRepository(databases.Database)
	iterationService := services.NewIterationService(iterationRepository)
	iterationController := controllers.NewIterationController(iterationService)

	e.GET("/iterations", iterationController.ListIterationsController)
	e.POST("/iterations", iterationController.CreateIterationController)
	e.GET("/iterations/:id", iterationController.FindIterationController)
	e.PUT("/iterations/:id", iterationController.UpdateIterationController)
	e.DELETE("/iterations/:id", iterationController.DeleteIterationController)
}