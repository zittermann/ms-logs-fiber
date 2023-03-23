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
