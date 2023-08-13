package controllers

import (
	"loop-notes-api/internal/entities"
	"loop-notes-api/internal/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IterationsController interface {
	ListIterationsController(context echo.Context) error
	CreateIterationController(context echo.Context) error
	FindIterationController(context echo.Context) error
	UpdateIterationController(context echo.Context) error
	DeleteIterationController(context echo.Context) error
}

type iterationsController struct {
	service services.IterationService
}

func NewIterationController(service services.IterationService) *iterationsController {
	return &iterationsController{
		service: service,
	}
}

func (controller *iterationsController) ListIterationsController(context echo.Context) error {
	var iterations []entities.Iteration

	if err := controller.service.ListIterationsService(&iterations); err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return context.JSON(http.StatusOK, iterations)
}

func (controller *iterationsController) CreateIterationController(context echo.Context) error {
	iteration := &entities.Iteration{}
	
	if bindError := context.Bind(iteration); bindError != nil {
		return context.JSON(http.StatusUnprocessableEntity, echo.Map{"error": "Não foi possível processar a estrutura de dados informada"})
	}

	if controllerError := controller.service.CreateIterationService(iteration); controllerError != nil {
		return context.JSON(http.StatusUnprocessableEntity, echo.Map{"error": "Erro na criação da tarefa os campos: 'name' e 'description' não podem ser nulos"})
	}

	return context.JSON(http.StatusOK, iteration)
}

func (controller *iterationsController) FindIterationController(context echo.Context) error {
	id, stringParseError := strconv.Atoi(context.Param("id"))

	if stringParseError != nil {
			return context.JSON(http.StatusBadRequest, echo.Map{"error": "Parâmetro de busca invalido ou fora do formato esperado"})
	}
	
	iteration := &entities.Iteration{
		Default: entities.Default{
			Id: uint(id),
		},
	}

	if controllerError := controller.service.FindIterationService(iteration); controllerError != nil {
		return context.JSON(http.StatusNotFound, echo.Map{"error": "Registro inexistente"})
	}

	return context.JSON(http.StatusOK, iteration)
}

func (controller *iterationsController) DeleteIterationController(context echo.Context) error {
	id, stringParseError := strconv.Atoi(context.Param("id"))

	if stringParseError != nil {
			return context.JSON(http.StatusBadRequest, echo.Map{"error": "Parâmetro de busca invalido ou fora do formato esperado"})
	}
	
	iteration := &entities.Iteration{
		Default: entities.Default{
			Id: uint(id),
		},
	}

	if controllerError := controller.service.FindIterationService(iteration); controllerError != nil {
		return context.JSON(http.StatusNotFound, echo.Map{"error": "Registro inexistente"})
	}

	if controllerError := controller.service.DeleteIterationService(iteration); controllerError != nil {
		return context.JSON(http.StatusNotFound, echo.Map{"error": "Registro inexistente"})
	}

	return context.JSON(http.StatusOK, nil)
}

func (controller *iterationsController) UpdateIterationController(context echo.Context) error {
	id, stringParseError := strconv.Atoi(context.Param("id"))

	if stringParseError != nil {
			return context.JSON(http.StatusBadRequest, echo.Map{"error": "Parâmetro de busca invalido ou fora do formato esperado"})
	}

	iteration := &entities.Iteration{}
	
	if bindError := context.Bind(iteration); bindError != nil {
		return context.JSON(http.StatusUnprocessableEntity, echo.Map{"error": "Não foi possível processar a estrutura de dados informada"})
	}

	iteration.Id = uint(id)

	if controllerError := controller.service.UpdateIterationService(iteration); controllerError != nil {
		return context.JSON(http.StatusNotFound, echo.Map{"error": "Registro inexistente"})
	}

	return context.JSON(http.StatusOK, iteration)
}
