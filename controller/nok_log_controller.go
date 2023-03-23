package controller

import (
	"ms_logs_go/data/request"
	"ms_logs_go/data/response"
	"ms_logs_go/helper"
	"ms_logs_go/service"
	"net/http"
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

	webResponse := response.Response {
		Code: http.StatusOK,
		Status: "Ok",
		Message: "Nok Log encontrado",
		Data: logResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}

func(controller *NokLogController) Create(ctx *fiber.Ctx) error {
	createRequest := request.CreateNokLogRequest{}
	
	err := ctx.BodyParser(&createRequest)
	helper.ErrorPanic(err)

	controller.nokLogService.Create(createRequest)
	webResponse := response.Response{
		Code: http.StatusOK,
		Status: "Ok",
		Message: "Nok Log creado correctamente",
		Data: nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)

}
