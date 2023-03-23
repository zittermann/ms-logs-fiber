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
