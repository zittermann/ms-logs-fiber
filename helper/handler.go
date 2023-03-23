package helper

import (
	"ms_logs_go/data/response"
	"ms_logs_go/util"

	"github.com/gofiber/fiber/v2"
)

func HandleNotFound(ctx *fiber.Ctx, msg string) error {
	
	webResponse := response.Response {
		Timestamp: util.GenerateFecha(),
		Message: msg,
		Error: "404 NOT_FOUND",
	}

	return ctx.Status(fiber.StatusNotFound).JSON(webResponse)

}
