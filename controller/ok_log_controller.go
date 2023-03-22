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
	helper.ErrorPanic(err, err.Error())

	logResponse := controller.okLogService.FindById(uint(id))

	webResponse := response.Response {
		Code: http.StatusOK,
		Status: "Ok",
		Message: "Ok Log encontrado",
		Data: logResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}

func(controller *OkLogController) Create(ctx *fiber.Ctx) error {
	createRequest := request.CreateOkLogRequest{}
	
	err := ctx.BodyParser(&createRequest)
	helper.ErrorPanic(err, err.Error())

	controller.okLogService.Create(createRequest)
	webResponse := response.Response{
		Code: http.StatusOK,
		Status: "Ok",
		Message: "Ok Log creado correctamente",
		Data: nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)

}
