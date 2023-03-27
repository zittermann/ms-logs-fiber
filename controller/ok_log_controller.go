package controller

import (
	"ms_logs_go/data/request"
	"ms_logs_go/helper"
	"ms_logs_go/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OkLogController struct {
	okLogService service.OkLogService
}

func NewOkLogController(service service.OkLogService) *OkLogController {
	return &OkLogController{
		okLogService: service,
	}
}

// @Summary Obtener un Ok Log por ID
// @Tags Ok Logs
// @ID get-ok-log-by-id
// @Produce json
// @Param id path string true "ok log ID"
// @Success 200 {object} response.OkLogResponse
// @Failure 404 {object} response.Response
// @Router /api/ok-logs/{id} [get]
func(controller *OkLogController) FindById(ctx *fiber.Ctx) error {
	logId := ctx.Params("id")

	id, err := strconv.Atoi(logId)
	helper.ErrorPanic(err)

	logResponse := controller.okLogService.FindById(uint(id))

		if (logResponse.ID == 0) {

		return helper.HandleNotFound(ctx, "No existe ningún " +
		 "Ok Log con ese ID")

	} 

	return ctx.Status(fiber.StatusOK).JSON(logResponse)

}

// @Summary Crea un nuevo Ok Log
// @Tags Ok Logs
// @ID create-ok-log
// @Produce json
// @Param data body response.OkLogResponse true "ok log data"
// @Success 200 {object} response.OkLogResponse
// @Failure 400 {object} response.Response
// @Router /api/ok-logs [post]
func(controller *OkLogController) Create(ctx *fiber.Ctx) error {
	createRequest := request.CreateOkLogRequest{}
	
	err := ctx.BodyParser(&createRequest)
	helper.ErrorPanic(err)

	// Guardamos el Ok Log
	controller.okLogService.Create(createRequest)

	// Recuperamos el último Ok Log guardado (es decir, el que recién creamos)
	// y devolvemos con toda la información que completó la BD
	logResponse := controller.okLogService.FindLatest()

	return ctx.Status(fiber.StatusCreated).JSON(logResponse)

}
