package controllers

import (
	"loop-notes-api/internal/entities"
	"loop-notes-api/internal/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TasksController interface {
	UpdateTaskController(context echo.Context) error
	MarkTaskAsFavoriteController(context echo.Context) error
	CreateTaskController(context echo.Context) error
}

type controller struct {
	service services.TaskService
}

func NewTaskController(service services.TaskService) *controller {
	return &controller{
		service: service,
	}
}

func (controller *controller) UpdateTaskController(context echo.Context) error {
	id, stringParseError := strconv.Atoi(context.Param("id"))

	if stringParseError != nil {
			return context.JSON(http.StatusBadRequest, echo.Map{"error": "Parâmetro de busca invalido ou fora do formato esperado"})
	}

	task := &entities.Task{}
	
	if bindError := context.Bind(task); bindError != nil {
		return context.JSON(http.StatusUnprocessableEntity, echo.Map{"error": "Não foi possível processar a estrutura de dados informada"})
	}

	task.Id = uint(id)

	if controllerError := controller.service.MarkTaskAsFavoriteService(task); controllerError != nil {
		return context.JSON(http.StatusNotFound, echo.Map{"error": "Registro inexistente"})
	}

	return context.JSON(http.StatusOK, task)
}

func (controller *controller) MarkTaskAsFavoriteController(context echo.Context) error {
	id, stringParseError := strconv.Atoi(context.Param("id"))

	if stringParseError != nil {
			return context.JSON(http.StatusBadRequest, echo.Map{"error": "Parâmetro de busca invalido ou fora do formato esperado"})
	}
	
	task := &entities.Task{}

	if bindError := context.Bind(task); bindError != nil {
		return context.JSON(http.StatusUnprocessableEntity, echo.Map{"error": "Não foi possível processar a estrutura de dados informada"})
	}

	task.Id = uint(id)

	if controllerError := controller.service.MarkTaskAsFavoriteService(task); controllerError != nil {
		return context.JSON(http.StatusNotFound, echo.Map{"error": "Registro inexistente"})
	}

	return context.JSON(http.StatusOK, task)
}

func (controller *controller) CreateTaskController(context echo.Context) error {
	task := &entities.Task{}

	if err := context.Bind(task); err != nil {
		return context.JSON(http.StatusUnprocessableEntity, echo.Map{"error": "Não foi possível processar a estrutura de dados informada"})
	}

	if err := controller.service.CreateTaskService(task); err != nil {
		return context.JSON(http.StatusNotFound, echo.Map{"error": err})
	}

	return context.JSON(http.StatusOK, task)
}