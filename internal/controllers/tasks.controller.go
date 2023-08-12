package controllers

import (
	"loop-notes-api/internal/entities"
	"loop-notes-api/internal/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TasksController interface {
	ListTaskController(context echo.Context) error
	CreateTaskController(context echo.Context) error
	FindTaskController(context echo.Context) error
	UpdateTaskController(context echo.Context) error
	DeleteTaskController(context echo.Context) error
	MarkTaskAsFavoriteController(context echo.Context) error
}

type controller struct {
	service services.TaskService
}

func NewTaskController(service services.TaskService) *controller {
	return &controller{
		service: service,
	}
}

func (controller *controller) ListTasksController(context echo.Context) error {
	var tasks []entities.Task

	if err := controller.service.ListTaskService(&tasks); err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return context.JSON(http.StatusOK, tasks)
}

func (controller *controller) CreateTaskController(context echo.Context) error {
	task := &entities.Task{}
	
	if bindError := context.Bind(task); bindError != nil {
		return context.JSON(http.StatusUnprocessableEntity, echo.Map{"error": "Não foi possível processar a estrutura de dados informada"})
	}

	if controllerError := controller.service.CreateTaskService(task); controllerError != nil {
		return context.JSON(http.StatusUnprocessableEntity, echo.Map{"error": "Erro na criação da tarefa os campos: 'name' e 'description' não podem ser nulos"})
	}

	return context.JSON(http.StatusOK, task)
}

func (controller *controller) FindTaskController(context echo.Context) error {
	id, stringParseError := strconv.Atoi(context.Param("id"))

	if stringParseError != nil {
			return context.JSON(http.StatusBadRequest, echo.Map{"error": "Parâmetro de busca invalido ou fora do formato esperado"})
	}
	
	task := &entities.Task{
		Default: entities.Default{
			Id: uint(id),
		},
	}

	if controllerError := controller.service.FindTaskService(task); controllerError != nil {
		return context.JSON(http.StatusNotFound, echo.Map{"error": "Registro inexistente"})
	}

	return context.JSON(http.StatusOK, task)
}

func (controller *controller) DeleteTaskController(context echo.Context) error {
	id, stringParseError := strconv.Atoi(context.Param("id"))

	if stringParseError != nil {
			return context.JSON(http.StatusBadRequest, echo.Map{"error": "Parâmetro de busca invalido ou fora do formato esperado"})
	}
	
	task := &entities.Task{
		Default: entities.Default{
			Id: uint(id),
		},
	}

	if controllerError := controller.service.FindTaskService(task); controllerError != nil {
		return context.JSON(http.StatusNotFound, echo.Map{"error": "Registro inexistente"})
	}

	if controllerError := controller.service.DeleteTaskService(task); controllerError != nil {
		return context.JSON(http.StatusNotFound, echo.Map{"error": "Registro inexistente"})
	}

	return context.JSON(http.StatusOK, nil)
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