package routes

import (
	"loop-notes-api/internal/controllers"
	"loop-notes-api/internal/databases"
	"loop-notes-api/internal/repositories"
	"loop-notes-api/internal/services"

	"github.com/labstack/echo/v4"
)

func TasksRoutes(e *echo.Group) {
	taskRepository := repositories.NewTaskRepository(databases.Database)
	taskService := services.NewTaskService(taskRepository)
	taskController := controllers.NewTaskController(taskService)

	e.PUT("/tasks/:id", taskController.UpdateTaskController)
	e.PATCH("/tasks/mark-as-favorite/:id", taskController.MarkTaskAsFavoriteController)
}