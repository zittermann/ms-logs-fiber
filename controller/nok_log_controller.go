package controller

import (
	"ms_logs_go/data/request"
	"ms_logs_go/helper"
	"ms_logs_go/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NokLogController struct {
	nokLogService service.NokLogService
}

func NewNokLogController(service service.NokLogService) *NokLogController {
	return &NokLogController{
		nokLogService: service,
	}
}

// @Summary Obtener un Nok Log por ID
// @Tags Nok Logs
// @ID get-nok-log-by-id
// @Produce json
// @Param id path string true "nok log ID"
// @Success 200 {object} response.NokLogResponse
// @Failure 404 {object} response.Response
// @Router /api/nok-logs/{id} [get]
func(controller *NokLogController) FindById(ctx *fiber.Ctx) error {
	logId := ctx.Params("id")

	id, err := strconv.Atoi(logId)
	helper.ErrorPanic(err)

	logResponse := controller.nokLogService.FindById(uint(id))

	if (logResponse.ID == 0) {

		return helper.HandleNotFound(ctx, "No existe ningún " +
		 "Nok Log con ese ID")

	} 

	return ctx.Status(fiber.StatusOK).JSON(logResponse)

}

// @Summary Crea un nuevo Nok Log
// @Tags Nok Logs
// @ID create-nok-log
// @Produce json
// @Param data body response.NokLogResponse true "noklog data"
// @Success 200 {object} response.NokLogResponse
// @Failure 400 {object} response.Response
// @Router /api/nok-logs [post]
func(controller *NokLogController) Create(ctx *fiber.Ctx) error {
	createRequest := request.CreateNokLogRequest{}
	
	err := ctx.BodyParser(&createRequest)
	helper.ErrorPanic(err)

	// Guardamos el Nok Log
	controller.nokLogService.Create(createRequest)

	// Recuperamos el último Nok Log guardado (es decir, el que recién creamos)
	// y devolvemos con toda la información que completó la BD
	logResponse := controller.nokLogService.FindLatest()

	return ctx.Status(fiber.StatusCreated).JSON(logResponse)

}
