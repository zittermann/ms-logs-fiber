package main

import (
	"fmt"
	"log"
	"ms_logs_go/config"
	"ms_logs_go/controller"
	"ms_logs_go/model"
	"ms_logs_go/repository"
	"ms_logs_go/router"
	"ms_logs_go/service"

	// swag init
	// Indispensable importar /docs al main para
	// poder mostrar correctamente el swagger
	_ "ms_logs_go/docs"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title MS Logs Go
// @version 1.0
// @description Microservicio encargado de la creación y recuperación de Logs
// @contact.name Germán Zittermann
// @contact.email zitterman@lacaja.com.ar
// @host localhost:33306
// @BasePath /api/
func main() {

	fmt.Println("Started server!")	
	fmt.Println("Connecting to database...")

	// Database
	db := config.CreateConnection()
	validate := validator.New()

	db.Table("ok_logs").AutoMigrate(&model.OkLog{})
	db.Table("nok_logs").AutoMigrate(&model.NokLog{})
	
	fmt.Println("Connection established!")

	// Repositories
	okLogsRepository := repository.NewOkLogRepositoryImpl(db)
	nokLogsRepository := repository.NewNokLogRepositoryImpl(db)

	// Services
	okLogsService := service.NewOkLogServiceImpl(okLogsRepository, 
		validate)
	nokLogsService := service.NewNokLogServiceImpl(nokLogsRepository, 
		validate) 

	// Controller
	okLogController := controller.NewOkLogController(okLogsService)
	nokLogController := controller.NewNokLogController(nokLogsService)

	// Routes
	routes := router.NewRouter(okLogController, nokLogController)

	app := fiber.New()
	app.Mount("/api", routes)

	// Prometheus config
	prometheus := fiberprometheus.New("my-service-name")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Swagger config
	// El swag debe ser generado previo a correr el servicio
	// swag init && go run .
	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":33306"))

}
