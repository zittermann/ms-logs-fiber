package router

import (
	"ms_logs_go/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(okLogController *controller.OkLogController,
	nokLogController *controller.NokLogController) *fiber.App {

	router := fiber.New()

	router.Route("/ok-logs", func(router fiber.Router) {
		router.Get("/:id", okLogController.FindById)
		router.Post("/", okLogController.Create)
	})

	router.Route("/nok-logs", func(router fiber.Router) {
		router.Get("/:id", nokLogController.FindById)
		router.Post("/", nokLogController.Create)
	})

	return router

}

// func NewOkLogRouter(okLogController *controller.OkLogController) *fiber.App {

// 	router := fiber.New()

// 	router.Route("/ok-logs", func(router fiber.Router) {
// 		router.Get("/:id", okLogController.FindById)
// 		router.Post("/", okLogController.Create)
// 	})

// 	return router

// }

// func NewNokLogRouter(nokLogController *controller.NokLogController) *fiber.App {
	
// 	router := fiber.New()

// 	router.Route("/nok-logs", func(router fiber.Router) {
// 		router.Get("/:id", nokLogController.FindById)
// 		router.Post("/", nokLogController.Create)
// 	})

// 	return router

// }
